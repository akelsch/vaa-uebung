package conf

import (
    "fmt"
    "sync"
)

type NeighborDirectory struct {
    mu       sync.Mutex
    sent     map[int]bool
    received map[int]bool
}

func NewNeighborDirectory() *NeighborDirectory {
    return &NeighborDirectory{sent: make(map[int]bool), received: make(map[int]bool)}
}

func (nd *NeighborDirectory) Lock() {
    nd.mu.Lock()
}

func (nd *NeighborDirectory) Unlock() {
    nd.mu.Unlock()
}

func (nd *NeighborDirectory) Reset() {
    fmt.Println("* resetting directory *")
    for key := range nd.sent {
        nd.sent[key] = false
    }

    for key := range nd.received {
        nd.received[key] = false
    }
}

func (nd *NeighborDirectory) HasNotSentTo(key int) bool {
    v, ok := nd.sent[key]
    if !ok {
        return true
    }
    return !v
}

func (nd *NeighborDirectory) HasAlreadyReceivedFrom(key int) bool {
    v, ok := nd.received[key]
    if !ok {
        return false
    }
    return v
}

func (nd *NeighborDirectory) SetSent(key int) {
    nd.sent[key] = true
}

func (nd *NeighborDirectory) SetReceived(key int) {
    nd.received[key] = true
}

func (nd *NeighborDirectory) ResetIfNecessary(supposedLen int) {
    sentDone := areAllValuesTrue(nd.sent, supposedLen)
    receivedDone := areAllValuesTrue(nd.received, supposedLen)

    if sentDone && receivedDone {
        nd.Reset()
    }
}

func areAllValuesTrue(m map[int]bool, supposedLen int) bool {
    if len(m) == 0 || len(m) != supposedLen {
        return false
    }

    for _, val := range m {
        if !val {
            return false
        }
    }

    return true
}
