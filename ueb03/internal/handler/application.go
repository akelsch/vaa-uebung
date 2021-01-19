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

func (h *ConnectionHandler) handleStart() {
    _, node := h.conf.FindRandomNode()
    percent := randutil.RoundedRandomInt(0, 100, 10)

    // TODO use flooding
    address := node.GetDialAddress()
    message := pbutil.CreateApplicationMessage(h.conf.Self.Id, h.conf.Params.Balance, percent)
    successMessage := fmt.Sprintf("Sent application message to node %s: B = %d, p = %d", node.Id, h.conf.Params.Balance, percent)
    netutil.SendMessage(address, message, successMessage)

    message = pbutil.CreateApplicationRequestMessage(h.conf.Self.Id, percent)
    successMessage = fmt.Sprintf("Sent balance request to node %s", node.Id)
    netutil.SendMessage(address, message, successMessage)
}

func (h *ConnectionHandler) handleApplicationMessage(message *pb.Message) {
    am := message.GetApplicationMessage()
    sender := message.GetSender()

    switch am.Type {
    case pb.ApplicationMessage_NUL:
        log.Printf("Received application message from %s\n", sender)
        h.handleApplicationDefault(am)
    case pb.ApplicationMessage_REQ:
        log.Printf("Received balance request from %s\n", sender)
        h.handleApplicationRequest(am, sender)
    case pb.ApplicationMessage_RES:
        log.Printf("Received balance response from %s\n", sender)
        h.handleApplicationResponse(am)
    case pb.ApplicationMessage_ACK:
        log.Printf("Received acknowledgment from %s\n", sender)
        h.handleApplicationAcknowledgment()
    }
}
func (h *ConnectionHandler) handleApplicationDefault(am *pb.ApplicationMessage) {
    p := int(am.GetPercent())
    bi := int(am.GetBalance())
    bj := h.conf.Params.Balance

    if bi >= bj {
        plus := 1 + (float64(p) / 100)
        h.conf.Params.Balance = int(math.Round(float64(bj) * plus))
        log.Printf("Increasing balance by %d percent: Old = %d, New = %d\n", p, bj, h.conf.Params.Balance)
    } else {
        minus := 1 - (float64(p) / 100)
        h.conf.Params.Balance = int(math.Round(float64(bj) * minus))
        log.Printf("Decreasing balance by %d percent: Old = %d, New = %d\n", p, bj, h.conf.Params.Balance)
    }
}

func (h *ConnectionHandler) handleApplicationRequest(am *pb.ApplicationMessage, sender string) {
    _, node := h.conf.FindById(sender)
    log.Println(sender)
    percent := int(am.GetPercent())

    // TODO use flooding
    address := node.GetDialAddress()
    message := pbutil.CreateApplicationResponseMessage(h.conf.Self.Id, h.conf.Params.Balance, percent)
    successMessage := fmt.Sprintf("Sent balance response to node %s: B = %d, p = %d", node.Id, h.conf.Params.Balance, percent)
    netutil.SendMessage(address, message, successMessage)
}

func (h *ConnectionHandler) handleApplicationResponse(am *pb.ApplicationMessage) {
    p := int(am.GetPercent())
    bi := h.conf.Params.Balance
    bj := int(am.GetBalance())

    if bj >= bi {
        plus := 1 + (float64(p) / 100)
        h.conf.Params.Balance = int(math.Round(float64(bi) * plus))
        log.Printf("Increasing balance by %d percent: Old = %d, New = %d\n", p, bi, h.conf.Params.Balance)
    } else {
        minus := 1 - (float64(p) / 100)
        h.conf.Params.Balance = int(math.Round(float64(bi) * minus))
        log.Printf("Decreasing balance by %d percent: Old = %d, New = %d\n", p, bi, h.conf.Params.Balance)
    }

    // TODO send ACK
}

func (h *ConnectionHandler) handleApplicationAcknowledgment() {
    // TODO unlock
}
