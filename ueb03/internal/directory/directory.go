package directory

import "sync"

type Directory struct {
    mu       sync.Mutex
    Flooding *FloodingDirectory
    Mutex    *MutexDirectory
    Election *ElectionDirectory
    Snapshot *SnapshotDirectory
}

func NewDirectory() *Directory {
    return &Directory{
        Flooding: NewFloodingDirectory(),
        Mutex:    NewMutexDirectory(),
        Election: NewElectionDirectory(),
        Snapshot: NewSnapshotDirectory(),
    }
}

func (md *Directory) Lock() {
    md.mu.Lock()
}

func (md *Directory) Unlock() {
    md.mu.Unlock()
}
