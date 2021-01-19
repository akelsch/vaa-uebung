package directory

import "sync"

type MessageDirectory struct {
    mu       sync.Mutex
    Flooding *FloodingDirectory
    Election *ElectionDirectory
}

func NewMessageDirectory() *MessageDirectory {
    return &MessageDirectory{
        Flooding: NewFloodingDirectory(),
        Election: NewElectionDirectory(),
    }
}

func (md *MessageDirectory) Lock() {
    md.mu.Lock()
}

func (md *MessageDirectory) Unlock() {
    md.mu.Unlock()
}
