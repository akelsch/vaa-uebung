package handler

import (
    "fmt"
    "github.com/akelsch/vaa/ueb03/api/pb"
    "github.com/akelsch/vaa/ueb03/internal/directory"
    "github.com/akelsch/vaa/ueb03/internal/util/netutil"
    "github.com/akelsch/vaa/ueb03/internal/util/pbutil"
    "log"
    "strconv"
    "time"
)

func (h *ConnectionHandler) handleStartElection() {
    h.dir.Lock()
    defer h.dir.Unlock()

    electionDir := h.dir.Election

    selfId, _ := strconv.Atoi(h.conf.Self.Id)
    currentId, _ := strconv.Atoi(electionDir.Initiator)

    if selfId > currentId {
        electionDir.Count = 0
        electionDir.Color = directory.RED
        electionDir.Initiator = h.conf.Self.Id
        // do not reset Predecessor as the node could potentially lose the election
        h.propagateExplorerToNeighbors(h.conf.Self.Id, "")
    } else {
        log.Printf("Skipping own election in favor of %d\n", currentId)
    }
}

func (h *ConnectionHandler) handleElectionMessage(message *pb.Message) {
    election := message.GetElection()

    switch election.Type {
    case pb.Election_EXPLORER:
        log.Printf("Received explorer message with ID %s from %s\n", election.GetInitiator(), message.GetSender())
    case pb.Election_ECHO:
        log.Printf("Received echo message with ID %s from %s\n", election.GetInitiator(), message.GetSender())
    }

    h.dir.Lock()
    defer h.dir.Unlock()

    electionDir := h.dir.Election
    debounceElectionVictoryCheck(electionDir.VictoryTimer)

    h.resetElectionIfNecessary(election)
    if electionDir.Color == directory.WHITE {
        electionDir.Color = directory.RED
        electionDir.Initiator = election.GetInitiator()
        electionDir.Predecessor = message.GetSender()
        h.propagateExplorerToNeighbors(electionDir.Initiator, electionDir.Predecessor)
    }

    electionDir.Count++
    if electionDir.Count == len(h.conf.Neighbors) {
        electionDir.Color = directory.GREEN
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

func (h *ConnectionHandler) resetElectionIfNecessary(election *pb.Election) {
    newInitiator, _ := strconv.Atoi(election.GetInitiator())
    oldInitiator, _ := strconv.Atoi(h.dir.Election.Initiator)
    isNotResetAlready := oldInitiator != 0
    if isNotResetAlready && newInitiator > oldInitiator {
        log.Printf("Discarding election of %d in favor of %d\n", oldInitiator, newInitiator)
        h.dir.Election.Reset()
    }
}

func (h *ConnectionHandler) propagateExplorerToNeighbors(initiator string, predecessor string) {
    for _, neighbor := range h.conf.Neighbors {
        if neighbor.Id != predecessor {
            address := neighbor.GetDialAddress()
            message := pbutil.CreateExplorerMessage(h.conf.Self.Id, initiator)
            successMessage := fmt.Sprintf("Sent explorer to node %s", neighbor.Id)
            netutil.SendMessage(address, message, successMessage)
        }
    }
}

func (h *ConnectionHandler) propagateEchoToNeighbors(initiator string, predecessor string) {
    if _, neighbor := h.conf.FindNeighborById(predecessor); neighbor != nil {
        address := neighbor.GetDialAddress()
        message := pbutil.CreateEchoMessage(h.conf.Self.Id, initiator)
        successMessage := fmt.Sprintf("Sent echo to node %s", neighbor.Id)
        netutil.SendMessage(address, message, successMessage)
    }
}

func (h *ConnectionHandler) checkElectionVictory() {
    h.dir.Lock()
    defer h.dir.Unlock()

    if h.dir.Election.IsCoordinator(h.conf.Self.Id) {
        log.Println("------- ELECTION VICTORY -------")
        h.conf.RegisterAllAsNeighbors()

        // propagate START command to random neighbors
        startingNodes := h.conf.GetRandomNeighbors(h.conf.Params.S)
        for _, neighbor := range startingNodes {
            address := neighbor.GetDialAddress()
            message := pbutil.CreateControlMessage(h.conf.Self.Id, pb.ControlMessage_START)
            successMessage := fmt.Sprintf("Sent START command to node %s", neighbor.Id)
            netutil.SendMessage(address, message, successMessage)
        }

        // start double count method to gather results
        h.dir.Status.Ticker = time.NewTicker(1000 * time.Millisecond)
        go func() {
            for {
                select {
                case <-h.dir.Status.Ticker.C:
                    log.Println("------- COUNTING RESULTS -------")
                    for _, neighbor := range h.conf.Neighbors {
                        address := neighbor.GetDialAddress()
                        message := pbutil.CreateControlMessage(h.conf.Self.Id, pb.ControlMessage_GET_STATUS)
                        successMessage := fmt.Sprintf("Sent GET_STATUS command to node %s", neighbor.Id)
                        netutil.SendMessage(address, message, successMessage)
                    }
                }
            }
        }()
    }
}
