package directory

import (
    "container/heap"
    "github.com/akelsch/vaa/ueb03/internal/collection"
)

// Used for Ricart-Agrawala algorithm
type MutexDirectory struct {
    lc  *collection.LamportClock
    pq  *collection.PriorityQueue
    res map[uint64]bool
    cur uint64
}

func NewMutexDirectory() *MutexDirectory {
    return &MutexDirectory{
        lc:  &collection.LamportClock{},
        pq:  &collection.PriorityQueue{},
        res: make(map[uint64]bool),
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

func (md *MutexDirectory) RegisterCurrentResource(resource uint64) {
    md.cur = resource
}

func (md *MutexDirectory) ResetCurrentResource() {
    md.cur = 0
}

func (md *MutexDirectory) IsInterestedInResource(resource uint64) bool {
    return resource == md.cur || md.pq.ContainsResource(resource)
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

func (md *MutexDirectory) RegisterResponse(node uint64) {
    md.res[node] = true
}

func (md *MutexDirectory) ResetResponses() {
    md.res = make(map[uint64]bool)
}

func (md *MutexDirectory) CheckResponseCount(expected int) bool {
    count := 0
    for _, b := range md.res {
        if b {
            count++
        }
    }

    return count == expected
}

func (md *MutexDirectory) HasLowerPriority(otherTimestamp uint64, otherId uint64, ownId uint64) bool {
    ownTimestamp := md.GetTimestamp()

    // handle concurrent events so the queue does not break
    if otherTimestamp == ownTimestamp {
        // let the node with the smaller id win
        return otherId < ownId
    }

    return otherTimestamp < ownTimestamp
}
