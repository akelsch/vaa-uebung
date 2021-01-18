package handler

import (
    "github.com/akelsch/vaa/ueb03/api/pb"
    "log"
)

func (h *ConnectionHandler) handleStart() {
    randomNode := h.conf.GetRandomNode()
    log.Printf("Chose node %s\n", randomNode.Id)
}

func (h *ConnectionHandler) handleApplicationMessage(message *pb.Message) {
    //am := message.GetApplicationMessage()
    log.Printf("Received application message from %s\n", message.GetSender())
}
