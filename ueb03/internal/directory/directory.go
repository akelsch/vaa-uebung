package directory

import "sync"

type Directory struct {
    mu       sync.Mutex
    Flooding *FloodingDirectory
    //Election *ElectionDirectory
}

func NewDirectory() *Directory {
    return &Directory{
        Flooding: NewFloodingDirectory(),
        //Election: NewElectionDirectory(),
    }
}

func (md *Directory) Lock() {
    md.mu.Lock()
}

func (md *Directory) Unlock() {
    md.mu.Unlock()
}
