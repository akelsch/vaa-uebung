package directory

import (
    "github.com/akelsch/vaa/ueb03/internal/config"
    "sync"
)

type Directory struct {
    mu       sync.Mutex
    Flooding *FloodingDirectory
    Mutex    *MutexDirectory
    Election *ElectionDirectory
    Snapshot *SnapshotDirectory
}

func NewDirectory(conf *config.Config) *Directory {
    return &Directory{
        Flooding: NewFloodingDirectory(),
        Mutex:    NewMutexDirectory(),
        Election: NewElectionDirectory(),
        Snapshot: NewSnapshotDirectory(conf.Params.Balance),
    }
}

func (md *Directory) Lock() {
    md.mu.Lock()
}

func (md *Directory) Unlock() {
    md.mu.Unlock()
}
