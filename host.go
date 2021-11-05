package main

import (
	"context"
	"sort"
	"time"

	"bou.ke/monkey"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/routing"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	pb "github.com/libp2p/go-libp2p-kad-dht/pb"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Host struct {
	id            string
	h             host.Host
	dht           *kaddht.IpfsDHT
	sc            chan Span
	cancelDiscard context.CancelFunc
}

func NewHost(ctx context.Context, optsFunc func(peer.ID, chan Span) libp2p.Option) (*Host, error) {
	ec := make(chan Span)

	key, _, err := crypto.GenerateKeyPair(crypto.Secp256k1, 256)
	if err != nil {
		return nil, errors.Wrap(err, "generate key pair")
	}

	localID, err := peer.IDFromPublicKey(key.GetPublic())
	if err != nil {
		return nil, errors.Wrap(err, "id from public key")
	}

	ms := &messageSenderImpl{
		protocols: kaddht.DefaultProtocols,
		strmap:    make(map[peer.ID]*peerMessageSender),
		local:     localID,
		sc:        ec,
	}
	pm, err := pb.NewProtocolMessenger(ms)
	if err != nil {
		return nil, err
	}

	// When kaddht tries to instantiate a new protocol messenger hand it our implementation. There is no option to
	// exchange the protocol messenger or message sender implementation.
	monkey.Patch(pb.NewProtocolMessenger, func(msgSender pb.MessageSender, opts ...pb.ProtocolMessengerOption) (*pb.ProtocolMessenger, error) {
		return pm, nil
	})

	var dht *kaddht.IpfsDHT
	h, err := libp2p.New(ctx,
		libp2p.Identity(key),
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			dht, err = kaddht.New(ctx, h)
			return dht, err
		}),
		optsFunc(localID, ec),
	)
	if err != nil {
		return nil, errors.Wrap(err, "new libp2p host")
	}

	// Set the remaining information
	// TODO: race condition if the above init already tries to access these fields?
	ms.host = h

	// Remove monkey patched function.
	monkey.UnpatchAll()

	newHost := &Host{
		id:  "host",
		h:   h,
		dht: dht,
		sc:  ec,
	}
	newHost.startDiscarding(ctx)

	return newHost, nil
}

func (h *Host) Init(ctx context.Context, refreshRt bool) error {
	if err := h.Bootstrap(ctx); err != nil {
		return err
	}
	// Check if we should wait until the routing table of the provider was refreshed.
	if refreshRt {
		h.RefreshRoutingTable(ctx)
	}
	return nil
}

func (h *Host) Bootstrap(ctx context.Context) error {
	for _, bp := range kaddht.GetDefaultBootstrapPeerAddrInfos() {
		h.logEntry().WithField("remoteID", FmtPeerID(bp.ID)).Infoln("Connecting to bootstrap peer")
		if err := h.h.Connect(ctx, bp); err != nil {
			return errors.Wrap(err, "connecting to bootstrap peer")
		}
	}
	return nil
}

func (h *Host) RefreshRoutingTable(ctx context.Context) {
	h.logEntry().Infoln("Start refreshing routing table")
	defer h.logEntry().Infoln("Done refreshing routing table")
	select {
	case <-h.dht.RefreshRoutingTable():
	case <-ctx.Done():
	}
}

func (h *Host) Run(ctx context.Context, content *Content, fn func(context.Context, *Content) error) (*Run, error) {
	h.stopDiscarding()
	defer h.startDiscarding(ctx)

	startTime := time.Now()
	involved := map[peer.ID]bool{}
	spans := []Span{}

	queryCtx, queryEvents := routing.RegisterForQueryEvents(ctx)
	done := make(chan error)
	go func() { done <- fn(queryCtx, content) }()

	for {
		select {
		case event := <-queryEvents:
			involved[event.ID] = false
		case span := <-h.sc:
			if span.LocalID() == h.h.ID() {
				spans = append(spans, span)
			}
		case err := <-done:
			if err != nil {
				return nil, err
			}
			var filtered []Span
			for _, span := range spans {
				if _, isInvolved := involved[span.RemoteID()]; isInvolved {
					filtered = append(filtered, span)
					involved[span.RemoteID()] = true
				}
			}
			sort.SliceStable(filtered, func(i, j int) bool {
				return filtered[i].StartedAt().Before(filtered[j].StartedAt())
			})
			for peerID, spanExists := range involved {
				if !spanExists {
					// this is safe it seems https://stackoverflow.com/questions/23229975/is-it-safe-to-remove-selected-keys-from-map-within-a-range-loop
					delete(involved, peerID)
				}
			}

			return &Run{
				StartedAt: startTime,
				EndedAt:   time.Now(),
				LocalID:   h.h.ID(),
				Spans:     filtered,
				Involved:  involved,
			}, nil
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}

func (h *Host) logEntry() *log.Entry {
	return log.WithField("type", h.id)
}

func (h *Host) startDiscarding(ctx context.Context) {
	if h.cancelDiscard != nil {
		// already discarding
		return
	}

	cancelCtx, cancel := context.WithCancel(ctx)
	go func() {
		for {
			select {
			case <-h.sc:
				// discard spans
			case <-cancelCtx.Done():
				return
			}
		}
	}()
	h.cancelDiscard = cancel
}

func (h *Host) stopDiscarding() {
	if h.cancelDiscard != nil {
		h.cancelDiscard()
		h.cancelDiscard = nil
	}
}
