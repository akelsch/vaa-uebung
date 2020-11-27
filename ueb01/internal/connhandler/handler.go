package connhandler

import (
    "github.com/akelsch/vaa/ueb01/internal/conf"
    "github.com/akelsch/vaa/ueb01/internal/errutil"
    "log"
    "net"
    "os"
    "strings"
)

type ConnectionHandler struct {
    bufferSize    int
    controlPrefix string // TODO replace with protobufs
    config        *conf.Config
    directory     *conf.NeighborDirectory
}

func NewConnectionHandler(config *conf.Config) *ConnectionHandler {
    return &ConnectionHandler{
        bufferSize:    4 << 10, // 4KB
        controlPrefix: "c:",
        config:        config,
        directory:     conf.NewNeighborDirectory(),
    }
}

func (handler *ConnectionHandler) HandleConnection(conn net.Conn) {
    defer conn.Close()
    buf := make([]byte, handler.bufferSize)
    n, err := conn.Read(buf)
    errutil.HandleError(err)

    message := strings.TrimSpace(string(buf[:n]))
    log.Printf("Received: %s\n", message)

    if strings.HasPrefix(message, handler.controlPrefix) {
        handler.handleControlMessage(message)
    } else {
        handler.handleApplicationMessage(message)
    }
}

func (handler *ConnectionHandler) handleControlMessage(message string) {
    command := strings.TrimPrefix(message, handler.controlPrefix)
    if command == "start" {
        directory := handler.directory
        directory.Lock()
        directory.Reset()
        handler.sendToAllNeighbors()
        directory.Unlock()
    } else if command == "exit" {
        // TODO exit all
        os.Exit(0)
    }
}

func (handler *ConnectionHandler) handleApplicationMessage(message string) {
    neighbors := handler.config.Neighbors
    directory := handler.directory
    for i := range neighbors {
        if message == neighbors[i].Id {
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
                message := []byte(self.Id)
                _, err := conn.Write(message)
                errutil.HandleError(err)
                directory.SetSent(i)
                log.Printf("Sent %s to node %s\n", message, neighbor.Id)
            }
        }
    }
}
