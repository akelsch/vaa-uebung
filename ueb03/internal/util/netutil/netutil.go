package netutil

import (
    "github.com/akelsch/vaa/ueb03/api/pb"
    "github.com/akelsch/vaa/ueb03/internal/util/errutil"
    "google.golang.org/protobuf/proto"
    "log"
    "net"
)

func SendMessage(address string, message *pb.Message, successLog string) bool {
    conn, err := net.Dial("tcp", address)
    if err != nil {
        log.Printf("Could not connect to address %s\n", address)
        return false
    }

    bytes, err := proto.Marshal(message)
    errutil.HandleError(err)

    _, err = conn.Write(bytes)
    errutil.HandleError(err)
    conn.Close()

    log.Println(successLog)
    return true
}

func SendMessageIgnoringErrors(address string, message *pb.Message, successLog string) {
    conn, err := net.Dial("tcp", address)
    if err == nil {
        bytes, err := proto.Marshal(message)
        errutil.HandleError(err) // still handle protobuf serialization errors

        _, err = conn.Write(bytes)
        conn.Close()

        if err == nil {
            log.Println(successLog)
        }
    }
}
