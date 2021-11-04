package main

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"go.uber.org/atomic"
)

var requesterID = atomic.NewInt32(0)

type Requester struct {
	*Host
	provider peer.AddrInfo
}

func NewRequester(ctx context.Context) (*Requester, error) {
	h, err := NewHost(ctx, func(localID peer.ID, ec chan Span) libp2p.Option {
		return libp2p.ChainOptions(
			libp2p.NoListenAddrs,
		)
	})
	if err != nil {
		return nil, errors.Wrap(err, "new host")
	}
	h.id = fmt.Sprintf("requester-%02d", requesterID.Inc())

	log.WithField("requesterID", FmtPeerID(h.h.ID())).Infoln("Initialized " + h.id)
	return &Requester{
		Host: h,
	}, nil
}

func (r *Requester) RunAction(ctx context.Context, content *Content) error {
	r.logEntry().Infoln("Start requesting content")
	defer r.logEntry().Infoln("Done requesting content")

	pi, ok := <-r.dht.FindProvidersAsync(ctx, content.cid, 1)
	r.logEntry().WithField("providerID", FmtPeerID(pi.ID)).Infoln("Found provider")
	r.provider = pi
	if ok {
		return nil
	}
	return fmt.Errorf("not ok")
}
