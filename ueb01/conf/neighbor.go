package conf

import "sync"

type NeighborDirectory struct {
    mu sync.Mutex
    v  map[int]bool
}

func NewNeighborDirectory() *NeighborDirectory {
    return &NeighborDirectory{v: make(map[int]bool)}
}

func (nd *NeighborDirectory) Lock() {
    nd.mu.Lock()
}

func (nd *NeighborDirectory) Unlock() {
    nd.mu.Unlock()
}

func (nd *NeighborDirectory) IsRemaining(key int) bool {
    v, ok := nd.v[key]
    if !ok {
        return true
    }
    return !v
}

func (nd *NeighborDirectory) Set(key int) {
    nd.v[key] = true
}
