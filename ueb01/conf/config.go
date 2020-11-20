package conf

import (
    "fmt"
    "github.com/akelsch/vaa/ueb01/csvutil"
    "math/rand"
    "time"
)

type config struct {
    nodes []Node
}

func NewConfig(filename string) *config {
    var c config
    for _, row := range csvutil.Parse(filename) {
        c.nodes = append(c.nodes, Node{
            Id:   row[0],
            Host: row[1],
            Port: row[2],
        })
    }

    return &c
}

func (c *config) Find(id string) (*Node, error) {
    for i := range c.nodes {
        if c.nodes[i].Id == id {
            return &c.nodes[i], nil
        }
    }

    return nil, fmt.Errorf("could not find configuration for entry with ID %s", id)
}

func (c *config) ChooseRandNeighbors(self *Node, n int) []*Node {
    rand.Seed(time.Now().UnixNano())

    var neighbors []*Node
    for _, randIndex := range rand.Perm(len(c.nodes)) {
        other := &c.nodes[randIndex]
        if len(neighbors) < n && other != self {
            neighbors = append(neighbors, other)
        }
    }

    return neighbors
}
