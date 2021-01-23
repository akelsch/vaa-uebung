package handler

import (
    "fmt"
    "github.com/akelsch/vaa/ueb03/api/pb"
    "github.com/akelsch/vaa/ueb03/internal/directory/election/color"
    "github.com/akelsch/vaa/ueb03/internal/util/netutil"
    "github.com/akelsch/vaa/ueb03/internal/util/pbutil"
    "log"
    "time"
)

func (h *ConnectionHandler) handleStartElection() {
    h.dir.Lock()
    defer h.dir.Unlock()

    electionDir := h.dir.Election

    selfId := h.conf.Self.Id
    currentId := electionDir.Initiator

    if selfId > currentId {
        electionDir.Count = 0
        electionDir.Color = color.RED
        electionDir.Initiator = h.conf.Self.Id
        // do not reset Predecessor as the node could potentially lose the election
        h.propagateExplorerToNeighbors(selfId, 0)
    } else {
        log.Printf("Skipping own election in favor of %d\n", currentId)
    }
}

func (h *ConnectionHandler) handleElectionMessage(message *pb.Message) {
    em := message.GetElectionMessage()

    switch em.Type {
    case pb.ElectionMessage_EXPLORER:
        log.Printf("Received explorer message with ID %d from %d\n", em.GetInitiator(), message.GetSender())
    case pb.ElectionMessage_ECHO:
        log.Printf("Received echo message with ID %d from %d\n", em.GetInitiator(), message.GetSender())
    }

    h.dir.Lock()
    defer h.dir.Unlock()

    electionDir := h.dir.Election
    debounceElectionVictoryCheck(electionDir.VictoryTimer)

    h.resetElectionIfNecessary(em)
    if electionDir.Color == color.WHITE {
        electionDir.Color = color.RED
        electionDir.Initiator = em.GetInitiator()
        electionDir.Predecessor = message.GetSender()
        h.propagateExplorerToNeighbors(electionDir.Initiator, electionDir.Predecessor)
    }

    electionDir.Count++
    if electionDir.Count == len(h.conf.Neighbors) {
        electionDir.Color = color.GREEN
        if electionDir.IsInitiator(h.conf.Self.Id) {
            log.Println("INITIATOR IS GREEN")
            electionDir.VictoryTimer = time.AfterFunc(1000*time.Millisecond, h.checkElectionVictory)
        } else {
            h.propagateEchoToNeighbors(electionDir.Initiator, electionDir.Predecessor)
        }
    }
}

func debounceElectionVictoryCheck(timer *time.Timer) {
    if timer != nil {
        timer.Reset(1000 * time.Millisecond)
    }
}

func (h *ConnectionHandler) resetElectionIfNecessary(em *pb.ElectionMessage) {
    newInitiator := em.GetInitiator()
    oldInitiator := h.dir.Election.Initiator
    isNotResetAlready := oldInitiator != 0
    if isNotResetAlready && newInitiator > oldInitiator {
        log.Printf("Discarding election of %d in favor of %d\n", oldInitiator, newInitiator)
        h.dir.Election.Reset()
    }
}

func (h *ConnectionHandler) propagateExplorerToNeighbors(initiator uint64, predecessor uint64) {
    for _, neighbor := range h.conf.Neighbors {
        if neighbor.Id != predecessor {
            address := neighbor.GetDialAddress()
            message := pbutil.CreateExplorerMessage(h.conf.Self.Id, initiator)
            successMessage := fmt.Sprintf("Sent explorer to node %d", neighbor.Id)
            netutil.SendMessage(address, message, successMessage)
        }
    }
}

func (h *ConnectionHandler) propagateEchoToNeighbors(initiator uint64, predecessor uint64) {
    if neighbor := h.conf.FindNeighborById(predecessor); neighbor != nil {
        address := neighbor.GetDialAddress()
        message := pbutil.CreateEchoMessage(h.conf.Self.Id, initiator)
        successMessage := fmt.Sprintf("Sent echo to node %d", neighbor.Id)
        netutil.SendMessage(address, message, successMessage)
    }
}

func (h *ConnectionHandler) checkElectionVictory() {
    h.dir.Lock()
    defer h.dir.Unlock()

    if h.dir.Election.IsCoordinator(h.conf.Self.Id) {
        log.Println("------- ELECTION VICTORY -------")

        // Flood START command
        metadata := pbutil.CreateMetadata(h.conf.Self.Id, 0, h.dir.Flooding.NextSequence())
        message := pbutil.CreateControlMessage(metadata, pb.ControlMessage_START)
        for _, neighbor := range h.conf.Neighbors {
            address := neighbor.GetDialAddress()
            netutil.SendMessageSilently(address, message)
        }

        h.startTakingSnapshots()
    }
}

func (h *ConnectionHandler) startTakingSnapshots() {
    time.AfterFunc(4000*time.Millisecond, func() {
        log.Println("------- TAKING SNAPSHOT -------")
        h.dir.Snapshot.Reset()
        h.dir.Snapshot.IsFirstMarker()
        h.sentSnapshotMarkerToNeighbors()

        time.AfterFunc(1000*time.Millisecond, func() {
            for _, neighbor := range h.conf.FindAllNeighbors() {
                metadata := pbutil.CreateMetadata(h.conf.Self.Id, neighbor.Id, h.dir.Flooding.NextSequence())
                message := pbutil.CreateSnapshotRequestMessage(metadata)
                h.dir.Flooding.MarkAsHandled(metadata.Identifier)
                successLog := fmt.Sprintf("Sent snapshot request to node %d", neighbor.Id)
                h.unicastMessage(neighbor, message, successLog)
            }
            h.startTakingSnapshots()
        })
    })
}
