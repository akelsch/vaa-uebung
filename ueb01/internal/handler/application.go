package handler

import (
    "github.com/akelsch/vaa/ueb01/api/pb"
    "github.com/akelsch/vaa/ueb01/internal/errutil"
    "github.com/akelsch/vaa/ueb01/internal/pbutil"
    "google.golang.org/protobuf/proto"
    "log"
    "net"
)

func (h *ConnectionHandler) handleApplicationMessage(message *pb.Message) {
    am := message.GetApplicationMessage()
    log.Printf("Received application message: %s\n", am.Body)

    for i := range h.conf.Neighbors {
        if am.Body == h.conf.Neighbors[i].Id {
            h.dir.Lock()
            if h.dir.Neighbors.HasAlreadyReceivedFrom(i) {
                h.dir.Neighbors.ResetIfNecessary(len(h.conf.Neighbors))
            }
            h.dir.Neighbors.SetReceived(i)
            h.propagateIdToNeighbors()
            h.dir.Unlock()
        }
    }
}

func (h *ConnectionHandler) propagateIdToNeighbors() {
    for i := range h.conf.Neighbors {
        if h.dir.Neighbors.HasNotSentTo(i) {
            neighbor := h.conf.Neighbors[i]
            conn, err := net.Dial("tcp", neighbor.GetDialAddress())
            if err != nil {
                log.Printf("Could not connect to node %s", neighbor.Id)
            } else {
                bytes, err := proto.Marshal(pbutil.CreateApplicationMessage(h.conf.Self.Id))
                errutil.HandleError(err)
                _, err = conn.Write(bytes)
                errutil.HandleError(err)
                conn.Close()
                h.dir.Neighbors.SetSent(i)
                log.Printf("Sent message to node %s\n", neighbor.Id)
            }
        }
    }
}
