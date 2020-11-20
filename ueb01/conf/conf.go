package conf

import (
    "fmt"
    "github.com/akelsch/vaa/ueb01/csvutil"
    "math/rand"
    "time"
)

type Config struct {
    All  []Node
}

type Node struct {
    Id   string
    Host string
    Port string
}

func Init(filename string, id string) *Config {
    var c Config
    for _, row := range csvutil.Parse(filename) {
        c.All = append(c.All, Node{
            Id:   row[0],
            Host: row[1],
            Port: row[2],
        })
    }

    return &c
}

func (c *Config) Find(id string) *Node {
    for i := range c.All {
        if c.All[i].Id == id {
            return &c.All[i]
        }
    }

    return nil
}

func (c *Config) ChooseRandNeighbors(self *Node, n int) []*Node {
    rand.Seed(time.Now().UnixNano())

    var neighbors []*Node
    for _, randIndex := range rand.Perm(len(c.All)) {
        other := &c.All[randIndex]
        if len(neighbors) < n && other != self {
            neighbors = append(neighbors, other)
        }
    }

    return neighbors
}

func (n *Node) CreateAddress() string {
    return fmt.Sprintf("%s:%s", n.Host, n.Port)
}
