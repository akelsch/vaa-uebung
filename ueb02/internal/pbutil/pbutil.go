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

func CreateApplicationMessage(sender string, body int) *pb.Message {
    return &pb.Message{
        Sender: sender,
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

func CreateStatusMessage(sender string, state pb.Status_State, sent int, received int, time int) *pb.Message {
    return &pb.Message{
        Sender: sender,
        Msg: &pb.Message_Status{
            Status: &pb.Status{
                State:    state,
                Sent:     int32(sent),
                Received: int32(received),
                Time:     int32(time),
            },
        },
    }
}

func CloneStatusMessage(message *pb.Message) *pb.Message {
    return &pb.Message{
        Sender: message.GetSender(),
        Msg: &pb.Message_Status{
            Status: message.GetStatus(),
        },
    }
}
