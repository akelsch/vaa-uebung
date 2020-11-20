package main

import (
    "fmt"
    "github.com/akelsch/vaa/ueb01/conf"
    "github.com/akelsch/vaa/ueb01/errutil"
    "log"
    "net"
    "os"
    "strings"
    "sync"
)

const (
    configFile    = "mapping.csv"
    protocol      = "tcp"
    neighborCount = 3
)

type NeighborDirectory struct {
    mu sync.Mutex
    v  map[int]bool
}

var directory = NeighborDirectory{v: make(map[int]bool)}

func main() {
    id := getIdFromArgs()

    configuration := conf.Init(configFile, id)
    self := configuration.Find(id)
    if self == nil {
        log.Fatalf("Could not determine configuration for entry with ID %s", id)
    }

    listener, err := net.Listen(protocol, ":"+self.Port)
    errutil.HandleError(err)
    fmt.Printf("Node %s is listening on port %s\n", id, self.Port)

    neighbors := configuration.ChooseRandNeighbors(self, neighborCount)
    printNeighbors(neighbors)

    defer listener.Close()
    for {
        conn, err := listener.Accept()
        errutil.HandleError(err)
        go handleConnection(conn, self, neighbors)
    }
}

func getIdFromArgs() string {
    args := os.Args[1:]
    if len(args) == 0 {
        log.Fatal("Usage: ueb01.exe <id>")
    }
    return args[0]
}

func printNeighbors(neighbors []*conf.Node) {
    fmt.Print("Neighbors: ")
    end := len(neighbors) - 1
    for i := range neighbors {
        fmt.Printf("%v", *neighbors[i])
        if i != end {
            fmt.Printf(", ")
        } else {
            fmt.Println()
        }
    }
}

func handleConnection(conn net.Conn, self *conf.Node, neighbors []*conf.Node) {
    defer conn.Close()
    buf := make([]byte, 1024)
    n, err := conn.Read(buf)
    errutil.HandleError(err)
    response := strings.TrimSpace(string(buf[:n]))
    log.Printf("Received: %s\n", response)

    if strings.HasPrefix(response, "control:") {
        if strings.HasSuffix(response, "start") {
            sendMessages(self, neighbors)
        } else if response == "exit" {
            os.Exit(2)
        }
    } else {
        sendMessages(self, neighbors)
    }
}

func sendMessages(self *conf.Node, neighbors []*conf.Node) {
    directory.mu.Lock()
    defer directory.mu.Unlock()
    for i := range neighbors {
        if called, ok := directory.v[i]; !ok || !called {
            node := neighbors[i]
            conn, err := net.Dial(protocol, node.CreateAddress())
            if err != nil {
                log.Printf("Could not connect to node %s", node.Id)
            } else {
                payload := []byte(self.Id)
                _, err := conn.Write(payload)
                errutil.HandleError(err)
                directory.v[i] = true
                log.Printf("Sent %s to node %s\n", payload, node.Id)
            }
        }
    }
}
