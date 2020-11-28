package handler

import (
    "github.com/akelsch/vaa/ueb01/api/pb"
    "github.com/akelsch/vaa/ueb01/internal/pbutil"
    "log"
)

func (h *ConnectionHandler) handleApplicationMessage(message *pb.ApplicationMessage) {
    log.Printf("Received application message: %s\n", message.Body)
    for i := range h.conf.Neighbors {
        if message.Body == h.conf.Neighbors[i].Id {
            h.dir.Lock()
            if h.dir.HasAlreadyReceivedFrom(i) {
                h.dir.ResetIfNecessary(len(h.conf.Neighbors))
            }
            h.dir.SetReceived(i)
            h.sendToRemainingNeighbors(pbutil.CreateApplicationMessage(h.conf.Self.Id))
            h.dir.Unlock()
        }
    }
}
