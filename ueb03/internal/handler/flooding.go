package handler

import (
    "fmt"
    "github.com/akelsch/vaa/ueb03/api/pb"
    "github.com/akelsch/vaa/ueb03/internal/config"
    "github.com/akelsch/vaa/ueb03/internal/util/netutil"
    "log"
)

func (h *ConnectionHandler) forwardMessage(message *pb.Message) {
    for _, neighbor := range h.conf.Neighbors {
        address := neighbor.GetDialAddress()
        successLog := fmt.Sprintf("Forwarded message '%s' to node %d", message.GetIdentifier(), neighbor.Id)
        netutil.SendMessage(address, message, successLog)
    }
}

func (h *ConnectionHandler) unicastMessage(node *config.Node, message *pb.Message, successLog string) {
    if h.conf.IsNodeNeighbor(node.Id) {
        // Direct message
        address := node.GetDialAddress()
        netutil.SendMessage(address, message, successLog)
    } else {
        // Flood neighbors
        log.Println(successLog, "*")
        for _, neighbor := range h.conf.Neighbors {
            address := neighbor.GetDialAddress()
            netutil.SendMessageSilently(address, message)
        }
    }
}
