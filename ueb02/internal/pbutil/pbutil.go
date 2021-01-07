package pbutil

import "github.com/akelsch/vaa/ueb02/api/pb"

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

func CreateApplicationMessage(body int) *pb.Message {
    return &pb.Message{
        Msg: &pb.Message_ApplicationMessage{
            ApplicationMessage: &pb.ApplicationMessage{
                Body: int32(body),
            },
        },
    }
}

func CreateRumorMessage(sender string, rumor *pb.Rumor) *pb.Message {
    return &pb.Message{
        Sender: sender,
        Msg: &pb.Message_Rumor{
            Rumor: rumor,
        },
    }
}

func CreateExplorerMessage(sender string, initiator string) *pb.Message {
    return &pb.Message{
        Sender: sender,
        Msg: &pb.Message_Election{
            Election: &pb.Election{
                Type: pb.Election_EXPLORER,
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
                Type: pb.Election_ECHO,
                Initiator: initiator,
            },
        },
    }
}
