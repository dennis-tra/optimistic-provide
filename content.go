package main

import (
	"crypto/rand"
	"crypto/sha256"

	u "github.com/ipfs/go-ipfs-util"
	"github.com/libp2p/go-libp2p-core/peer"
	kbucket "github.com/libp2p/go-libp2p-kbucket"

	"github.com/ipfs/go-cid"
	mh "github.com/multiformats/go-multihash"
	"github.com/pkg/errors"
)

// Content encapsulates multiple representations of the same data.
type Content struct {
	raw   []byte
	mhash mh.Multihash
	cid   cid.Cid
}

// NewRandomContent reads 1024 bytes from crypto/rand and builds a content struct.
func NewRandomContent() (*Content, error) {
	raw := make([]byte, 1024)
	if _, err := rand.Read(raw); err != nil {
		return nil, errors.Wrap(err, "read rand data")
	}
	hash := sha256.New()
	hash.Write(raw)

	mhash, err := mh.Encode(hash.Sum(nil), mh.SHA2_256)
	if err != nil {
		return nil, errors.Wrap(err, "encode multi hash")
	}

	return &Content{
		raw:   raw,
		mhash: mhash,
		cid:   cid.NewCidV0(mhash),
	}, nil
}

func (c *Content) DistanceTo(peerID peer.ID) []byte {
	return u.XOR(kbucket.ConvertPeerID(peerID), kbucket.ConvertKey(string(c.mhash)))
}
