package handler

import (
    "fmt"
    "github.com/akelsch/vaa/ueb02/api/pb"
    "github.com/akelsch/vaa/ueb02/internal/util/netutil"
    "github.com/akelsch/vaa/ueb02/internal/util/pbutil"
    "log"
)

func (h *ConnectionHandler) handleGetStatus() {
    for _, neighbor := range h.conf.Neighbors {
        if neighbor.Id == h.dir.Election.Predecessor { // TODO
            // status message params
            state := pb.Status_ACTIVE
            if !h.dir.Status.Busy {
                state = pb.Status_PASSIVE
            }
            sent, received := h.dir.Neighbors.Stats()

            address := neighbor.GetDialAddress()
            message := pbutil.CreateStatusMessage(h.conf.Self.Id, state, sent, received, h.conf.Params.T)
            successMessage := fmt.Sprintf("Sent status message to node %s", neighbor.Id)
            netutil.SendMessage(address, message, successMessage)
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
                address := neighbor.GetDialAddress()
                message := pbutil.CloneStatusMessage(message)
                successMessage := fmt.Sprintf("Forwarded status message from node %s to node %s", message.GetSender(), neighbor.Id)
                netutil.SendMessage(address, message, successMessage)
            }
        }
    }
}
