package directory

import (
    "container/heap"
    "github.com/akelsch/vaa/ueb03/internal/collection"
    "github.com/akelsch/vaa/ueb03/internal/directory/state"
)

// Used for Ricart-Agrawala algorithm
type MutexDirectory struct {
    lc        *collection.LamportClock
    pq        *collection.PriorityQueue
    state     state.State
    responses map[uint64]bool
}

func NewMutexDirectory() *MutexDirectory {
    return &MutexDirectory{
        lc:        &collection.LamportClock{},
        pq:        &collection.PriorityQueue{},
        responses: make(map[uint64]bool),
    }
}

func (md *MutexDirectory) GetTimestamp() uint64 {
    return uint64(md.lc.Time())
}

func (md *MutexDirectory) IncrementTimestamp() uint64 {
    return uint64(md.lc.Increment())
}

func (md *MutexDirectory) UpdateTimestamp(timestamp uint64) {
    md.lc.Witness(collection.LamportTime(timestamp))
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

func (md *MutexDirectory) RegisterWant() {
    md.state = state.WANTED
}

func (md *MutexDirectory) RegisterLock() {
    md.state = state.HELD
}

func (md *MutexDirectory) NeedsToQueue(timestamp uint64, resource uint64, selfId uint64) bool {
    return md.state == state.HELD ||
        (md.state == state.WANTED && md.GetTimestamp() < timestamp) ||
        (md.state == state.WANTED && resource == selfId) // FIXME guarantees correctness but does potentially cause deadlocks
}

func (md *MutexDirectory) RegisterResponse(node uint64) {
    md.responses[node] = true
}

func (md *MutexDirectory) CheckResponseCount(expected int) bool {
    count := 0
    for _, b := range md.responses {
        if b {
            count++
        }
    }

    return count == expected
}

func (md *MutexDirectory) Reset() {
    md.state = state.RELEASED
    md.responses = make(map[uint64]bool)
}
