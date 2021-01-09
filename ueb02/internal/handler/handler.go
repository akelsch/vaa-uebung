package handler

import (
    "github.com/akelsch/vaa/ueb02/api/pb"
    "github.com/akelsch/vaa/ueb02/internal/config"
    "github.com/akelsch/vaa/ueb02/internal/directory"
    "github.com/akelsch/vaa/ueb02/internal/errutil"
    "google.golang.org/protobuf/proto"
    "io"
    "log"
    "net"
)

type ConnectionHandler struct {
    ln   *net.Listener
    Quit chan interface{}
    conf *config.Config
    dir  *directory.MessageDirectory
}

func NewConnectionHandler(ln *net.Listener, conf *config.Config) *ConnectionHandler {
    return &ConnectionHandler{
        ln:   ln,
        Quit: make(chan interface{}),
        conf: conf,
        dir:  directory.NewMessageDirectory(),
    }
}

func (h *ConnectionHandler) HandleConnection(conn net.Conn) {
    defer conn.Close()

    buf := make([]byte, 4<<10) // 4KB
    n, err := conn.Read(buf)
    if err == io.EOF {
        return
    } else {
        errutil.HandleError(err)
    }

    message := &pb.Message{}
    err = proto.Unmarshal(buf[:n], message)
    if err != nil {
        log.Println("Could not parse protobuf message")
        return
    }

    switch message.Msg.(type) {
    case *pb.Message_ControlMessage:
        h.handleControlMessage(message)
    case *pb.Message_ApplicationMessage:
        h.handleApplicationMessage(message)
    case *pb.Message_Election:
        h.handleElectionMessage(message)
    case *pb.Message_Status:
        h.handleStatusMessage(message)
    }
}
