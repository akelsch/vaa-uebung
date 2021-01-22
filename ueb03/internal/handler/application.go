package handler

import (
    "fmt"
    "github.com/akelsch/vaa/ueb03/api/pb"
    "github.com/akelsch/vaa/ueb03/internal/util/pbutil"
    "github.com/akelsch/vaa/ueb03/internal/util/randutil"
    "log"
    "math"
)

func (h *ConnectionHandler) startApplicationSteps(resource uint64) {
    // Step 4
    percent := randutil.RoundedRandomUint(0, 100, 10)

    // Step 5
    node := h.conf.FindNodeById(resource)
    metadata := pbutil.CreateMetadata(h.conf.Self.Id, node.Id, h.dir.Flooding.NextSequence())
    message := pbutil.CreateApplicationRequestMessage(metadata, percent)
    successLog := fmt.Sprintf("Sent balance request to node %d", node.Id)
    h.unicastMessage(node, message, successLog)
}

func (h *ConnectionHandler) handleApplicationMessage(message *pb.Message) {
    am := message.GetApplicationMessage()
    identifier := message.GetIdentifier()
    sender := message.GetSender()
    receiver := message.GetReceiver()

    h.dir.Lock()
    defer h.dir.Unlock()
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
                h.handleApplicationAcknowledgment(sender)
            }
        }
    } else {
        //log.Printf("Message '%s' got handled already\n", identifier)
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
        log.Printf("Increasing balance by %d: Old = %d, New = %d (8A)\n", plus, bj, h.conf.Params.Balance)
    } else {
        minus := int64(math.Round(float64(bj) * (float64(p) / 100)))
        h.conf.Params.Balance = bj - minus
        log.Printf("Decreasing balance by %d: Old = %d, New = %d (8B)\n", minus, bj, h.conf.Params.Balance)
    }

    // Step 9
    node := h.conf.FindNodeById(sender)
    metadata := pbutil.CreateMetadata(h.conf.Self.Id, node.Id, h.dir.Flooding.NextSequence())
    message := pbutil.CreateApplicationAcknowledgmentMessage(metadata)
    successLog := fmt.Sprintf("Sent acknowledgment to node %d", node.Id)
    h.unicastMessage(node, message, successLog)
}

func (h *ConnectionHandler) handleApplicationRequest(am *pb.ApplicationMessage, sender uint64) {
    node := h.conf.FindNodeById(sender)
    percent := am.GetPercent()

    metadata := pbutil.CreateMetadata(h.conf.Self.Id, node.Id, h.dir.Flooding.NextSequence())
    message := pbutil.CreateApplicationResponseMessage(metadata, h.conf.Params.Balance, percent)
    successLog := fmt.Sprintf("Sent balance response to node %d: B = %d", node.Id, h.conf.Params.Balance)
    h.unicastMessage(node, message, successLog)
}

func (h *ConnectionHandler) handleApplicationResponse(am *pb.ApplicationMessage, sender uint64) {
    p := am.GetPercent()
    bi := h.conf.Params.Balance
    bj := am.GetBalance()

    // Step 6
    node := h.conf.FindNodeById(sender)
    metadata := pbutil.CreateMetadata(h.conf.Self.Id, node.Id, h.dir.Flooding.NextSequence())
    message := pbutil.CreateApplicationMessage(metadata, h.conf.Params.Balance, p)
    successLog := fmt.Sprintf("Sent application message to node %d: B = %d, p = %d", node.Id, h.conf.Params.Balance, p)
    h.unicastMessage(node, message, successLog)

    // Step 7
    if bj >= bi {
        plus := int64(math.Round(float64(bj) * (float64(p) / 100)))
        h.conf.Params.Balance = bi + plus
        log.Printf("Increasing balance by %d: Old = %d, New = %d (7A)\n", plus, bi, h.conf.Params.Balance)
    } else {
        minus := int64(math.Round(float64(bi) * (float64(p) / 100)))
        h.conf.Params.Balance = bi - minus
        log.Printf("Decreasing balance by %d: Old = %d, New = %d (7B)\n", minus, bi, h.conf.Params.Balance)
    }
}

func (h *ConnectionHandler) handleApplicationAcknowledgment(sender uint64) {
    log.Printf("--- UNLOCKING RESOURCE %d ---\n", sender)
    h.dir.Mutex.ResetInterestInResource()
    h.dir.Mutex.ResetOk()

    // Step 10
    item := h.dir.Mutex.PopLockRequest()
    for item != nil {
        h.sendMutexResponse(item.Sender, item.Resource)
        item = h.dir.Mutex.PopLockRequest()
    }

    // Step 11
    h.StartFirstStep()
}
