package main

import (
    "flag"
    "fmt"
    "github.com/akelsch/vaa/ueb01/internal/config"
    "github.com/akelsch/vaa/ueb01/internal/errutil"
    "github.com/akelsch/vaa/ueb01/internal/handler"
    "github.com/akelsch/vaa/ueb01/internal/randutil"
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
    n := flag.Int("n", 0, "number of random neighbors if not using Graphviz")
    flag.Parse()

    log.SetPrefix(fmt.Sprintf("[node-%03s] ", *id))

    // 1-2
    conf := config.NewConfig(*file, *id)

    // 3
    addr := conf.Self.GetListenAddress()
    ln, err := net.Listen("tcp", addr)
    errutil.HandleError(err)
    log.Printf("Listening on port %s\n", addr)

    // 4
    if *n != 0 {
        conf.ChooseNeighborsRandomly(*n)
    } else {
        conf.ChooseNeighborsByGraph(*gvFile)
    }
    conf.PrintNeighbors()

    // 5-9
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
