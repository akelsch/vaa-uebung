package directory

import "sync"

type MessageDirectory struct {
    mu        sync.Mutex
    Neighbors *NeighborDirectory
    Election  *ElectionDirectory
}

func NewMessageDirectory() *MessageDirectory {
    return &MessageDirectory{
        Neighbors: NewNeighborDirectory(),
        Election:  NewElectionDirectory(),
    }
}

func (md *MessageDirectory) Lock() {
    md.mu.Lock()
}

func (md *MessageDirectory) Unlock() {
    md.mu.Unlock()
}
