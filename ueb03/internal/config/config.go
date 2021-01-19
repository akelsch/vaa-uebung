package config

import (
    "fmt"
    "github.com/akelsch/vaa/ueb03/internal/util/errutil"
    "github.com/akelsch/vaa/ueb03/internal/util/fileutil"
    "github.com/awalterschulze/gographviz"
    "math/rand"
    "strconv"
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
        parsedId, err := strconv.ParseUint(row[0], 10, 0)
        errutil.HandleError(err)
        c.all = append(c.all, Node{
            Id:   parsedId,
            Host: row[1],
            Port: row[2],
        })
    }

    // Self
    self, err := c.find(id)
    errutil.HandleError(err)
    c.Self = self

    return c
}

func (c *Config) find(id uint64) (*Node, error) {
    for i := range c.all {
        if c.all[i].Id == id {
            return &c.all[i], nil
        }
    }

    return nil, fmt.Errorf("could not find configuration for entry with ID %d", id)
}

func (c *Config) ChooseNeighborsByGraph(filename string) {
    graphAst, err := gographviz.Parse(fileutil.ReadBytes(filename))
    errutil.HandleError(err)

    graph := gographviz.NewGraph()
    err = gographviz.Analyse(graphAst, graph)
    errutil.HandleError(err)

    for _, edge := range graph.Edges.Edges {
        l, err := strconv.ParseUint(edge.Src, 10, 0)
        errutil.HandleError(err)
        r, err := strconv.ParseUint(edge.Dst, 10, 0)
        errutil.HandleError(err)
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

func (c *Config) FindNeighborById(id uint64) (int, *Node) {
    for i := range c.Neighbors {
        if c.Neighbors[i].Id == id {
            return i, c.Neighbors[i]
        }
    }

    return -1, nil
}

// TODO remove once using flooding
func (c *Config) FindById(id uint64) (int, *Node) {
    for i := range c.all {
        if c.all[i].Id == id {
            return i, &c.all[i]
        }
    }

    return -1, nil
}

func (c *Config) FindRandomNode() (int, *Node) {
    for _, randIndex := range rand.Perm(len(c.all)) {
        node := &c.all[randIndex]
        if node != c.Self {
            return randIndex, node
        }
    }

    return -1, nil
}
