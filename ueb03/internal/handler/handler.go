package handler

import (
    "github.com/akelsch/vaa/ueb03/api/pb"
    "github.com/akelsch/vaa/ueb03/internal/config"
    "github.com/akelsch/vaa/ueb03/internal/directory"
    "github.com/akelsch/vaa/ueb03/internal/util/errutil"
    "github.com/akelsch/vaa/ueb03/internal/util/netutil"
    "github.com/akelsch/vaa/ueb03/internal/util/randutil"
    "google.golang.org/protobuf/proto"
    "io"
    "log"
    "net"
)

type ConnectionHandler struct {
    ln   *net.Listener
    quit chan interface{}
    conf *config.Config
    dir  *directory.Directory
}

func NewConnectionHandler(ln *net.Listener, conf *config.Config) *ConnectionHandler {
    return &ConnectionHandler{
        ln:   ln,
        quit: make(chan interface{}),
        conf: conf,
        dir:  directory.NewDirectory(conf),
    }
}

func (h *ConnectionHandler) HandleError(err error) {
    select {
    case <-h.quit:
        // Listener has been closed by a goroutine
        log.Println("Goodbye!")
    default:
        errutil.HandleError(err)
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
    case *pb.Message_MutexMessage:
        h.handleMutexMessage(message)
    case *pb.Message_ElectionMessage:
        h.handleElectionMessage(message)
    case *pb.Message_SnapshotMessage:
        h.handleSnapshotMessage(message)
    }
}

func (h *ConnectionHandler) StartElection() {
    if randutil.RandomBool() {
        h.handleStartElection()
    }
}

func (h *ConnectionHandler) forwardMessage(message *pb.Message) {
    for _, neighbor := range h.conf.Neighbors {
        address := neighbor.GetDialAddress()
        //successLog := fmt.Sprintf("Forwarded message '%s' to node %d", message.GetIdentifier(), neighbor.Id)
        netutil.SendMessageSilently(address, message)
    }
}

func (h *ConnectionHandler) unicastMessage(node *config.Node, message *pb.Message, successLog string) {
    if h.conf.IsNodeNeighbor(node.Id) {
        // Direct message
        address := node.GetDialAddress()
        netutil.SendMessage(address, message, successLog)
    } else {
        // Flood neighbors
        log.Println(successLog, "*")
        for _, neighbor := range h.conf.Neighbors {
            address := neighbor.GetDialAddress()
            netutil.SendMessageSilently(address, message)
        }
    }
}
