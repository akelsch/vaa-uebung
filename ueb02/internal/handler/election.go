package handler

import (
    "github.com/akelsch/vaa/ueb02/api/pb"
    "github.com/akelsch/vaa/ueb02/internal/directory"
    "github.com/akelsch/vaa/ueb02/internal/errutil"
    "github.com/akelsch/vaa/ueb02/internal/pbutil"
    "google.golang.org/protobuf/proto"
    "log"
    "net"
    "strconv"
    "time"
)

const victoryTimeout = 1 * time.Second

func (h *ConnectionHandler) handleStartElection() {
    electionDir := h.dir.Election

    electionDir.Count = 0
    electionDir.Color = directory.RED
    electionDir.Initiator = h.conf.Self.Id
    // do not reset Predecessor as the node could potentially lose the election
    h.propagateExplorerToNeighbors(h.conf.Self.Id, "")
}

func (h *ConnectionHandler) handleElectionMessage(message *pb.Message) {
    election := message.GetElection()

    switch election.Type {
    case pb.Election_EXPLORER:
        log.Printf("Received explorer message with ID %s from %s\n", election.GetInitiator(), message.GetSender())
    case pb.Election_ECHO:
        log.Printf("Received echo message with ID %s from %s\n", election.GetInitiator(), message.GetSender())
    }

    electionDir := h.dir.Election
    debounceElectionVictory(electionDir.VictoryTimer)

    h.dir.Lock()
    h.resetForHigherInitiator(election)
    electionDir.Count++

    if electionDir.Color == directory.WHITE {
        electionDir.Color = directory.RED
        electionDir.Initiator = election.GetInitiator()
        electionDir.Predecessor = message.GetSender()
        h.propagateExplorerToNeighbors(electionDir.Initiator, message.GetSender())
    }

    if electionDir.Count == len(h.conf.Neighbors) {
        electionDir.Color = directory.GREEN
        if electionDir.IsNotInitiator(h.conf.Self.Id) {
            h.propagateEchoToNeighbors(electionDir.Initiator, electionDir.Predecessor)
        } else {
            log.Println("INITIATOR IS GREEN")
            electionDir.VictoryTimer = time.AfterFunc(victoryTimeout, h.checkElectionVictory)
        }
    }
    h.dir.Unlock()
}

func debounceElectionVictory(timer *time.Timer) {
    // await "last" election message before declaring election victory
    if timer != nil {
        timer.Reset(victoryTimeout)
    }
}

func (h *ConnectionHandler) resetForHigherInitiator(election *pb.Election) {
    newInitiator, _ := strconv.Atoi(election.GetInitiator())
    oldInitiator, _ := strconv.Atoi(h.dir.Election.Initiator)
    if oldInitiator != 0 && newInitiator > oldInitiator {
        log.Printf("Discarding election of %d in favor for %d\n", oldInitiator, newInitiator)
        h.dir.Election.Reset()
    }
}

func (h *ConnectionHandler) propagateExplorerToNeighbors(initiator string, sender string) {
    for _, neighbor := range h.conf.Neighbors {
        if neighbor.Id != sender {
            conn, err := net.Dial("tcp", neighbor.GetDialAddress())
            if err != nil {
                log.Printf("Could not connect to node %s\n", neighbor.Id)
            } else {
                bytes, err := proto.Marshal(pbutil.CreateExplorerMessage(h.conf.Self.Id, initiator))
                errutil.HandleError(err)
                _, err = conn.Write(bytes)
                errutil.HandleError(err)
                conn.Close()
                log.Printf("Sent explorer to node %s\n", neighbor.Id)
            }
        }
    }
}

func (h *ConnectionHandler) propagateEchoToNeighbors(initiator string, predecessor string) {
    for _, neighbor := range h.conf.Neighbors {
        if neighbor.Id == predecessor {
            conn, err := net.Dial("tcp", neighbor.GetDialAddress())
            if err != nil {
                log.Printf("Could not connect to node %s\n", neighbor.Id)
            } else {
                bytes, err := proto.Marshal(pbutil.CreateEchoMessage(h.conf.Self.Id, initiator))
                errutil.HandleError(err)
                _, err = conn.Write(bytes)
                errutil.HandleError(err)
                conn.Close()
                log.Printf("Sent echo to node %s\n", neighbor.Id)
            }
        }
    }
}

func (h *ConnectionHandler) checkElectionVictory() {
    h.dir.Lock()
    // check if current node is still the initiator of the last election message
    if !h.dir.Election.IsNotInitiator(h.conf.Self.Id) {
        log.Println("ELECTION VICTORY")
        h.conf.RegisterAllAsNeighbors()
        // TODO choose "s" random neighbors and send them a START
    }
    h.dir.Unlock()
}
