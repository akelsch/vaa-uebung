package directory

import (
    "container/heap"
    "github.com/akelsch/vaa/ueb03/internal/collection"
)

// Used for Ricart-Agrawala algorithm
type MutexDirectory struct {
    lc        *collection.LamportClock
    pq        *collection.PriorityQueue
    current   uint64          // Represents the resource the node is currently interested in
    responses map[uint64]bool // Tracks mutex responses from other nodes
    locked    bool            // Flag that indicates whether the critical section got entered (true) or not (false)
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

func (md *MutexDirectory) IncrementTimestampBy(n int) uint64 {
    return uint64(md.lc.IncrementBy(uint64(n)))
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

func (md *MutexDirectory) RegisterCurrentResource(resource uint64) {
    md.current = resource
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

func (md *MutexDirectory) RegisterLock() {
    md.locked = true
}

func (md *MutexDirectory) Reset() {
    md.current = 0
    md.responses = make(map[uint64]bool)
    md.locked = false
}

func (md *MutexDirectory) IsInterestedIn(resource uint64, ownId uint64) bool {
    isInterestedInSameResource := resource == md.current
    isRequestingAndInterestedInSelf := md.current != 0 && resource == ownId
    return isInterestedInSameResource || isRequestingAndInterestedInSelf
}

func (md *MutexDirectory) IsUsing(resource uint64, ownId uint64, otherId uint64) bool {
    if md.locked {
        return resource == md.current || md.current == otherId
    }

    return resource == ownId // FIXME causing deadlocks
}

func (md *MutexDirectory) IsLowerPriority(otherTimestamp uint64, otherId uint64, ownId uint64) bool {
    ownTimestamp := md.GetTimestamp()

    // handle concurrent events so the queue does not break
    if otherTimestamp == ownTimestamp {
        // let the node with the smaller id win
        return otherId < ownId
    }

    return otherTimestamp < ownTimestamp
}
