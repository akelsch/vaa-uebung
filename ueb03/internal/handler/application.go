package handler

import (
    "fmt"
    "github.com/akelsch/vaa/ueb03/api/pb"
    "github.com/akelsch/vaa/ueb03/internal/util/netutil"
    "github.com/akelsch/vaa/ueb03/internal/util/pbutil"
    "github.com/akelsch/vaa/ueb03/internal/util/randutil"
    "log"
    "math"
)

// TODO use flooding

func (h *ConnectionHandler) handleStart() {
    node := h.conf.FindRandomNode()

    // Step 4
    percent := randutil.RoundedRandomUint(0, 100, 10)

    // Step 6 (swapped places with 5)
    address := node.GetDialAddress()
    metadata := pbutil.CreateMetadata(h.conf.Self.Id, node.Id, 0)
    message := pbutil.CreateApplicationRequestMessage(metadata, percent)
    successMessage := fmt.Sprintf("Sent balance request to node %d", node.Id)
    netutil.SendMessage(address, message, successMessage)
}

func (h *ConnectionHandler) handleApplicationMessage(message *pb.Message) {
    am := message.GetApplicationMessage()
    sender := message.GetSender()

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
    address := node.GetDialAddress()
    metadata := pbutil.CreateMetadata(h.conf.Self.Id, node.Id, 0)
    message := pbutil.CreateApplicationAcknowledgmentMessage(metadata)
    successMessage := fmt.Sprintf("Sent acknowledgment to node %d", node.Id)
    netutil.SendMessage(address, message, successMessage)
}

func (h *ConnectionHandler) handleApplicationRequest(am *pb.ApplicationMessage, sender uint64) {
    node := h.conf.FindNodeById(sender)
    percent := am.GetPercent()

    address := node.GetDialAddress()
    metadata := pbutil.CreateMetadata(h.conf.Self.Id, node.Id, 0)
    message := pbutil.CreateApplicationResponseMessage(metadata, h.conf.Params.Balance, percent)
    successMessage := fmt.Sprintf("Sent balance response to node %d: B = %d", node.Id, h.conf.Params.Balance)
    netutil.SendMessage(address, message, successMessage)
}

func (h *ConnectionHandler) handleApplicationResponse(am *pb.ApplicationMessage, sender uint64) {
    p := am.GetPercent()
    bi := h.conf.Params.Balance
    bj := am.GetBalance()

    // Step 5 (swapped places with 6)
    node := h.conf.FindNodeById(sender)
    address := node.GetDialAddress()
    metadata := pbutil.CreateMetadata(h.conf.Self.Id, node.Id, 0)
    message := pbutil.CreateApplicationMessage(metadata, h.conf.Params.Balance, p)
    successMessage := fmt.Sprintf("Sent application message to node %d: B = %d, p = %d", node.Id, h.conf.Params.Balance, p)
    netutil.SendMessage(address, message, successMessage)

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
