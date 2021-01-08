package directory

import "sync"

type MessageDirectory struct {
    mu        sync.Mutex
    Neighbors *NeighborDirectory
    Rumors    *RumorDirectory
    Election  *ElectionDirectory
    Status    *StatusDirectory
}

func NewMessageDirectory() *MessageDirectory {
    return &MessageDirectory{
        Neighbors: NewNeighborDirectory(),
        Rumors:    NewRumorDirectory(),
        Election:  NewElectionDirectory(),
        Status:    NewStatusDirectory(),
    }
}

func (md *MessageDirectory) Lock() {
    md.mu.Lock()
}

func (md *MessageDirectory) Unlock() {
    md.mu.Unlock()
}
