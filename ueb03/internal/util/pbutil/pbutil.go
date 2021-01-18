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

func CreateApplicationStartMessage(sender string, body int) *pb.Message {
    return &pb.Message{
        Sender: sender,
        Msg: &pb.Message_ApplicationMessage{
            ApplicationMessage: &pb.ApplicationMessage{
                Type: pb.ApplicationMessage_START,
                Body: int32(body),
            },
        },
    }
}

func CreateApplicationAckMessage(sender string, body int) *pb.Message {
    return &pb.Message{
        Sender: sender,
        Msg: &pb.Message_ApplicationMessage{
            ApplicationMessage: &pb.ApplicationMessage{
                Type: pb.ApplicationMessage_ACK,
                Body: int32(body),
            },
        },
    }
}

func CreateApplicationResultMessage(sender string, body int) *pb.Message {
    return &pb.Message{
        Sender: sender,
        Msg: &pb.Message_ApplicationMessage{
            ApplicationMessage: &pb.ApplicationMessage{
                Type: pb.ApplicationMessage_RESULT,
                Body: int32(body),
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
