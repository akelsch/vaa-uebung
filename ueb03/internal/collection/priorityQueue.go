package collection

// Original source: https://golang.org/pkg/container/heap/

// An Item is something we manage in a priority queue (lock request).
type Item struct {
    Sender   uint64 // sender id of the lock request
    Resource uint64 // resource id of the lock request
    priority uint64 // priority of the item in the queue (lamport clock timestamp)
    index    int    // index of the item in the heap (maintained by heap.Interface methods)
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
    return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
    // Lowest lamport clock timestamp has highest priority
    return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].index = i
    pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
    n := len(*pq)
    item := x.(*Item)
    item.index = n
    *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    old[n-1] = nil  // avoid memory leak
    item.index = -1 // for safety
    *pq = old[0 : n-1]
    return item
}

func (pq *PriorityQueue) NewItem(sender uint64, resource uint64, priority uint64) *Item {
    return &Item{
        Sender:   sender,
        Resource: resource,
        priority: priority,
    }
}

func (pq *PriorityQueue) ContainsResource(resource uint64) bool {
    for _, item := range *pq {
        if item.Resource == resource {
            return true
        }
    }

    return false
}

func (pq *PriorityQueue) HasNext() bool {
    return pq.Len() > 0
}
