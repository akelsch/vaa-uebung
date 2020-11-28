package handler

import (
    "github.com/akelsch/vaa/ueb01/api/pb"
    "github.com/akelsch/vaa/ueb01/internal/pbutil"
    "log"
)

func (h *ConnectionHandler) handleControlMessage(message *pb.ControlMessage) {
    log.Printf("Received control message: %s\n", message.Command)
    switch message.Command {
    case pb.ControlMessage_START:
        h.dir.Lock()
        h.dir.Reset()
        h.sendToRemainingNeighbors(pbutil.CreateApplicationMessage(h.conf.Self.Id))
        h.dir.Unlock()
    case pb.ControlMessage_EXIT:
        close(h.Quit)
        (*h.ln).Close()
    case pb.ControlMessage_EXIT_ALL:
        //fmt.Println("received exit")
        //h.sendExitToAllNeighbors()
        //fmt.Println("lock")
        //h.mu.Lock()
        //defer func() {
        //    fmt.Println("unlock")
        //    h.mu.Unlock()
        //}()
        //select {
        //case <-h.ch:
        //    // Already closed. Don't close again.
        //default:
        //    // Safe to close here. We're the only closer, guarded by mutex.
        //    close(h.ch)
        //    (*h.ln).Close()
        //}
    }
}
