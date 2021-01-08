package handler

import (
    "github.com/akelsch/vaa/ueb02/api/pb"
    "github.com/akelsch/vaa/ueb02/internal/errutil"
    "github.com/akelsch/vaa/ueb02/internal/pbutil"
    "google.golang.org/protobuf/proto"
    "log"
    "net"
    "runtime"
)

func (h *ConnectionHandler) handleGetStatus() {
    for _, neighbor := range h.conf.Neighbors {
        if neighbor.Id == h.dir.Election.Predecessor {
            conn, err := net.Dial("tcp", neighbor.GetDialAddress())
            if err != nil {
                log.Printf("Could not connect to node %s\n", neighbor.Id)
            } else {
                // status message params
                state := pb.Status_ACTIVE
                if runtime.NumGoroutine() == 2 {
                    state = pb.Status_PASSIVE
                }
                sent, received := h.dir.Neighbors.Stats()

                bytes, err := proto.Marshal(pbutil.CreateStatusMessage(h.conf.Self.Id, state, sent, received, h.conf.Params.T))
                errutil.HandleError(err)
                _, err = conn.Write(bytes)
                errutil.HandleError(err)
                conn.Close()
                log.Printf("Sent status message to node %s\n", neighbor.Id)
            }
        }
    }
}

func (h *ConnectionHandler) handleStatusMessage(message *pb.Message) {
    if h.dir.Election.IsCoordinator(h.conf.Self.Id) {
        log.Printf("Coordinator got status from %s\n", message.GetSender())
    } else {
        for _, neighbor := range h.conf.Neighbors {
            if neighbor.Id == h.dir.Election.Predecessor {
                conn, err := net.Dial("tcp", neighbor.GetDialAddress())
                if err != nil {
                    log.Printf("Could not connect to node %s\n", neighbor.Id)
                } else {
                    bytes, err := proto.Marshal(pbutil.CloneStatusMessage(message))
                    errutil.HandleError(err)
                    _, err = conn.Write(bytes)
                    errutil.HandleError(err)
                    conn.Close()
                    log.Printf("Forwarded status message from node %s to node %s\n", message.GetSender(), neighbor.Id)
                }
            }
        }
    }
}
