package main

import (
    "flag"
    "fmt"
    "github.com/akelsch/vaa/ueb01/a2/conf"
    "github.com/akelsch/vaa/ueb01/a2/errutil"
    "log"
    "net"
    "os"
    "strings"
)

const (
    protocol      = "tcp"
    controlPrefix = "c:"
)

var directory = *conf.NewNeighborDirectory()

func main() {
    file := flag.String("f", "./testdata/config.csv", "path to the CSV file containing the node configuration")
    gvFile := flag.String("gv", "./testdata/topology.gv", "path to the Graphviz file containing the network topology")
    id := flag.String("id", "1", "ID of the node")
    flag.Parse()

    // 1-2
    config := conf.NewConfig(*file)
    self, err := config.Find(*id)
    errutil.HandleError(err)

    // 3
    _, err = net.Listen(protocol, self.GetListenAddress())
    errutil.HandleError(err)
    fmt.Printf("Node %s is listening on port %s\n", self.Id, self.Port)

    // 4
    neighbors := config.ChooseNeighborsByGraph(self, *gvFile)
    printNeighbors(neighbors)

    // 5-9
    //defer listener.Close()
    //for {
    //    conn, err := listener.Accept()
    //    errutil.HandleError(err)
    //    go handleConnection(conn, self, neighbors)
    //}
}

func printNeighbors(neighbors []*conf.Node) {
    output := "Neighbors: "
    for i := range neighbors {
        output += fmt.Sprintf("%v", *neighbors[i])
        if i != len(neighbors)-1 {
            output += ", "
        }
    }
    fmt.Println(output)
}

func handleConnection(conn net.Conn, self *conf.Node, neighbors []*conf.Node) {
    defer conn.Close()
    buf := make([]byte, 1024)
    n, err := conn.Read(buf)
    errutil.HandleError(err)
    response := strings.TrimSpace(string(buf[:n]))
    log.Printf("Received: %s\n", response)

    // Control messages
    if strings.HasPrefix(response, controlPrefix) {
        command := strings.TrimPrefix(response, controlPrefix)
        if command == "start" {
            directory.Lock()
            directory.Reset()
            sendMessages(self, neighbors)
            directory.Unlock()
        } else if command == "exit" {
            // TODO exit all
            os.Exit(2)
        }
    }

    // Messages by other nodes
    for i := range neighbors {
        if response == neighbors[i].Id {
            directory.Lock()
            if directory.HasAlreadyReceivedFrom(i) {
                directory.ResetIfNecessary(len(neighbors))
            }
            directory.SetReceived(i)
            sendMessages(self, neighbors)
            directory.Unlock()
        }
    }
}

func sendMessages(self *conf.Node, neighbors []*conf.Node) {
    for i := range neighbors {
        if directory.HasNotSentTo(i) {
            neighbor := neighbors[i]
            conn, err := net.Dial(protocol, neighbor.GetDialAddress())
            if err != nil {
                log.Printf("Could not connect to node %s", neighbor.Id)
            } else {
                payload := []byte(self.Id)
                _, err := conn.Write(payload)
                errutil.HandleError(err)
                directory.SetSent(i)
                log.Printf("Sent %s to node %s\n", payload, neighbor.Id)
            }
        }
    }
}
