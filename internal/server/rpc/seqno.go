package rpc

import (
	"sync"
)

var (
	seqnoMu  sync.Mutex
	seqnoMap = map[uint32]chan interface{}{}
)

func waitSeqno(seqno uint32) chan interface{} {
	seqnoMu.Lock()
	defer seqnoMu.Unlock()

	ch, ok := seqnoMap[seqno]
	if ok {
		return ch
	}

	ch = make(chan interface{}, 1)
	seqnoMap[seqno] = ch
	return ch
}

func stopSeqno(seqno uint32, op interface{}) {
	seqnoMu.Lock()
	defer seqnoMu.Unlock()

	ch, ok := seqnoMap[seqno]
	if !ok {
		return
	}
	ch <- op
	delete(seqnoMap, seqno)
}
