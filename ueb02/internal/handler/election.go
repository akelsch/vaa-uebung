package handler

import (
    "github.com/akelsch/vaa/ueb02/api/pb"
    "github.com/akelsch/vaa/ueb02/internal/directory"
    "github.com/akelsch/vaa/ueb02/internal/errutil"
    "github.com/akelsch/vaa/ueb02/internal/pbutil"
    "google.golang.org/protobuf/proto"
    "log"
    "net"
)

func (h *ConnectionHandler) handleStartElection() {
    electionDir := h.dir.Election

    electionDir.Color = directory.RED
    electionDir.Initiator = h.conf.Self.Id
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

    h.dir.Lock()
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
        }
    }
    h.dir.Unlock()
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
