package main

import (
    "flag"
    "fmt"
    "github.com/akelsch/vaa/ueb02/internal/config"
    "github.com/akelsch/vaa/ueb02/internal/errutil"
    "github.com/akelsch/vaa/ueb02/internal/handler"
    "github.com/akelsch/vaa/ueb02/internal/randutil"
    "log"
    "net"
)

func init() {
    log.SetFlags(log.Ltime | log.Lmicroseconds)
    randutil.Init()
}

func main() {
    file := flag.String("f", "config.csv", "path to the CSV file containing the node configuration")
    gvFile := flag.String("gv", "topology.gv", "path to the Graphviz file containing the network topology")
    id := flag.String("id", "1", "ID of this particular node")
    m := flag.Int("m", 1, "Upper bound for preferred time t")
    flag.Parse()

    log.SetPrefix(fmt.Sprintf("[p-%02s] ", *id))

    // Setup configuration
    conf := config.NewConfig(*file, *id)
    conf.Params.T = randutil.RandomInt(1, *m)

    // Listen on own port from configuration
    addr := conf.Self.GetListenAddress()
    ln, err := net.Listen("tcp", addr)
    errutil.HandleError(err)
    log.Printf("Listening on port %s\n", addr)

    // Choose neighbors using Graphviz graph
    conf.ChooseNeighborsByGraph(*gvFile)
    conf.PrintNeighbors()

    log.Printf("Preferred t = %d\n", conf.Params.T)

    // Handle connections
    h := handler.NewConnectionHandler(&ln, conf)
    for {
        conn, err := ln.Accept()
        if err != nil {
            select {
            case <-h.Quit:
                // Listener has been closed by a goroutine
                log.Println("Goodbye!")
                return
            default:
                errutil.HandleError(err)
            }
        } else {
            go h.HandleConnection(conn)
        }
    }
}
