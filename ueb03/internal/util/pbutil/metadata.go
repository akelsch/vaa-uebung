package pbutil

import "fmt"

type Metadata struct {
    Identifier string
    sender     uint64
    receiver   uint64
}

func CreateMetadata(sender uint64, receiver uint64, seq uint64) *Metadata {
    return &Metadata{
        Identifier: fmt.Sprintf("%d-%d", sender, seq),
        sender:     sender,
        receiver:   receiver,
    }
}
