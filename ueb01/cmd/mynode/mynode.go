package main

import (
    "flag"
    "fmt"
    "github.com/akelsch/vaa/ueb01/internal/conf"
    "github.com/akelsch/vaa/ueb01/internal/connhandler"
    "github.com/akelsch/vaa/ueb01/internal/errutil"
    "github.com/akelsch/vaa/ueb01/internal/randutil"
    "net"
)

func init() {
    randutil.Init()
}

func main() {
    file := flag.String("f", "config.csv", "path to the CSV file containing the node configuration")
    gvFile := flag.String("gv", "topology.gv", "path to the Graphviz file containing the network topology")
    id := flag.String("id", "1", "ID of this particular node")
    n := flag.Int("n", 0, "number of random neighbors if not using Graphviz")
    flag.Parse()

    // 1-2
    config := conf.NewConfig(*file, *id)

    // 3
    listener, err := net.Listen("tcp", config.Self.GetListenAddress())
    errutil.HandleError(err)
    fmt.Printf("Node %s is listening on port %s\n", config.Self.Id, config.Self.Port)

    // 4
    if *n != 0 {
        config.ChooseNeighborsRandomly(*n)
    } else {
        config.ChooseNeighborsByGraph(*gvFile)
    }
    config.PrintNeighbors()

    // 5-9
    handler := connhandler.NewConnectionHandler(config)
    defer listener.Close()
    for {
        conn, err := listener.Accept()
        errutil.HandleError(err)
        go handler.HandleConnection(conn)
    }
}
