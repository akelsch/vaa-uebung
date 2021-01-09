package handler

import (
    "github.com/akelsch/vaa/ueb02/api/pb"
    "github.com/akelsch/vaa/ueb02/internal/util/errutil"
    "github.com/akelsch/vaa/ueb02/internal/util/pbutil"
    "google.golang.org/protobuf/proto"
    "log"
    "net"
)

func (h *ConnectionHandler) handleControlMessage(message *pb.Message) {
    cm := message.GetControlMessage()
    log.Printf("Received control message: %s\n", cm.Command)

    switch cm.Command {
    case pb.ControlMessage_START:
        h.exchangeTimeWithNeighbors()
    case pb.ControlMessage_EXIT:
        close(h.Quit)
        (*h.ln).Close()
    case pb.ControlMessage_EXIT_ALL:
        select {
        case <-h.Quit:
            // Already exiting, ignore
        default:
            close(h.Quit)
            h.propagateExitToNeighbors(message.Sender)
            (*h.ln).Close()
        }
    case pb.ControlMessage_START_ELECTION:
        h.dir.Lock()
        h.handleStartElection()
        h.dir.Unlock()
    case pb.ControlMessage_GET_STATUS:
        h.handleGetStatus()
    }
}

func (h *ConnectionHandler) propagateExitToNeighbors(sender string) {
    for _, neighbor := range h.conf.Neighbors {
        if neighbor.Id != sender {
            conn, err := net.Dial("tcp", neighbor.GetDialAddress())
            if err == nil {
                bytes, err := proto.Marshal(pbutil.CreateControlMessage(h.conf.Self.Id, pb.ControlMessage_EXIT_ALL))
                errutil.HandleError(err)
                _, err = conn.Write(bytes)
                conn.Close()
                // Ignore write errors as other node could have exited
                if err == nil {
                    log.Printf("Propagated exit to node %s\n", neighbor.Id)
                }
            }
        }
    }
}
