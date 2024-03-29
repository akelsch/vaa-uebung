package config

import (
    "fmt"
    "github.com/akelsch/vaa/ueb02/internal/util/errutil"
    "github.com/akelsch/vaa/ueb02/internal/util/fileutil"
    "github.com/awalterschulze/gographviz"
    "math/rand"
)

type Config struct {
    all       []Node
    Self      *Node
    Neighbors []*Node
    Params    struct {
        T    int
        S    int
        P    int
        AMax int
    }
}

func NewConfig(filename string, id string) *Config {
    c := &Config{}
    for _, row := range fileutil.ReadCsvRows(filename) {
        c.all = append(c.all, Node{
            Id:   row[0],
            Host: row[1],
            Port: row[2],
        })
    }

    self, err := c.find(id)
    errutil.HandleError(err)
    c.Self = self

    return c
}

func (c *Config) find(id string) (*Node, error) {
    for i := range c.all {
        if c.all[i].Id == id {
            return &c.all[i], nil
        }
    }

    return nil, fmt.Errorf("could not find configuration for entry with ID %s", id)
}

func (c *Config) ChooseNeighborsByGraph(filename string) {
    graphAst, err := gographviz.Parse(fileutil.ReadBytes(filename))
    errutil.HandleError(err)

    graph := gographviz.NewGraph()
    err = gographviz.Analyse(graphAst, graph)
    errutil.HandleError(err)

    for _, edge := range graph.Edges.Edges {
        l := edge.Src
        r := edge.Dst
        if l == c.Self.Id {
            node, err := c.find(r)
            errutil.HandleError(err)
            c.Neighbors = append(c.Neighbors, node)
        } else if r == c.Self.Id {
            node, err := c.find(l)
            errutil.HandleError(err)
            c.Neighbors = append(c.Neighbors, node)
        }
    }
}

func (c *Config) NeighborsToString() string {
    lastIndex := len(c.Neighbors) - 1

    output := "Neighbors: "
    for i := range c.Neighbors {
        output += fmt.Sprintf("%v", c.Neighbors[i].Id)
        if i != lastIndex {
            output += ", "
        }
    }

    return output
}

func (c *Config) RegisterAllAsNeighbors() {
    c.Neighbors = nil
    for i := range c.all {
        neighbor := &c.all[i]
        if neighbor != c.Self {
            c.Neighbors = append(c.Neighbors, neighbor)
        }
    }
}

func (c *Config) GetRandomNeighbors(n int) []*Node {
    var neighbors []*Node
    for _, randIndex := range rand.Perm(len(c.Neighbors)) {
        neighbors = append(neighbors, c.Neighbors[randIndex])
        if len(neighbors) >= n {
            break
        }
    }
    return neighbors
}

func (c *Config) FindNeighborById(id string) (int, *Node) {
    for i := range c.Neighbors {
        if c.Neighbors[i].Id == id {
            return i, c.Neighbors[i]
        }
    }

    return -1, nil
}
