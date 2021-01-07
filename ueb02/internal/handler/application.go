package handler

import (
    "github.com/akelsch/vaa/ueb02/api/pb"
    "github.com/akelsch/vaa/ueb02/internal/errutil"
    "github.com/akelsch/vaa/ueb02/internal/pbutil"
    "google.golang.org/protobuf/proto"
    "log"
    "net"
)

func (h *ConnectionHandler) handleApplicationMessage(message *pb.Message) {
    am := message.GetApplicationMessage()
    t2 := int(am.GetBody())
    log.Printf("Received application message: %d\n", t2)

    if h.conf.Params.AMax > 0 {
        h.conf.Params.AMax--
        oldT := h.conf.Params.T
        h.conf.Params.T = (h.conf.Params.T + t2) / 2
        log.Printf("Old t = %d, new t = %d\n", oldT, h.conf.Params.T)
    }
}

func (h *ConnectionHandler) exchangeTimeWithNeighbors() {
    randomNeighbors := h.conf.GetRandomNeighbors(h.conf.Params.P)
    for _, neighbor := range randomNeighbors {
        conn, err := net.Dial("tcp", neighbor.GetDialAddress())
        if err != nil {
            log.Printf("Could not connect to node %s\n", neighbor.Id)
        } else {
            bytes, err := proto.Marshal(pbutil.CreateApplicationMessage(h.conf.Params.T))
            errutil.HandleError(err)
            _, err = conn.Write(bytes)
            errutil.HandleError(err)
            conn.Close()
            log.Printf("Sent application message to node %s: %d\n", neighbor.Id, h.conf.Params.T)
        }
    }
}
