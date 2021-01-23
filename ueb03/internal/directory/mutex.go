package directory

import (
    "github.com/akelsch/vaa/ueb03/internal/directory/mutex/state"
    "github.com/akelsch/vaa/ueb03/internal/util/collection"
    "github.com/akelsch/vaa/ueb03/internal/util/collection/queue"
)

// Used for Ricart-Agrawala algorithm
type MutexDirectory struct {
    lc        *collection.LamportClock
    queue     *collection.Queue
    state     state.State
    responses int
}

func NewMutexDirectory() *MutexDirectory {
    return &MutexDirectory{
        lc:    &collection.LamportClock{},
        queue: collection.NewQueue(),
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

func (md *MutexDirectory) PushLockRequest(sender uint64, resource uint64) {
    item := queue.NewItem(sender, resource)
    md.queue.Push(item)
}

func (md *MutexDirectory) PopLockRequests(f func(item *queue.Item)) {
    for md.queue.HasNext() {
        item := md.queue.Pop()
        f(item)
    }
}

func (md *MutexDirectory) RegisterWant() {
    md.state = state.WANTED
}

func (md *MutexDirectory) RegisterLock() {
    md.state = state.HELD
}

func (md *MutexDirectory) NeedsToQueue(timestamp uint64, resource uint64, selfId uint64) bool {
    return md.state == state.HELD ||
        (md.state == state.WANTED && md.GetTimestamp() < timestamp) /* ||
       (md.state == state.WANTED && resource == selfId) // FIXME guarantees correctness but does potentially cause deadlocks*/
}

func (md *MutexDirectory) RegisterResponse() {
    md.responses++
}

func (md *MutexDirectory) CheckResponseCount(expected int) bool {
    return md.responses == expected
}

func (md *MutexDirectory) Reset() {
    md.state = state.RELEASED
    md.responses = 0
}
