package account

import "sync"

var noceMap sync.Map

type Nonce struct {
	sync.Mutex

	nonce uint64
}

func NewNonce(nonce uint64) *Nonce {
	return &Nonce{nonce: nonce}
}

func (n *Nonce) Value() uint64 {
	n.Lock()
	defer n.Unlock()
	return n.nonce
}

func (n *Nonce) Next(nonces ...uint64) uint64 {
	n.Lock()
	defer n.Unlock()
	var nonce uint64
	if len(nonces) > 0 && nonces[0] > 0 {
		nonce = nonces[0]
	}
	if n.nonce < nonce {
		n.nonce = nonce
		return nonce
	}
	n.nonce++
	return n.nonce
}

func (n *Nonce) Reset(nonce uint64, omitLower ...bool) {
	n.Lock()
	defer n.Unlock()
	if len(omitLower) > 0 && omitLower[0] && nonce <= n.nonce {
		return
	}
	n.nonce = nonce

}
