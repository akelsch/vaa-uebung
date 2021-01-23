package directory

import "sync"

type Directory struct {
    mu       sync.Mutex
    Flooding *FloodingDirectory
    Mutex    *MutexDirectory
    Election *ElectionDirectory
}

func NewDirectory() *Directory {
    return &Directory{
        Flooding: NewFloodingDirectory(),
        Mutex:    NewMutexDirectory(),
        Election: NewElectionDirectory(),
    }
}

func (md *Directory) Lock() {
    md.mu.Lock()
}

func (md *Directory) Unlock() {
    md.mu.Unlock()
}
