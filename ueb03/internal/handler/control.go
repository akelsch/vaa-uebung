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
        h.handleStart(message)
    case pb.ControlMessage_EXIT:
        h.handleExit()
    case pb.ControlMessage_EXIT_ALL:
        h.handleExitAll(message)
    }
}

func (h *ConnectionHandler) handleStart(message *pb.Message) {
    identifier := message.GetIdentifier()

    h.dir.Lock()
    defer h.dir.Unlock()

    if !h.dir.Flooding.IsHandled(identifier) {
        h.forwardMessage(message)
        log.Println("Starting...")
        h.startFirstStep()
    }
}

func (h *ConnectionHandler) handleExit() {
    close(h.quit)
    (*h.ln).Close()
}

func (h *ConnectionHandler) handleExitAll(message *pb.Message) {
    sender := message.GetSender()

    select {
    case <-h.quit:
        // Already exiting, ignore
    default:
        close(h.quit)
        for _, neighbor := range h.conf.Neighbors {
            if neighbor.Id != sender {
                address := neighbor.GetDialAddress()
                metadata := pbutil.CreateMetadata(h.conf.Self.Id, 0, 0)
                message := pbutil.CreateControlMessage(metadata, pb.ControlMessage_EXIT_ALL)
                successLog := fmt.Sprintf("Propagated exit to node %d", neighbor.Id)
                netutil.SendMessageIgnoringErrors(address, message, successLog)
            }
        }
        (*h.ln).Close()
    }
}
