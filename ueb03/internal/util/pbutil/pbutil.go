package pbutil

import (
    "fmt"
    "github.com/akelsch/vaa/ueb03/api/pb"
)

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

func CreateControlMessage(sender uint64, command pb.ControlMessage_Command) *pb.Message {
    return &pb.Message{
        Sender: sender,
        Msg: &pb.Message_ControlMessage{
            ControlMessage: &pb.ControlMessage{
                Command: command,
            },
        },
    }
}

func CreateApplicationMessage(metadata *Metadata, balance int64, percent uint64) *pb.Message {
    return &pb.Message{
        Identifier: metadata.identifier,
        Sender:     metadata.sender,
        Receiver:   metadata.receiver,
        Msg: &pb.Message_ApplicationMessage{
            ApplicationMessage: &pb.ApplicationMessage{
                Balance: balance,
                Percent: percent,
            },
        },
    }
}

func CreateApplicationRequestMessage(metadata *Metadata, percent uint64) *pb.Message {
    return &pb.Message{
        Identifier: metadata.identifier,
        Sender:     metadata.sender,
        Receiver:   metadata.receiver,
        Msg: &pb.Message_ApplicationMessage{
            ApplicationMessage: &pb.ApplicationMessage{
                Type:    pb.ApplicationMessage_REQ,
                Percent: percent,
            },
        },
    }
}

func CreateApplicationResponseMessage(metadata *Metadata, balance int64, percent uint64) *pb.Message {
    return &pb.Message{
        Identifier: metadata.identifier,
        Sender:     metadata.sender,
        Receiver:   metadata.receiver,
        Msg: &pb.Message_ApplicationMessage{
            ApplicationMessage: &pb.ApplicationMessage{
                Type:    pb.ApplicationMessage_RES,
                Balance: balance,
                Percent: percent,
            },
        },
    }
}

func CreateApplicationAcknowledgmentMessage(metadata *Metadata) *pb.Message {
    return &pb.Message{
        Identifier: metadata.identifier,
        Sender:     metadata.sender,
        Receiver:   metadata.receiver,
        Msg: &pb.Message_ApplicationMessage{
            ApplicationMessage: &pb.ApplicationMessage{
                Type: pb.ApplicationMessage_ACK,
            },
        },
    }
}

func CreateExplorerMessage(sender uint64, initiator string) *pb.Message {
    return &pb.Message{
        Sender: sender,
        Msg: &pb.Message_Election{
            Election: &pb.Election{
                Type:      pb.Election_EXPLORER,
                Initiator: initiator,
            },
        },
    }
}

func CreateEchoMessage(sender uint64, initiator string) *pb.Message {
    return &pb.Message{
        Sender: sender,
        Msg: &pb.Message_Election{
            Election: &pb.Election{
                Type:      pb.Election_ECHO,
                Initiator: initiator,
            },
        },
    }
}
