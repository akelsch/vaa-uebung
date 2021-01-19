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
    _, node := h.conf.FindRandomNode()

    // Step 4
    percent := randutil.RoundedRandomInt(0, 100, 10)

    // Step 6 (swapped places with 5)
    address := node.GetDialAddress()
    message := pbutil.CreateApplicationRequestMessage(h.conf.Self.Id, percent)
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
    p := int(am.GetPercent())
    bi := int(am.GetBalance())
    bj := h.conf.Params.Balance

    // Step 8
    if bi >= bj {
        plus := int(math.Round(float64(bi) * (float64(p) / 100)))
        h.conf.Params.Balance = bj + plus
        log.Printf("Increasing balance by %d: Old = %d, New = %d\n", plus, bj, h.conf.Params.Balance)
    } else {
        minus := 1 - (float64(p) / 100)
        h.conf.Params.Balance = int(math.Round(float64(bj) * minus))
        log.Printf("Decreasing balance by %d percent: Old = %d, New = %d\n", p, bj, h.conf.Params.Balance)
    }

    _, node := h.conf.FindById(sender)
    address := node.GetDialAddress()
    message := pbutil.CreateApplicationAcknowledgmentMessage(h.conf.Self.Id)
    successMessage := fmt.Sprintf("Sent acknowledgment to node %d", node.Id)
    netutil.SendMessage(address, message, successMessage)
}

func (h *ConnectionHandler) handleApplicationRequest(am *pb.ApplicationMessage, sender uint64) {
    _, node := h.conf.FindById(sender)
    percent := int(am.GetPercent())

    address := node.GetDialAddress()
    message := pbutil.CreateApplicationResponseMessage(h.conf.Self.Id, h.conf.Params.Balance, percent)
    successMessage := fmt.Sprintf("Sent balance response to node %d: B = %d", node.Id, h.conf.Params.Balance)
    netutil.SendMessage(address, message, successMessage)
}

func (h *ConnectionHandler) handleApplicationResponse(am *pb.ApplicationMessage, sender uint64) {
    p := int(am.GetPercent())
    bi := h.conf.Params.Balance
    bj := int(am.GetBalance())

    // Step 5 (swapped places with 6)
    _, node := h.conf.FindById(sender)
    address := node.GetDialAddress()
    message := pbutil.CreateApplicationMessage(h.conf.Self.Id, h.conf.Params.Balance, p)
    successMessage := fmt.Sprintf("Sent application message to node %d: B = %d, p = %d", node.Id, h.conf.Params.Balance, p)
    netutil.SendMessage(address, message, successMessage)

    // Step 7
    if bj >= bi {
        plus := int(math.Round(float64(bj) * (float64(p) / 100)))
        h.conf.Params.Balance = bi + plus
        log.Printf("Increasing balance by %d: Old = %d, New = %d\n", plus, bi, h.conf.Params.Balance)
    } else {
        minus := 1 - (float64(p) / 100)
        h.conf.Params.Balance = int(math.Round(float64(bi) * minus))
        log.Printf("Decreasing balance by %d percent: Old = %d, New = %d\n", p, bi, h.conf.Params.Balance)
    }
}

func (h *ConnectionHandler) handleApplicationAcknowledgment() {
    // TODO unlock
}
