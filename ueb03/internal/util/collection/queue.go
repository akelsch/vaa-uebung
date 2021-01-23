package collection

import (
    "container/list"
    "github.com/akelsch/vaa/ueb03/internal/util/collection/queue"
)

type Queue struct {
    l *list.List
}

func NewQueue() *Queue {
    return &Queue{
        l: list.New(),
    }
}

func (q *Queue) Push(item *queue.Item) {
    q.l.PushBack(item)
}

func (q *Queue) Pop() *queue.Item {
    e := q.l.Front()
    v := q.l.Remove(e)
    return v.(*queue.Item)
}

func (q *Queue) HasNext() bool {
    return q.l.Len() > 0
}
