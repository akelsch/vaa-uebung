package config

import "fmt"

type Node struct {
    Id   uint64
    Host string
    Port string
}

func (n *Node) GetListenAddress() string {
    return fmt.Sprintf(":%s", n.Port)
}

func (n *Node) GetDialAddress() string {
    return fmt.Sprintf("%s:%s", n.Host, n.Port)
}
