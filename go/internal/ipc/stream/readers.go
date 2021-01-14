package stream

import (
	"sync"

	"github.com/diamondburned/cchat-ipc/go/internal/ipc/core"
)

// MaxBufferSize is the constant for the maximum buffer size allowed. It is
// capped to 512MB.
const MaxBufferSize = 512 * 1024 * 1024

func maxBufSz(bufSz uint32) uint32 {
	if bufSz > MaxBufferSize {
		return MaxBufferSize
	}
	return bufSz
}

// ReceiverMap is a map of stream receivers.
type ReceiverMap struct {
	mu      sync.RWMutex
	tickets map[core.StreamTicket]*receiver
}

// ReceiverFunc is the callback type that takes in an incoming stream data. The
// callback must not retain the byte slice. If false is returned, then the
// stream is removed from the registry and the callback won't be called again.
type ReceiverFunc func(core.StreamData) (ok bool)

type receiver struct {
	fn ReceiverFunc
}

// Register registers a reader.
func (smap *ReceiverMap) RegisterReader(s core.Stream, fn ReceiverFunc) {
	smap.mu.Lock()
	smap.tickets[s.Ticket] = &receiver{fn}
	smap.mu.Unlock()
}

// WriteStream writes the given stream data to the reader associated with it.
// It does nothing if the given ticket is not found.
func (smap *ReceiverMap) WriteStream(dat core.StreamData) {
	smap.mu.RLock()
	stream, ok := smap.tickets[dat.Ticket]
	smap.mu.RUnlock()
	if !ok {
		return
	}

	if !stream.fn(dat) || dat.Error != nil {
		smap.mu.Lock()
		delete(smap.tickets, dat.Ticket)
		smap.mu.Unlock()
	}
}
