package neuralnet

import (
	"github.com/aclander/golang-neural-net/pkg/node"
)

type Args struct {
	Nodes []*node.Node
}

type NeuralNet struct {
	nodes []*node.Node
}

func New(args *Args) *NeuralNet {
	return &NeuralNet{nodes: args.Nodes}
}

func (nn *NeuralNet) Print() {
	for _, n := range nn.nodes {
		n.Print()
	}
}

func (n *NeuralNet) Backpropogate() {
	n.reverseTopoSort()
	for _, node := range n.nodes {
		node.Backwards()
	}
}

func (nn *NeuralNet) reverseTopoSort() {
	var sorted []*node.Node
	visited := make(map[*node.Node]bool)
	for len(visited) != len(nn.nodes) {
		for _, n := range nn.nodes {
			if visited[n] {
				continue
			}
			if n.Parents == nil || (visited[n.Parents[0]] && visited[n.Parents[1]]) {
				sorted = append([]*node.Node{n}, sorted...)
				visited[n] = true
			}
		}
	}
	nn.nodes = sorted
}
