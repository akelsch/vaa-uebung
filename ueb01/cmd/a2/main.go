package main

import (
    "flag"
    "fmt"
    "github.com/akelsch/vaa/ueb01/internal/conf"
    "github.com/akelsch/vaa/ueb01/internal/errutil"
    "github.com/akelsch/vaa/ueb01/internal/randutil"
    "log"
    "net"
    "os"
    "strings"
)

const (
    protocol      = "tcp"
    controlPrefix = "c:"
)

var config = conf.NewConfig()
var directory = conf.NewNeighborDirectory()

func init() {
    randutil.Init()
}

func main() {
    file := flag.String("f", "config.csv", "path to the CSV file containing the node configuration")
    gvFile := flag.String("gv", "topology.gv", "path to the Graphviz file containing the network topology")
    id := flag.String("id", "1", "ID of the node")
    flag.Parse()

    // 1-2
    config.Init(*file, *id)

    // 3
    listener, err := net.Listen(protocol, config.Self.GetListenAddress())
    errutil.HandleError(err)
    fmt.Printf("Node %s is listening on port %s\n", config.Self.Id, config.Self.Port)

    // 4
    config.ChooseNeighborsByGraph(*gvFile)
    config.PrintNeighbors()

    // 5-9
    defer listener.Close()
    for {
        conn, err := listener.Accept()
        errutil.HandleError(err)
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    buf := make([]byte, 1024)
    n, err := conn.Read(buf)
    errutil.HandleError(err)

    message := strings.TrimSpace(string(buf[:n]))
    log.Printf("Received: %s\n", message)

    if strings.HasPrefix(message, controlPrefix) {
        handleControlMessage(message)
    } else {
        handleApplicationMessage(message)
    }
}

func handleControlMessage(message string) {
    command := strings.TrimPrefix(message, controlPrefix)
    if command == "start" {
        directory.Lock()
        directory.Reset()
        sendToAllNeighbors()
        directory.Unlock()
    } else if command == "exit" {
        // TODO exit all
        os.Exit(0)
    }
}

func handleApplicationMessage(message string) {
    neighbors := config.Neighbors
    for i := range neighbors {
        if message == neighbors[i].Id {
            directory.Lock()
            if directory.HasAlreadyReceivedFrom(i) {
                directory.ResetIfNecessary(len(neighbors))
            }
            directory.SetReceived(i)
            sendToAllNeighbors()
            directory.Unlock()
        }
    }
}

func sendToAllNeighbors() {
    neighbors := config.Neighbors
    for i := range neighbors {
        if directory.HasNotSentTo(i) {
            neighbor := neighbors[i]
            conn, err := net.Dial(protocol, neighbor.GetDialAddress())
            if err != nil {
                log.Printf("Could not connect to node %s", neighbor.Id)
            } else {
                payload := []byte(config.Self.Id)
                _, err := conn.Write(payload)
                errutil.HandleError(err)
                directory.SetSent(i)
                log.Printf("Sent %s to node %s\n", payload, neighbor.Id)
            }
        }
    }
}
