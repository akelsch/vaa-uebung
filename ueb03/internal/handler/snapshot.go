package handler

import (
    "fmt"
    "github.com/akelsch/vaa/ueb03/api/pb"
    "github.com/akelsch/vaa/ueb03/internal/util/netutil"
    "github.com/akelsch/vaa/ueb03/internal/util/pbutil"
    "log"
)

func (h *ConnectionHandler) handleSnapshotMessage(message *pb.Message) {
    sm := message.GetSnapshotMessage()
    identifier := message.GetIdentifier()
    sender := message.GetSender()
    receiver := message.GetReceiver()

    h.dir.Lock()
    defer h.dir.Unlock()

    if sm.Type == pb.SnapshotMessage_MARKER {
        log.Printf("Received snapshot marker from node %d\n", sender)
        h.handleSnapshotMarker(sender)
    } else {
        if !h.dir.Flooding.IsHandled(identifier) {
            if receiver != h.conf.Self.Id {
                h.forwardMessage(message)
            } else {
                switch sm.Type {
                case pb.SnapshotMessage_REQ: // non-coordinator
                    log.Printf("Received snapshot request from node %d\n", sender)
                    h.handleSnapshotRequest(sender)
                case pb.SnapshotMessage_RES: // coordinator
                    log.Printf("Received snapshot response from node %d\n", sender)
                    h.handleSnapshotResponse(message)
                }
            }
        }
    }
}

func (h *ConnectionHandler) handleSnapshotRequest(sender uint64) {
    node := h.conf.FindNodeById(sender)

    metadata := pbutil.CreateMetadata(h.conf.Self.Id, node.Id, h.dir.Flooding.NextSequence())
    message := pbutil.CreateSnapshotResponseMessage(metadata, h.dir.Snapshot.Balance, h.dir.Snapshot.ChangesAsArray())
    h.dir.Flooding.MarkAsHandled(metadata.Identifier)

    successLog := fmt.Sprintf("Sent snapshot response to node %d", node.Id)
    h.unicastMessage(node, message, successLog)

    h.dir.Snapshot.Reset()
}

func (h *ConnectionHandler) handleSnapshotResponse(message *pb.Message) {
    h.dir.Snapshot.StoreResponse(message)

    if len(h.dir.Snapshot.Responses) == h.conf.GetAllNeighborsLength() {
        fmt.Println("done")
        for _, res := range h.dir.Snapshot.Responses {
            fmt.Println(res)
        }
    }
}

func (h *ConnectionHandler) handleSnapshotMarker(sender uint64) {
    if h.dir.Snapshot.IsFirstMarker() {
        h.dir.Snapshot.RecordState(h.conf.Params.Balance)
        h.dir.Snapshot.MarkSenderAsEmpty(sender)
        h.sentSnapshotMarkerToNeighbors()
        h.dir.Snapshot.StartRecording(sender, h.conf.Neighbors)
    } else {
        h.dir.Snapshot.StopRecording(sender)
    }
}

func (h *ConnectionHandler) sentSnapshotMarkerToNeighbors() {
    metadata := pbutil.CreateMetadata(h.conf.Self.Id, 0, 0)
    message := pbutil.CreateSnapshotMarkerMessage(metadata)

    for _, neighbor := range h.conf.Neighbors {
        address := neighbor.GetDialAddress()
        successLog := fmt.Sprintf("Sent snapshot marker to node %d", neighbor.Id)
        netutil.SendMessage(address, message, successLog)
    }
}
