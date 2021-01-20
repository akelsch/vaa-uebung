package handler

import (
    "fmt"
    "github.com/akelsch/vaa/ueb03/api/pb"
    "github.com/akelsch/vaa/ueb03/internal/config"
    "github.com/akelsch/vaa/ueb03/internal/util/netutil"
    "github.com/akelsch/vaa/ueb03/internal/util/pbutil"
    "github.com/akelsch/vaa/ueb03/internal/util/randutil"
    "log"
    "math"
)

func (h *ConnectionHandler) handleStart2() { // TODO
    node := h.conf.FindRandomNode()

    // Step 4
    percent := randutil.RoundedRandomUint(0, 100, 10)

    // Step 6 (swapped places with 5)
    metadata := pbutil.CreateMetadata(h.conf.Self.Id, node.Id, h.dir.Flooding.NextSequence())
    message := pbutil.CreateApplicationRequestMessage(metadata, percent)
    successLog := fmt.Sprintf("Sent balance request to node %d", node.Id)
    h.floodMessage(node, message, successLog)
}

func (h *ConnectionHandler) handleApplicationMessage(message *pb.Message) {
    am := message.GetApplicationMessage()
    identifier := message.GetIdentifier()
    sender := message.GetSender()
    receiver := message.GetReceiver()

    if !h.dir.Flooding.IsHandled(identifier) {
        h.dir.Flooding.MarkAsHandled(identifier)

        if receiver != h.conf.Self.Id {
            h.forwardMessage(message)
        } else {
            switch am.Type {
            case pb.ApplicationMessage_NUL:
                log.Printf("Received application message from node %d\n", sender)
                h.handleApplicationDefault(am, sender)
            case pb.ApplicationMessage_REQ:
                log.Printf("Received balance request from node %d\n", sender)
                h.handleApplicationRequest(am, sender)
            case pb.ApplicationMessage_RES:
                log.Printf("Received balance response from node %d\n", sender)
                h.handleApplicationResponse(am, sender)
            case pb.ApplicationMessage_ACK:
                log.Printf("Received acknowledgment from node %d\n", sender)
                h.handleApplicationAcknowledgment()
            }
        }
    } else {
        log.Printf("Message '%s' got handled already\n", identifier)
    }
}

func (h *ConnectionHandler) handleApplicationDefault(am *pb.ApplicationMessage, sender uint64) {
    p := am.GetPercent()
    bi := am.GetBalance()
    bj := h.conf.Params.Balance

    // Step 8
    if bi >= bj {
        plus := int64(math.Round(float64(bi) * (float64(p) / 100)))
        h.conf.Params.Balance = bj + plus
        log.Printf("Increasing balance by %d: Old = %d, New = %d\n", plus, bj, h.conf.Params.Balance)
    } else {
        minus := 1 - (float64(p) / 100)
        h.conf.Params.Balance = int64(math.Round(float64(bj) * minus))
        log.Printf("Decreasing balance by %d percent: Old = %d, New = %d\n", p, bj, h.conf.Params.Balance)
    }

    node := h.conf.FindNodeById(sender)
    metadata := pbutil.CreateMetadata(h.conf.Self.Id, node.Id, h.dir.Flooding.NextSequence())
    message := pbutil.CreateApplicationAcknowledgmentMessage(metadata)
    successLog := fmt.Sprintf("Sent acknowledgment to node %d", node.Id)
    h.floodMessage(node, message, successLog)
}

func (h *ConnectionHandler) handleApplicationRequest(am *pb.ApplicationMessage, sender uint64) {
    node := h.conf.FindNodeById(sender)
    percent := am.GetPercent()

    metadata := pbutil.CreateMetadata(h.conf.Self.Id, node.Id, h.dir.Flooding.NextSequence())
    message := pbutil.CreateApplicationResponseMessage(metadata, h.conf.Params.Balance, percent)
    successLog := fmt.Sprintf("Sent balance response to node %d: B = %d", node.Id, h.conf.Params.Balance)
    h.floodMessage(node, message, successLog)
}

func (h *ConnectionHandler) handleApplicationResponse(am *pb.ApplicationMessage, sender uint64) {
    p := am.GetPercent()
    bi := h.conf.Params.Balance
    bj := am.GetBalance()

    // Step 5 (swapped places with 6)
    node := h.conf.FindNodeById(sender)
    metadata := pbutil.CreateMetadata(h.conf.Self.Id, node.Id, h.dir.Flooding.NextSequence())
    message := pbutil.CreateApplicationMessage(metadata, h.conf.Params.Balance, p)
    successLog := fmt.Sprintf("Sent application message to node %d: B = %d, p = %d", node.Id, h.conf.Params.Balance, p)
    h.floodMessage(node, message, successLog)

    // Step 7
    if bj >= bi {
        plus := int64(math.Round(float64(bj) * (float64(p) / 100)))
        h.conf.Params.Balance = bi + plus
        log.Printf("Increasing balance by %d: Old = %d, New = %d\n", plus, bi, h.conf.Params.Balance)
    } else {
        minus := 1 - (float64(p) / 100)
        h.conf.Params.Balance = int64(math.Round(float64(bi) * minus))
        log.Printf("Decreasing balance by %d percent: Old = %d, New = %d\n", p, bi, h.conf.Params.Balance)
    }
}

func (h *ConnectionHandler) handleApplicationAcknowledgment() {
    // TODO unlock
}

func (h *ConnectionHandler) forwardMessage(message *pb.Message) {
    for _, neighbor := range h.conf.Neighbors {
        address := neighbor.GetDialAddress()
        successLog := fmt.Sprintf("Forwarded message '%s' to node %d", message.GetIdentifier(), neighbor.Id)
        netutil.SendMessage(address, message, successLog)
    }
}

func (h *ConnectionHandler) floodMessage(node *config.Node, message *pb.Message, successLog string) {
    if h.conf.IsNodeNeighbor(node.Id) {
        // Direct message
        address := node.GetDialAddress()
        netutil.SendMessage(address, message, successLog)
    } else {
        // Flood neighbors
        log.Printf("*** Flooding neighbors with message '%s' as node %d is not a neighbor ***\n", message.GetIdentifier(), node.Id)
        log.Println(successLog)
        for _, neighbor := range h.conf.Neighbors {
            address := neighbor.GetDialAddress()
            netutil.SendMessageSilently(address, message)
        }
    }
}
