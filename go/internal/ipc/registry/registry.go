package registry

import (
	"github.com/diamondburned/cchat-ipc/go/internal/ipc/core"
	"github.com/diamondburned/cchat-ipc/go/internal/ipc/registry/slab"
	"github.com/diamondburned/cchat-ipc/go/internal/ipc/stream"
)

// Registry describes a global registry of instances, stop callbacks and
// contexts.
type Registry struct {
	instances slab.SlabMutex
	stopFuncs slab.SlabMutex

	Receivers stream.ReceiverMap
	Contexts  ContextMap
}

// stopFunc is a container to wrap a stop callback with more data.
type stopFunc struct {
	stopFn func()
	instID core.InstanceID
}

// RegisterStopFn registers a stop callback in relation to the given instance ID
// and returns a unique stop handle as a reference to said callback. The given
// instance ID will be used to remove the instance off the registry once the
// stop callback is called.
func (reg *Registry) RegisterStopFn(instID core.InstanceID, stop func()) core.StopHandle {
	return core.StopHandle(reg.stopFuncs.Put(stopFunc{
		stopFn: stop,
		instID: instID,
	}))
}

// Stop pops the stop callback with the given handle and calls it. It will also
// unregister the instance related to the stop callback.
func (reg *Registry) Stop(handle core.StopHandle) {
	fn, ok := reg.stopFuncs.Pop(uint64(handle)).(stopFunc)
	if !ok {
		return
	}
	fn.stopFn()
	reg.instances.Pop(uint64(fn.instID))
}
