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
    sender := message.GetSender()
    receiver := message.GetReceiver()

    h.dir.Lock()
    defer h.dir.Unlock()
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
            case pb.SnapshotMessage_MARKER: // both
                log.Printf("Received snapshot marker from node %d\n", sender)
                h.handleSnapshotMarker(sender)
            }
        }
    }
}

func (h *ConnectionHandler) handleSnapshotRequest(sender uint64) {
    node := h.conf.FindNodeById(sender)

    balance := h.dir.Snapshot.Balance
    changes := h.dir.Snapshot.ChangesAsArray()
    finished := h.dir.Snapshot.AreAllRecordingsStopped()

    metadata := pbutil.CreateMetadata(h.conf.Self.Id, node.Id, h.dir.Flooding.NextSequence())
    message := pbutil.CreateSnapshotResponseMessage(metadata, balance, changes, finished)
    h.dir.Flooding.MarkAsHandled(metadata.Identifier)

    successLog := fmt.Sprintf("Sent snapshot response to node %d", node.Id)
    h.unicastMessage(node, message, successLog)

    h.dir.Snapshot.Reset()
}

func (h *ConnectionHandler) handleSnapshotResponse(message *pb.Message) {
    h.dir.Snapshot.StoreResponse(message)

    if len(h.dir.Snapshot.Responses) == h.conf.GetAllNeighborsLength() {
        previousBalance := h.dir.Snapshot.PreviousBalance
        currentBalance := h.dir.Snapshot.Balance

        fmt.Println(currentBalance, h.dir.Snapshot.ChangesAsArray(), h.dir.Snapshot.AreAllRecordingsStopped())
        for _, res := range h.dir.Snapshot.Responses {
            sm := res.GetSnapshotMessage()
            fmt.Println(sm.GetBalance(), sm.GetChanges(), sm.GetFinished())
            currentBalance += sm.GetBalance()
        }

        logMessage := fmt.Sprintf("Previous system balance = %d\n", previousBalance)
        logMessage += fmt.Sprintf("Current system balance = %d\n", currentBalance)
        if previousBalance != currentBalance && previousBalance != 0 {
            logMessage += fmt.Sprintf("BALANCES DO NOT MATCH! DIFFERENCE = %d\n", currentBalance-previousBalance)
        }

        h.dir.Snapshot.PreviousBalance = currentBalance
        fmt.Print(logMessage)
    }
}

func (h *ConnectionHandler) handleSnapshotMarker(sender uint64) {
    if h.dir.Snapshot.IsFirstMarker() {
        h.dir.Snapshot.RecordState(h.conf.Params.Balance)
        h.dir.Snapshot.MarkSenderAsEmpty(sender)
        h.sentSnapshotMarkerToNeighbors()
        h.dir.Snapshot.StartRecording(sender, h.conf.FindAllNeighbors())
    } else {
        h.dir.Snapshot.StopRecording(sender)
    }
}

func (h *ConnectionHandler) sentSnapshotMarkerToNeighbors() {
    for _, neighbor := range h.conf.FindAllNeighbors() {
        metadata := pbutil.CreateMetadata(h.conf.Self.Id, neighbor.Id, h.dir.Flooding.NextSequence())
        message := pbutil.CreateSnapshotMarkerMessage(metadata)
        h.dir.Flooding.MarkAsHandled(metadata.Identifier)
        successLog := fmt.Sprintf("Sent snapshot marker to node %d", neighbor.Id)
        h.unicastMessage(neighbor, message, successLog)
    }
}
