package directory

type NeighborDirectory struct {
    sent     map[int]int
    received map[int]int
}

func NewNeighborDirectory() *NeighborDirectory {
    return &NeighborDirectory{
        sent:     make(map[int]int),
        received: make(map[int]int),
    }
}

func (nd *NeighborDirectory) Stats() (int, int) {
    sentCount := 0
    receivedCount := 0

    for key := range nd.sent {
        sentCount += nd.sent[key]
    }

    for key := range nd.received {
        receivedCount += nd.received[key]
    }

    return sentCount, receivedCount
}

func (nd *NeighborDirectory) SetSent(key int) {
    nd.sent[key]++
}

func (nd *NeighborDirectory) SetReceived(key int) {
    nd.received[key]++
}
