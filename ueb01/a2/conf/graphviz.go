package conf;

import (
    "github.com/akelsch/vaa/ueb01/a2/errutil"
    "github.com/akelsch/vaa/ueb01/a2/fileutil"
    "github.com/awalterschulze/gographviz"
)

func (c *config) ChooseNeighborsByGraph(self *Node, filename string) []*Node {
    graphAst, err := gographviz.Parse(fileutil.ReadBytes(filename))
    errutil.HandleError(err)

    graph := gographviz.NewGraph()
    err = gographviz.Analyse(graphAst, graph)
    errutil.HandleError(err)

    var neighbors []*Node
    for _, edge := range graph.Edges.Edges {
        l := edge.Src
        r := edge.Dst
        if l == self.Id {
            node, err := c.Find(r)
            errutil.HandleError(err)
            neighbors = append(neighbors, node)
        } else if r == self.Id {
            node, err := c.Find(l)
            errutil.HandleError(err)
            neighbors = append(neighbors, node)
        }
    }

    return neighbors
}
