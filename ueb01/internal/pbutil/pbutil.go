package pbutil

import "github.com/akelsch/vaa/ueb01/api/pb"

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

func CreateApplicationMessage(body string) *pb.Message {
    return &pb.Message{
        Msg: &pb.Message_ApplicationMessage{
            ApplicationMessage: &pb.ApplicationMessage{
                Body: body,
            },
        },
    }
}
