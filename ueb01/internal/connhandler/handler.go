package connhandler

import (
    "github.com/akelsch/vaa/ueb01/api/pb"
    "github.com/akelsch/vaa/ueb01/internal/conf"
    "github.com/akelsch/vaa/ueb01/internal/errutil"
    "github.com/akelsch/vaa/ueb01/internal/pbutil"
    "google.golang.org/protobuf/proto"
    "io"
    "log"
    "net"
    "os"
)

type ConnectionHandler struct {
    bufferSize int
    config     *conf.Config
    directory  *conf.NeighborDirectory
}

func NewConnectionHandler(config *conf.Config) *ConnectionHandler {
    return &ConnectionHandler{
        bufferSize: 4 << 10, // 4KB
        config:     config,
        directory:  conf.NewNeighborDirectory(),
    }
}

func (handler *ConnectionHandler) HandleConnection(conn net.Conn) {
    defer conn.Close()

    buf := make([]byte, handler.bufferSize)
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
        handler.handleControlMessage(message.GetControlMessage())
    case *pb.Message_ApplicationMessage:
        handler.handleApplicationMessage(message.GetApplicationMessage())
    }
}

func (handler *ConnectionHandler) handleControlMessage(message *pb.ControlMessage) {
    log.Printf("Received control message: %s\n", message.Command)
    switch message.Command {
    case pb.ControlMessage_START:
        directory := handler.directory
        directory.Lock()
        directory.Reset()
        handler.sendToAllNeighbors()
        directory.Unlock()
    case pb.ControlMessage_EXIT:
        log.Println("Goodbye!")
        os.Exit(0)
    case pb.ControlMessage_EXIT_ALL:
        // TODO
    }
}

func (handler *ConnectionHandler) handleApplicationMessage(message *pb.ApplicationMessage) {
    log.Printf("Received application message: %s\n", message.Body)
    neighbors := handler.config.Neighbors
    directory := handler.directory
    for i := range neighbors {
        if message.Body == neighbors[i].Id {
            directory.Lock()
            if directory.HasAlreadyReceivedFrom(i) {
                directory.ResetIfNecessary(len(neighbors))
            }
            directory.SetReceived(i)
            handler.sendToAllNeighbors()
            directory.Unlock()
        }
    }
}

func (handler *ConnectionHandler) sendToAllNeighbors() {
    self := handler.config.Self
    neighbors := handler.config.Neighbors
    directory := handler.directory
    for i := range neighbors {
        if directory.HasNotSentTo(i) {
            neighbor := neighbors[i]
            conn, err := net.Dial("tcp", neighbor.GetDialAddress())
            if err != nil {
                log.Printf("Could not connect to node %s", neighbor.Id)
            } else {
                message, err := proto.Marshal(pbutil.CreateApplicationMessage(self.Id))
                errutil.HandleError(err)
                _, err = conn.Write(message)
                errutil.HandleError(err)
                directory.SetSent(i)
                log.Printf("Sent application message to node %s\n", neighbor.Id)
            }
        }
    }
}
