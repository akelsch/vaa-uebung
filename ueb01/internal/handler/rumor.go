package handler

import (
    "github.com/akelsch/vaa/ueb01/api/pb"
    "github.com/akelsch/vaa/ueb01/internal/errutil"
    "github.com/akelsch/vaa/ueb01/internal/pbutil"
    "google.golang.org/protobuf/proto"
    "log"
    "net"
)

func (h *ConnectionHandler) handleRumorMessage(message *pb.Message) {
    rumor := message.GetRumor()
    log.Printf("Received rumor: %s\n", rumor.Content)

    h.dir.Lock()
    h.dir.Rumors.RememberRumor(rumor.Content)

    if h.dir.Rumors.IsNewRumor(rumor.Content) {
        h.propagateRumorToNeighbors(rumor, message.Sender)
    } else {
        log.Printf("Heard about rumor '%s' %d times", rumor.Content, h.dir.Rumors.GetRumorCount(rumor.Content))
    }

    h.dir.Unlock()
}

func (h *ConnectionHandler) propagateRumorToNeighbors(rumor *pb.Rumor, sender string) {
    for _, neighbor := range h.conf.Neighbors {
        if neighbor.Id != sender {
            conn, err := net.Dial("tcp", neighbor.GetDialAddress())
            if err != nil {
                log.Printf("Could not connect to node %s", neighbor.Id)
            } else {
                bytes, err := proto.Marshal(pbutil.CreateRumorMessage(h.conf.Self.Id, rumor))
                errutil.HandleError(err)
                _, err = conn.Write(bytes)
                errutil.HandleError(err)
                conn.Close()
                log.Printf("Told node %s about rumor '%s'\n", neighbor.Id, rumor.Content)
            }
        }
    }
}
