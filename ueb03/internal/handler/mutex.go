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
    h.dir.Mutex.IncrementTimestampBy(h.conf.GetAllNeighborsLength())

    metadata := pbutil.CreateMetadata(h.conf.Self.Id, 0, h.dir.Flooding.NextSequence())
    message := pbutil.CreateMutexRequestMessage(metadata, node.Id, h.dir.Mutex.GetTimestamp())
    h.dir.Flooding.MarkAsHandled(metadata.Identifier)

    log.Printf("Broadcasting mutex request '%s' with resource = %d, timestamp = %d\n",
        message.GetIdentifier(), node.Id, h.dir.Mutex.GetTimestamp())
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

    if !h.dir.Flooding.IsHandled(identifier) {
        h.dir.Flooding.MarkAsHandled(identifier)

        // Step 3b - Request Handling
        switch mm.Type {
        case pb.MutexMessage_REQ:
            h.forwardMessage(message)

            // Ricart-Agrawala
            if !h.dir.Mutex.IsUsingResource(resource) || h.dir.Mutex.GetTimestamp() >= timestamp {
                // send ok
                h.sendMutexResponse(sender, resource)
                h.dir.Mutex.UpdateTimestamp(timestamp)
            } else {
                // queue
                h.dir.Mutex.PushLockRequest(sender, resource, timestamp)
            }
        case pb.MutexMessage_RES:
            if receiver != h.conf.Self.Id {
                h.forwardMessage(message)
            } else {
                log.Printf("Received mutex response from node %d\n", sender)
                h.dir.Mutex.RegisterOk(sender)
                if h.dir.Mutex.CheckIfAllOk(h.conf.GetAllNeighborsLength()) {
                    log.Println("---- LOCK START ---")
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
