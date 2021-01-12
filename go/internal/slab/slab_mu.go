package slab

import "sync"

// SlabMutex implements a thread-safe free list.
type SlabMutex struct {
	mut  sync.RWMutex
	slab *Slab
}

func (smu *SlabMutex) Put(entry interface{}) uint64 {
	smu.mut.Lock()
	id := smu.slab.Put(entry)
	smu.mut.Unlock()

	return id
}

func (smu *SlabMutex) Get(i uint64) interface{} {
	smu.mut.RLock()
	v := smu.slab.Get(i)
	smu.mut.RUnlock()

	return v
}

func (smu *SlabMutex) Pop(i uint64) interface{} {
	smu.mut.Lock()
	v := smu.slab.Pop(i)
	smu.mut.Unlock()

	return v
}
