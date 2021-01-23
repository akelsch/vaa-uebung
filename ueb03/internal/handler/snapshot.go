package handler

import (
    "fmt"
    "github.com/akelsch/vaa/ueb03/api/pb"
    "github.com/akelsch/vaa/ueb03/internal/util/pbutil"
    "log"
)

func (h *ConnectionHandler) handleSnapshotMessage(message *pb.Message) {
    sm := message.GetSnapshotMessage()
    identifier := message.GetIdentifier()
    receiver := message.GetReceiver()

    h.dir.Lock()
    defer h.dir.Unlock()
    if !h.dir.Flooding.IsHandled(identifier) {
        if receiver != h.conf.Self.Id {
            h.forwardMessage(message)
        } else {
            switch sm.Type {
            case pb.SnapshotMessage_REQ:
                log.Printf("Received snapshot request from node %d\n", message.GetSender())
                h.handleSnapshotRequest(message.GetSender())
            case pb.SnapshotMessage_RES:
                log.Printf("Received snapshot response from node %d\n", message.GetSender())
                fmt.Println("Node", message.GetSender(), "Balance", sm.GetBalance(), "Changes", sm.GetChanges())
                fmt.Println("Self", "Balance", h.dir.Snapshot.State.Balance, "Changes", h.dir.Snapshot.State.Changes)
            }
        }
    }
}

func (h *ConnectionHandler) handleSnapshotRequest(sender uint64) {
    node := h.conf.FindNodeById(sender)

    metadata := pbutil.CreateMetadata(h.conf.Self.Id, node.Id, h.dir.Flooding.NextSequence())
    message := pbutil.CreateSnapshotResponseMessage(metadata, h.dir.Snapshot.State)
    h.dir.Flooding.MarkAsHandled(metadata.Identifier)

    successLog := fmt.Sprintf("Sent snapshot response to node %d", node.Id)
    h.unicastMessage(node, message, successLog)
}
