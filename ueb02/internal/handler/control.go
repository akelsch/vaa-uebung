package handler

import (
    "fmt"
    "github.com/akelsch/vaa/ueb02/api/pb"
    "github.com/akelsch/vaa/ueb02/internal/util/netutil"
    "github.com/akelsch/vaa/ueb02/internal/util/pbutil"
    "log"
)

func (h *ConnectionHandler) handleControlMessage(message *pb.Message) {
    cm := message.GetControlMessage()
    log.Printf("Received control message: %s\n", cm.Command)

    switch cm.Command {
    case pb.ControlMessage_START:
        h.handleStartCommand()
    case pb.ControlMessage_EXIT:
        h.handleExitCommand()
    case pb.ControlMessage_EXIT_ALL:
        h.handleExitAllCommand(message.GetSender())
    case pb.ControlMessage_START_ELECTION:
        h.handleStartElectionCommand()
    case pb.ControlMessage_GET_STATUS:
        h.handleGetStatusCommand()
    }
}

func (h *ConnectionHandler) handleStartCommand() {
    h.exchangeTimeWithNeighbors()
}

func (h *ConnectionHandler) handleExitCommand() {
    close(h.quit)
    (*h.ln).Close()
}

func (h *ConnectionHandler) handleExitAllCommand(sender string) {
    select {
    case <-h.quit:
        // Already exiting, ignore
    default:
        close(h.quit)
        for _, neighbor := range h.conf.Neighbors {
            if neighbor.Id != sender {
                address := neighbor.GetDialAddress()
                message := pbutil.CreateControlMessage(h.conf.Self.Id, pb.ControlMessage_EXIT_ALL)
                successMessage := fmt.Sprintf("Propagated exit to node %s", neighbor.Id)
                netutil.SendMessageIgnoringErrors(address, message, successMessage)
            }
        }
        (*h.ln).Close()
    }
}

func (h *ConnectionHandler) handleStartElectionCommand() {
    h.handleStartElection()
}

func (h *ConnectionHandler) handleGetStatusCommand() {
    h.handleGetStatus()
}
