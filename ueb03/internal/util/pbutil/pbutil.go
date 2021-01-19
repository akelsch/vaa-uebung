package pbutil

import "github.com/akelsch/vaa/ueb03/api/pb"

func CreateControlMessage(sender string, command pb.ControlMessage_Command) *pb.Message {
    return &pb.Message{
        Sender: sender,
        Msg: &pb.Message_ControlMessage{
            ControlMessage: &pb.ControlMessage{
                Command: command,
            },
        },
    }
}

func CreateApplicationMessage(sender string, balance, percent int) *pb.Message {
    return &pb.Message{
        Sender: sender,
        Msg: &pb.Message_ApplicationMessage{
            ApplicationMessage: &pb.ApplicationMessage{
                Balance: int32(balance),
                Percent: int32(percent),
            },
        },
    }
}

func CreateApplicationRequestMessage(sender string, percent int) *pb.Message {
    return &pb.Message{
        Sender: sender,
        Msg: &pb.Message_ApplicationMessage{
            ApplicationMessage: &pb.ApplicationMessage{
                Type:    pb.ApplicationMessage_REQ,
                Percent: int32(percent),
            },
        },
    }
}

func CreateApplicationResponseMessage(sender string, balance, percent int) *pb.Message {
    return &pb.Message{
        Sender: sender,
        Msg: &pb.Message_ApplicationMessage{
            ApplicationMessage: &pb.ApplicationMessage{
                Type:    pb.ApplicationMessage_RES,
                Balance: int32(balance),
                Percent: int32(percent),
            },
        },
    }
}

func CreateExplorerMessage(sender string, initiator string) *pb.Message {
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

func CreateEchoMessage(sender string, initiator string) *pb.Message {
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
