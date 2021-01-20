package pbutil

import "fmt"

type Metadata struct {
    identifier string
    sender     uint64
    receiver   uint64
}

func CreateMetadata(sender uint64, receiver uint64, seq uint64) *Metadata {
    return &Metadata{
        identifier: fmt.Sprintf("%d-%d", sender, seq),
        sender:     sender,
        receiver:   receiver,
    }
}

func (m *Metadata) GetIdentifier() string {
    return m.identifier
}
