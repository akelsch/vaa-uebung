package handler

import (
    "fmt"
    "github.com/akelsch/vaa/ueb03/api/pb"
    "github.com/akelsch/vaa/ueb03/internal/util/netutil"
    "github.com/akelsch/vaa/ueb03/internal/util/pbutil"
    "log"
)

func (h *ConnectionHandler) handleStart() {
    // Step 2
    node := h.conf.FindRandomNode()

    // Step 3a - Request Critical Section
    h.dir.Lock()
    defer h.dir.Unlock()
    h.dir.Mutex.RegisterWant()

    timestamp := h.dir.Mutex.IncrementTimestamp()
    metadata := pbutil.CreateMetadata(h.conf.Self.Id, 0, h.dir.Flooding.NextSequence())
    message := pbutil.CreateMutexRequestMessage(metadata, node.Id, timestamp)
    h.dir.Flooding.MarkAsHandled(metadata.Identifier)

    log.Printf("Broadcasting mutex request '%s' with resource = %d, timestamp = %d\n",
        message.GetIdentifier(), node.Id, timestamp)
    for _, neighbor := range h.conf.Neighbors {
        address := neighbor.GetDialAddress()
        netutil.SendMessageSilently(address, message)
    }
}

func (h *ConnectionHandler) handleMutexMessage(message *pb.Message) {
    mm := message.GetMutexMessage()
    identifier := message.GetIdentifier()
    sender := message.GetSender()
    receiver := message.GetReceiver()
    resource := mm.GetResource()
    timestamp := mm.GetTimestamp()

    h.dir.Lock()
    defer h.dir.Unlock()
    if !h.dir.Flooding.IsHandled(identifier) {
        h.dir.Flooding.MarkAsHandled(identifier)

        // Step 3b - Request Handling
        switch mm.Type {
        case pb.MutexMessage_REQ:
            log.Printf("Received mutex request from node %d\n", sender)
            h.forwardMessage(message)

            if h.dir.Mutex.NeedsToQueue(timestamp, resource, h.conf.Self.Id) {
                log.Printf("*** Queueing s=%d, r=%d, %d<%d\n", sender, resource, h.dir.Mutex.GetTimestamp(), timestamp)
                // queue
                h.dir.Mutex.PushLockRequest(sender, resource)
            } else {
                // send ok
                h.sendMutexResponse(sender, resource)
            }

            h.dir.Mutex.UpdateTimestamp(timestamp)
        case pb.MutexMessage_RES:
            if receiver != h.conf.Self.Id {
                h.forwardMessage(message)
            } else {
                log.Printf("Received mutex response from node %d\n", sender)
                h.dir.Mutex.RegisterResponse(sender)
                if h.dir.Mutex.CheckResponseCount(h.conf.GetAllNeighborsLength()) {
                    log.Printf("--- LOCKING RESOURCE %d ---\n", resource)
                    h.dir.Mutex.RegisterLock()
                    h.startApplicationSteps(resource)
                }
            }
        }
    } else {
        //log.Printf("Mutex message '%s' got handled already\n", identifier)
    }
}

func (h *ConnectionHandler) sendMutexResponse(receiver uint64, resource uint64) {
    metadata := pbutil.CreateMetadata(h.conf.Self.Id, receiver, h.dir.Flooding.NextSequence())
    message := pbutil.CreateMutexResponseMessage(metadata, resource)
    h.dir.Flooding.MarkAsHandled(metadata.Identifier)

    node := h.conf.FindNodeById(receiver)
    successLog := fmt.Sprintf("Sent mutex response to node %d", node.Id)
    h.unicastMessage(node, message, successLog)
}
