package directory

import (
    "container/heap"
    "github.com/akelsch/vaa/ueb03/internal/collection"
)

// Used for Ricart-Agrawala algorithm
type MutexDirectory struct {
    lc *collection.LamportClock
    pq *collection.PriorityQueue
}

func NewMutexDirectory() *MutexDirectory {
    return &MutexDirectory{
        lc: &collection.LamportClock{},
        pq: &collection.PriorityQueue{},
    }
}

func (md *MutexDirectory) GetTimestamp() uint64 {
    return uint64(md.lc.Time())
}

func (md *MutexDirectory) IncrementTimestampBy(n int) uint64 {
    return uint64(md.lc.IncrementBy(uint64(n)))
}

func (md *MutexDirectory) UpdateTimestamp(timestamp uint64) {
    md.lc.Witness(collection.LamportTime(timestamp))
}

func (md *MutexDirectory) IsUsingResource(resource uint64) bool {
    return md.pq.ContainsValue(resource)
}

func (md *MutexDirectory) QueueLockRequest(resource uint64, timestamp uint64) {
    heap.Push(md.pq, md.pq.NewItem(resource, timestamp))
}
