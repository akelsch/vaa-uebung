package handler

import (
    "fmt"
    "github.com/akelsch/vaa/ueb02/api/pb"
    "github.com/akelsch/vaa/ueb02/internal/util/netutil"
    "github.com/akelsch/vaa/ueb02/internal/util/pbutil"
    "log"
    "math"
)

func (h *ConnectionHandler) handleStartVote() {
    h.dir.Status.Busy = true

    randomNeighbors := h.conf.GetRandomNeighbors(h.conf.Params.P)
    for i, neighbor := range randomNeighbors {
        address := neighbor.GetDialAddress()
        message := pbutil.CreateApplicationMessage(h.conf.Self.Id, h.conf.Params.T)
        successMessage := fmt.Sprintf("Sent application message to node %s: %d", neighbor.Id, h.conf.Params.T)

        if netutil.SendMessage(address, message, successMessage) {
            h.dir.Neighbors.SetSent(i)
        }
    }

    h.dir.Status.Busy = false
}

func (h *ConnectionHandler) handleApplicationMessage(message *pb.Message) {
    h.dir.Status.Busy = true
    am := message.GetApplicationMessage()
    t2 := int(am.GetBody())
    log.Printf("Received application message: %d\n", t2)

    // register received message for double counting stats
    if i, neighbor := h.conf.FindNeighborById(message.GetSender()); neighbor != nil {
        h.dir.Neighbors.SetReceived(i)
    }

    if h.conf.Params.AMax > 0 {
        h.conf.Params.AMax--
        oldT := h.conf.Params.T
        h.conf.Params.T = int(math.Ceil(float64(h.conf.Params.T+t2) / 2))
        log.Printf("Old t = %d, new t = %d\n", oldT, h.conf.Params.T)
    }
    h.dir.Status.Busy = false
}
