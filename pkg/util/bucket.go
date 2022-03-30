package util

import (
	"github.com/libp2p/go-libp2p-core/peer"
	kbucket "github.com/libp2p/go-libp2p-kbucket"
)

const (
	DefaultBucketSize = 20
)

func BucketIdForPeer(localPeer, remotePeer peer.ID) int16 {
	peerID := kbucket.ConvertPeerID(remotePeer)
	cpl := kbucket.CommonPrefixLen(peerID, kbucket.ConvertPeerID(localPeer))
	bucketID := cpl
	if bucketID >= DefaultBucketSize {
		bucketID = DefaultBucketSize - 1
	}
	return int16(bucketID)
}
