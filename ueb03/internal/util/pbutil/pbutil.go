package pbutil

import "github.com/akelsch/vaa/ueb03/api/pb"

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
        Identifier: metadata.Identifier,
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
        Identifier: metadata.Identifier,
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
        Identifier: metadata.Identifier,
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
        Identifier: metadata.Identifier,
        Sender:     metadata.sender,
        Receiver:   metadata.receiver,
        Msg: &pb.Message_ApplicationMessage{
            ApplicationMessage: &pb.ApplicationMessage{
                Type: pb.ApplicationMessage_ACK,
            },
        },
    }
}

func CreateMutexRequestMessage(metadata *Metadata, resource uint64, timestamp uint64) *pb.Message {
    return &pb.Message{
        Identifier: metadata.Identifier,
        Sender:     metadata.sender,
        // Broadcast -> no receiver necessary
        Msg: &pb.Message_MutexMessage{
            MutexMessage: &pb.MutexMessage{
                Type:      pb.MutexMessage_REQ,
                Resource:  resource,
                Timestamp: timestamp,
            },
        },
    }
}

func CreateMutexResponseMessage(metadata *Metadata, resource uint64) *pb.Message {
    return &pb.Message{
        Identifier: metadata.Identifier,
        Sender:     metadata.sender,
        Receiver:   metadata.receiver,
        Msg: &pb.Message_MutexMessage{
            MutexMessage: &pb.MutexMessage{
                Type:     pb.MutexMessage_RES,
                Resource: resource,
            },
        },
    }
}

func CreateExplorerMessage(sender uint64, initiator string) *pb.Message {
    return &pb.Message{
        Sender: sender,
        Msg: &pb.Message_ElectionMessage{
            ElectionMessage: &pb.Election{
                Type:      pb.Election_EXPLORER,
                Initiator: initiator,
            },
        },
    }
}

func CreateEchoMessage(sender uint64, initiator string) *pb.Message {
    return &pb.Message{
        Sender: sender,
        Msg: &pb.Message_ElectionMessage{
            ElectionMessage: &pb.Election{
                Type:      pb.Election_ECHO,
                Initiator: initiator,
            },
        },
    }
}
