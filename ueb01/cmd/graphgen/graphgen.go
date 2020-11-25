package main

import (
    "flag"
    "fmt"
    "github.com/akelsch/vaa/ueb01/internal/errutil"
    "github.com/awalterschulze/gographviz"
    "log"
    "math/rand"
    "strconv"
    "time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func main() {
    nNodes := flag.Int("n", 5, "number of nodes n")
    mEdges := flag.Int("m", 6, "number of edges m where m > n")
    flag.Parse()

    if *mEdges <= *nNodes {
        log.Fatal("The number of edges m must be greater than the number of nodes n")
    }

    graph := gographviz.NewGraph()

    for i := 1; i <= *nNodes; i++ {
        err := graph.AddNode("", strconv.Itoa(i), map[string]string{})
        errutil.HandleError(err)
    }

    for i := 2; i <= *nNodes; i++ {
        j := getRandomNumber(1, i-1)
        err := graph.AddEdge(strconv.Itoa(i), strconv.Itoa(j), false, map[string]string{})
        errutil.HandleError(err)
    }

    // TODO possibly endless
    for len(graph.Edges.Edges) < *mEdges {
        p := strconv.Itoa(getRandomNumber(1, *nNodes))
        q := strconv.Itoa(getRandomNumber(1, *nNodes))

        if isEdgeSuitable(graph, p, q) {
            err := graph.AddEdge(p, q, false, map[string]string{})
            errutil.HandleError(err)
        }
    }

    fmt.Print(graph)
}

func getRandomNumber(min, max int) int {
    return rand.Intn(max-min+1) + min
}

func isEdgeSuitable(graph *gographviz.Graph, p, q string) bool {
    edges := graph.Edges.Edges
    for i := range edges {
        edge := edges[i]
        // do not allow duplicates or loops
        if (edge.Dst == p && edge.Src == q) || (edge.Dst == q && edge.Src == p) || p == q {
            return false
        }
    }

    return true
}
