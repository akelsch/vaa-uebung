package queue

type Item struct {
    Sender   uint64
    Resource uint64
}

func NewItem(sender uint64, resource uint64) *Item {
    return &Item{
        Sender:   sender,
        Resource: resource,
    }
}
