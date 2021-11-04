package main

import (
	"context"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Provider struct {
	*Host
}

func NewProvider(ctx context.Context) (*Provider, error) {
	h, err := NewHost(ctx, func(localID peer.ID, ec chan Event) libp2p.Option {
		return libp2p.ChainOptions(
			libp2p.DefaultListenAddrs,
			libp2p.Transport(NewTCPTransport(localID, ec)),
			libp2p.Transport(NewWSTransport(localID, ec)),
		)
	})
	if err != nil {
		return nil, errors.Wrap(err, "new host")
	}
	h.id = "provider"

	log.WithField("providerID", FmtPeerID(h.h.ID())).Infoln("Initialized provider")
	return &Provider{
		Host: h,
	}, nil
}

func (p *Provider) RunAction(ctx context.Context, content *Content) error {
	p.logEntry().Infoln("Start providing content")
	defer p.logEntry().Infoln("Done providing content")
	return p.dht.Provide(ctx, content.cid, true)
}
