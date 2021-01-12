package cchatipc

import "github.com/diamondburned/cchat-ipc/languages/go/internal/slab"

// Repository contains state-keeping structures of shared object instances.
type Repository struct {
	objects  slab.SlabMutex // T: any
	contexts slab.SlabMutex // T: context.Context
}
