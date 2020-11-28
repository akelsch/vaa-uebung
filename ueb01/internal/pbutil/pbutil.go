package pbutil

import "github.com/akelsch/vaa/ueb01/api/pb"

func CreateApplicationMessage(body string) *pb.Message {
    return &pb.Message{
        Msg: &pb.Message_ApplicationMessage{
            ApplicationMessage: &pb.ApplicationMessage{
                Body: body,
            },
        },
    }
}
