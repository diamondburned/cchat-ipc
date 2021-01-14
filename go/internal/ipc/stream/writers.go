package stream

import (
	"github.com/diamondburned/cchat-ipc/go/internal/ipc/core"
	"github.com/diamondburned/cchat-ipc/go/internal/ipc/registry/slab"
)

// WriterMap is a map of stream writers.
type WriterMap struct {
	slab slab.SlabMutex
}

// stub type in case of future additions.
type writer struct{}

// NewTicket creates a new stream ticket for writing.
func (wmap *WriterMap) NewTicket() core.StreamTicket {
	return core.StreamTicket(wmap.slab.Put(writer{}))
}
