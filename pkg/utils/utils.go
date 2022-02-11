package utils

import (
	"github.com/libp2p/go-libp2p-core/peer"
	kbucket "github.com/libp2p/go-libp2p-kbucket"
)

const (
	DefaultBucketSize = 20
)

// IDLength is here as a variable so that it can be decreased for tests with mocknet where IDs are way shorter.
// The call to FmtPeerID would panic if this value stayed at 16.
var IDLength = 16

func FmtPeerID(id peer.ID) string {
	if len(id.Pretty()) <= IDLength {
		return id.Pretty()
	}
	return id.Pretty()[:IDLength]
}

func BucketIdForPeer(localPeer, remotePeer peer.ID) int16 {
	peerID := kbucket.ConvertPeerID(remotePeer)
	cpl := kbucket.CommonPrefixLen(peerID, kbucket.ConvertPeerID(localPeer))
	bucketID := cpl
	if bucketID >= DefaultBucketSize {
		bucketID = DefaultBucketSize - 1
	}
	return int16(bucketID)
}
