package handler

import (
    "github.com/akelsch/vaa/ueb02/api/pb"
    "github.com/akelsch/vaa/ueb02/internal/errutil"
    "github.com/akelsch/vaa/ueb02/internal/pbutil"
    "google.golang.org/protobuf/proto"
    "log"
    "net"
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
                if !h.dir.Status.Busy {
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
        statusDir := h.dir.Status

        ready := statusDir.AddStatus(message, len(h.conf.Neighbors))
        if ready {
            log.Println("------- DOUBLE COUNT DONE -------")
            statusDir.Ticker.Stop()

            _, selfReceived := h.dir.Neighbors.Stats()
            isValidCount := statusDir.CheckStatesAndNumberOfMessages(selfReceived)
            if isValidCount {
                log.Println("Double count is valid!")
                statusDir.GetAndPrintResults(h.conf.Params.T, h.conf.Self.Id)
                // TODO spread the result
            } else {
                log.Println("Double count is invalid! Restarting...")
                statusDir.Restart()
            }
        }
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
