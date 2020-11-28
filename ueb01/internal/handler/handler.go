package handler

import (
    "github.com/akelsch/vaa/ueb01/api/pb"
    "github.com/akelsch/vaa/ueb01/internal/config"
    "github.com/akelsch/vaa/ueb01/internal/directory"
    "github.com/akelsch/vaa/ueb01/internal/errutil"
    "google.golang.org/protobuf/proto"
    "io"
    "log"
    "net"
    "sync"
)

type ConnectionHandler struct {
    mu   sync.Mutex
    ln   *net.Listener
    Quit chan interface{}
    conf *config.Config
    dir  *directory.NeighborDirectory
}

func NewConnectionHandler(ln *net.Listener, conf *config.Config) *ConnectionHandler {
    return &ConnectionHandler{
        ln:   ln,
        Quit: make(chan interface{}),
        conf: conf,
        dir:  directory.NewNeighborDirectory(),
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
        h.handleControlMessage(message.GetControlMessage())
    case *pb.Message_ApplicationMessage:
        h.handleApplicationMessage(message.GetApplicationMessage())
    }
}

func (h *ConnectionHandler) sendToRemainingNeighbors(message *pb.Message) {
    for i := range h.conf.Neighbors {
        if h.dir.HasNotSentTo(i) {
            neighbor := h.conf.Neighbors[i]
            conn, err := net.Dial("tcp", neighbor.GetDialAddress())
            if err != nil {
                log.Printf("Could not connect to node %s", neighbor.Id)
            } else {
                bytes, err := proto.Marshal(message)
                errutil.HandleError(err)
                _, err = conn.Write(bytes)
                errutil.HandleError(err)
                h.dir.SetSent(i)
                log.Printf("Sent message to node %s\n", neighbor.Id)
            }
        }
    }
}
