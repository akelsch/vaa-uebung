package config

import (
    "fmt"
    "github.com/akelsch/vaa/ueb03/internal/util/convutil"
    "github.com/akelsch/vaa/ueb03/internal/util/errutil"
    "github.com/akelsch/vaa/ueb03/internal/util/fileutil"
    "github.com/awalterschulze/gographviz"
    "math/rand"
)

type Config struct {
    all       []Node
    Self      *Node
    Neighbors []*Node
    Params    struct {
        Balance int
    }
}

func NewConfig(filename string, id uint64) *Config {
    c := &Config{}

    // all
    for _, row := range fileutil.ReadCsvRows(filename) {
        c.all = append(c.all, Node{
            Id:   convutil.StringToUint(row[0]),
            Host: row[1],
            Port: row[2],
        })
    }

    // Self
    c.Self = c.find(id)

    return c
}

func (c *Config) ChooseNeighborsByGraph(filename string) {
    graphAst, err := gographviz.Parse(fileutil.ReadBytes(filename))
    errutil.HandleError(err)

    graph := gographviz.NewGraph()
    err = gographviz.Analyse(graphAst, graph)
    errutil.HandleError(err)

    // Neighbors
    for _, edge := range graph.Edges.Edges {
        l := convutil.StringToUint(edge.Src)
        r := convutil.StringToUint(edge.Dst)
        if l == c.Self.Id {
            c.Neighbors = append(c.Neighbors, c.find(r))
        } else if r == c.Self.Id {
            c.Neighbors = append(c.Neighbors, c.find(l))
        }
    }
}

func (c *Config) find(id uint64) *Node {
    node := c.FindNodeById(id)

    if node == nil {
        errutil.HandleError(fmt.Errorf("could not find configuration for entry with ID %d", id))
    }

    return node
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

func (c *Config) FindNodeById(id uint64) *Node {
    for i := range c.all {
        node := &c.all[i]
        if node.Id == id {
            return node
        }
    }

    return nil
}

func (c *Config) FindRandomNode() *Node {
    for _, randIndex := range rand.Perm(len(c.all)) {
        node := &c.all[randIndex]
        if node != c.Self {
            return node
        }
    }

    return nil
}

func (c *Config) FindNeighborById(id uint64) *Node {
    for i := range c.Neighbors {
        neighbor := c.Neighbors[i]
        if neighbor.Id == id {
            return neighbor
        }
    }

    return nil
}
