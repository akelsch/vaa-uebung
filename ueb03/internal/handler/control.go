package handler

import (
    "fmt"
    "github.com/akelsch/vaa/ueb03/api/pb"
    "github.com/akelsch/vaa/ueb03/internal/util/netutil"
    "github.com/akelsch/vaa/ueb03/internal/util/pbutil"
    "log"
)

func (h *ConnectionHandler) handleControlMessage(message *pb.Message) {
    cm := message.GetControlMessage()
    log.Printf("Received control message: %s\n", cm.Command)

    switch cm.Command {
    case pb.ControlMessage_START:
        h.handleStart()
    case pb.ControlMessage_EXIT:
        h.handleExit()
    case pb.ControlMessage_EXIT_ALL:
        h.handleExitAll(message.GetSender())
    }
}

func (h *ConnectionHandler) handleStart() {
    h.startFirstStep()
}

func (h *ConnectionHandler) handleExit() {
    close(h.quit)
    (*h.ln).Close()
}

func (h *ConnectionHandler) handleExitAll(sender uint64) {
    select {
    case <-h.quit:
        // Already exiting, ignore
    default:
        close(h.quit)
        for _, neighbor := range h.conf.Neighbors {
            if neighbor.Id != sender {
                address := neighbor.GetDialAddress()
                message := pbutil.CreateControlMessage(h.conf.Self.Id, pb.ControlMessage_EXIT_ALL)
                successLog := fmt.Sprintf("Propagated exit to node %d", neighbor.Id)
                netutil.SendMessageIgnoringErrors(address, message, successLog)
            }
        }
        (*h.ln).Close()
    }
}
