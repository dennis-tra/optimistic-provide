package wrap

import (
	"sync"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

type Notifee interface {
	DialStarted(trpt string, raddr ma.Multiaddr, p peer.ID, start time.Time)
	DialEnded(trpt string, raddr ma.Multiaddr, p peer.ID, start time.Time, end time.Time, err error)
}

type Notifier struct {
	trpt     string
	notiflk  sync.RWMutex
	notifees map[Notifee]struct{}
}

func newNotifier(trpt string) *Notifier {
	return &Notifier{
		trpt:     trpt,
		notiflk:  sync.RWMutex{},
		notifees: map[Notifee]struct{}{},
	}
}

func (n *Notifier) notifyDialStarted(raddr ma.Multiaddr, p peer.ID, start time.Time) {
	n.notiflk.RLock()
	defer n.notiflk.RUnlock()

	for notifee := range n.notifees {
		notifee.DialStarted(n.trpt, raddr, p, start)
	}
}

func (n *Notifier) notifyDialEnded(raddr ma.Multiaddr, p peer.ID, start time.Time, end time.Time, err error) {
	n.notiflk.RLock()
	defer n.notiflk.RUnlock()

	for notifee := range n.notifees {
		notifee.DialEnded(n.trpt, raddr, p, start, end, err)
	}
}

func (n *Notifier) Notify(notifee Notifee) {
	n.notiflk.Lock()
	defer n.notiflk.Unlock()
	n.notifees[notifee] = struct{}{}
}

func (n *Notifier) StopNotify(notifee Notifee) {
	n.notiflk.Lock()
	defer n.notiflk.Unlock()
	delete(n.notifees, notifee)
}
