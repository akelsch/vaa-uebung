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
        message := pbutil.CreateApplicationStartMessage(h.conf.Self.Id, h.conf.Params.T)
        successMessage := fmt.Sprintf("Sent application message to node %s: t = %d", neighbor.Id, h.conf.Params.T)

        if netutil.SendMessage(address, message, successMessage) {
            h.dir.Neighbors.SetSent(i)
        }
    }

    h.dir.Status.Busy = false
}

func (h *ConnectionHandler) handleApplicationMessage(message *pb.Message) {
    h.dir.Status.Busy = true

    am := message.GetApplicationMessage()
    log.Printf("Received application message from %s\n", message.GetSender())

    if i, neighbor := h.conf.FindNeighborById(message.GetSender()); neighbor != nil {
        h.dir.Neighbors.SetReceived(i)
    }

    switch am.Type {
    case pb.ApplicationMessage_START:
        if h.conf.Params.AMax > 0 {
            h.conf.Params.AMax--
            h.recalculateVotedTime(int(am.GetBody()))

            if i, neighbor := h.conf.FindNeighborById(message.GetSender()); neighbor != nil {
                address := neighbor.GetDialAddress()
                message := pbutil.CreateApplicationAckMessage(h.conf.Self.Id, h.conf.Params.T)
                successMessage := fmt.Sprintf("Acknowledged application message of node %s", neighbor.Id)
                if netutil.SendMessage(address, message, successMessage) {
                    h.dir.Neighbors.SetSent(i)
                }
            }
        } else {
            log.Println("Ignoring vote, AMax reached")
        }
    case pb.ApplicationMessage_ACK:
        h.recalculateVotedTime(int(am.GetBody()))
    case pb.ApplicationMessage_RESULT:
        if am.GetBody() != 0 {
            log.Printf("The vote result is t = %d\n", am.GetBody())
        } else {
            log.Println("The vote result is indecision")
        }
    }

    h.dir.Status.Busy = false
}

func (h *ConnectionHandler) recalculateVotedTime(received int) {
    oldT := h.conf.Params.T
    h.conf.Params.T = int(math.Ceil(float64(oldT+received) / 2))
    log.Printf("Neighbor t = %d, Old t = %d, New t = %d\n", received, oldT, h.conf.Params.T)
}
