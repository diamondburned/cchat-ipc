package registry

import (
	"context"
	"sync"

	"github.com/diamondburned/cchat-ipc/go/internal/ipc/core"
)

// ContextMap implements a map of call IDs to context cancellation functions.
type ContextMap struct {
	mu   sync.Mutex
	ctxs map[core.CallID]context.CancelFunc
}

// New creates a new context associated with the given call ID.
func (ctxMap *ContextMap) New(callID core.CallID) context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	ctxMap.mu.Lock()
	ctxMap.ctxs[callID] = cancel
	ctxMap.mu.Unlock()

	return ctx
}

// Remove cancels the context associated with the given call ID and removes it.
func (ctxMap *ContextMap) Remove(callID core.CallID) {
	ctxMap.mu.Lock()
	cancel, ok := ctxMap.ctxs[callID]
	delete(ctxMap.ctxs, callID)
	ctxMap.mu.Unlock()

	if ok {
		cancel()
	}
}
