package main

import (
    "fmt"
    "github.com/akelsch/vaa/ueb01/conf"
    "github.com/akelsch/vaa/ueb01/errutil"
    "log"
    "net"
    "os"
    "strings"
)

const (
    protocol      = "tcp"
    neighborCount = 3
    controlPrefix = "c:"
)

var directory = *conf.NewNeighborDirectory()

func main() {
    file, id := parseArgs()

    config := conf.NewConfig(file)
    self, err := config.Find(id)
    errutil.HandleError(err)

    listener, err := net.Listen(protocol, self.GetListenAddress())
    errutil.HandleError(err)
    fmt.Printf("Node %s is listening on port %s\n", id, self.Port)

    neighbors := config.ChooseRandNeighbors(self, neighborCount)
    printNeighbors(neighbors)

    defer listener.Close()
    for {
        conn, err := listener.Accept()
        errutil.HandleError(err)
        go handleConnection(conn, self, neighbors)
    }
}

func parseArgs() (string, string) {
    args := os.Args[1:]
    if len(args) != 2 {
        log.Fatal("Usage: ueb01.exe <file> <id>")
    }
    return args[0], args[1]
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

    if strings.HasPrefix(response, controlPrefix) {
        command := strings.TrimPrefix(response, controlPrefix)
        if command == "start" {
            sendMessages(self, neighbors)
        } else if command == "exit" {
            os.Exit(2)
        }
    } else {
        sendMessages(self, neighbors)
    }
}

func sendMessages(self *conf.Node, neighbors []*conf.Node) {
    directory.Lock()
    for i := range neighbors {
        if directory.IsRemaining(i) {
            neighbor := neighbors[i]
            conn, err := net.Dial(protocol, neighbor.GetDialAddress())
            if err != nil {
                log.Printf("Could not connect to node %s", neighbor.Id)
            } else {
                payload := []byte(self.Id)
                _, err := conn.Write(payload)
                errutil.HandleError(err)
                directory.Set(i)
                log.Printf("Sent %s to node %s\n", payload, neighbor.Id)
            }
        }
    }
    directory.Unlock()
}
