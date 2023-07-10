package main

import (
	"github.com/aclander/golang-neural-net/pkg/neuralnet"
	"github.com/aclander/golang-neural-net/pkg/node"
)

func main() {
	a := node.NewRoot(&node.Args{Label: "a", Val: 2.0})
	b := node.NewRoot(&node.Args{Label: "b", Val: 3.0})
	c := a.Add(b, "c")
	d := node.NewRoot(&node.Args{Label: "d", Val: 6.0})
	e := c.Multiply(d, "e")
	e.Grad = 1.0
	nodes := []*node.Node{a, b, c, d, e}
	nn := neuralnet.New(&neuralnet.Args{Nodes: nodes})
	nn.Backpropogate()
	nn.Print()
}
