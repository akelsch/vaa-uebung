package directory

import (
    "container/heap"
    "github.com/akelsch/vaa/ueb03/internal/collection"
)

// Used for Ricart-Agrawala algorithm
type MutexDirectory struct {
    lc *collection.LamportClock
    pq *collection.PriorityQueue
    ok map[uint64]bool
}

func NewMutexDirectory() *MutexDirectory {
    return &MutexDirectory{
        lc: &collection.LamportClock{},
        pq: &collection.PriorityQueue{},
        ok: make(map[uint64]bool),
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
    return md.pq.ContainsResource(resource)
}

func (md *MutexDirectory) PushLockRequest(sender uint64, resource uint64, timestamp uint64) {
    heap.Push(md.pq, md.pq.NewItem(sender, resource, timestamp))
}

func (md *MutexDirectory) PopLockRequest() *collection.Item {
    if md.pq.HasNext() {
        item := heap.Pop(md.pq).(*collection.Item)
        return item
    }

    return nil
}

func (md *MutexDirectory) RegisterOk(node uint64) {
    md.ok[node] = true
}

func (md *MutexDirectory) CheckIfAllOk(expected int) bool {
    count := 0
    for _, b := range md.ok {
        if b {
            count++
        }
    }

    return count == expected
}

func (md *MutexDirectory) ResetOk() {
    md.ok = make(map[uint64]bool)
}
