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
    h.dir.Flooding.MarkAsHandled(metadata.GetIdentifier())

    log.Printf("*** Flooding lock request '%s' with resource = %d & timestamp = %d ***\n",
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
            h.forwardMutexMessage(message)

            // Ricart-Agrawala
            if !h.dir.Mutex.IsUsingResource(resource) || h.dir.Mutex.GetTimestamp() >= timestamp {
                // send ok
                metadata := pbutil.CreateMetadata(h.conf.Self.Id, sender, h.dir.Flooding.NextSequence())
                message := pbutil.CreateMutexResponseMessage(metadata)
                h.dir.Flooding.MarkAsHandled(metadata.GetIdentifier())

                log.Printf("Flooding lock response for node %d\n", sender)
                for _, neighbor := range h.conf.Neighbors {
                    address := neighbor.GetDialAddress()
                    netutil.SendMessageSilently(address, message)
                }

                h.dir.Mutex.UpdateTimestamp(timestamp)
            } else {
                // queue
                h.dir.Mutex.QueueLockRequest(resource, timestamp)
            }
        case pb.MutexMessage_RES:
            if receiver != h.conf.Self.Id {
                h.forwardMutexMessage(message)
            } else {
                log.Printf("Received OK from node %d\n", sender)
            }
        }
    } else {
        log.Printf("Mutex message '%s' got handled already\n", identifier)
    }
}

func (h *ConnectionHandler) forwardMutexMessage(message *pb.Message) {
    for _, neighbor := range h.conf.Neighbors {
        address := neighbor.GetDialAddress()
        successLog := fmt.Sprintf("Forwarded mutex message '%s' to node %d", message.GetIdentifier(), neighbor.Id)
        netutil.SendMessage(address, message, successLog)
    }
}
