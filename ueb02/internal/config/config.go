package config

import (
    "fmt"
    "github.com/akelsch/vaa/ueb02/internal/errutil"
    "github.com/akelsch/vaa/ueb02/internal/fileutil"
    "github.com/awalterschulze/gographviz"
    "log"
    "math/rand"
)

type Config struct {
    all       []Node
    Self      *Node
    Neighbors []*Node
    Params    struct {
        T int
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

func (c *Config) ChooseNeighborsRandomly(n int) {
    for _, randIndex := range rand.Perm(len(c.all)) {
        other := &c.all[randIndex]
        if len(c.Neighbors) < n && other != c.Self {
            c.Neighbors = append(c.Neighbors, other)
        }
    }
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

func (c *Config) PrintNeighbors() {
    lastIndex := len(c.Neighbors) - 1

    output := "Neighbors: "
    for i := range c.Neighbors {
        output += fmt.Sprintf("%v", c.Neighbors[i].Id)
        if i != lastIndex {
            output += ", "
        }
    }

    log.Println(output)
}
