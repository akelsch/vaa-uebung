package main

import (
    "flag"
    "fmt"
    "github.com/akelsch/vaa/ueb03/internal/config"
    "github.com/akelsch/vaa/ueb03/internal/handler"
    "github.com/akelsch/vaa/ueb03/internal/util/errutil"
    "github.com/akelsch/vaa/ueb03/internal/util/randutil"
    "log"
    "net"
)

func main() {
    file := flag.String("f", "config.csv", "path to the CSV file containing the node configuration")
    gvFile := flag.String("gv", "topology.gv", "path to the Graphviz file containing the network topology")
    id := flag.String("id", "1", "ID of this particular node")
    flag.Parse()

    // Init rand & log
    randutil.Init(*id)
    log.SetFlags(log.Ltime | log.Lmicroseconds)
    log.SetPrefix(fmt.Sprintf("[account-%03s] ", *id))

    // Setup configuration
    conf := config.NewConfig(*file, *id)
    conf.Params.Balance = randutil.RandomInt(0, 100_000)
    conf.Params.Balance = conf.Params.Balance - (conf.Params.Balance % 1_000)

    // Listen on own port from configuration
    ln, err := net.Listen("tcp", conf.Self.GetListenAddress())
    errutil.HandleError(err)
    defer ln.Close()

    // Choose neighbors using Graphviz graph
    conf.ChooseNeighborsByGraph(*gvFile)

    // Print node details
    log.Printf("Listening on port %s\n", conf.Self.GetListenAddress())
    log.Println(conf.NeighborsToString())
    log.Printf("Balance: %d\n", conf.Params.Balance)

    // Handle connections
    h := handler.NewConnectionHandler(&ln, conf)
    for {
        conn, err := ln.Accept()
        if err != nil {
            h.HandleError(err)
            return
        }

        go h.HandleConnection(conn)
    }
}
