package neuralnet

import (
	"testing"

	"github.com/aclander/golang-neural-net/pkg/node"
	"github.com/google/go-cmp/cmp"
)

func TestBackpropogate(t *testing.T) {
	a := node.NewRoot(&node.Args{Label: "a", Val: 2.0})
	b := node.NewRoot(&node.Args{Label: "b", Val: 3.0})
	c := a.Add(b, "c")
	d := node.NewRoot(&node.Args{Label: "d", Val: 6.0})
	e := c.Multiply(d, "e")
	nodes := []*node.Node{b, c, d, a, e}
	nn := New(&Args{Nodes: nodes})
	nn.Backpropogate()
	wantOrder := []string{"e", "c", "a", "d", "b"}
	var gotOrder []string
	for _, n := range nn.nodes {
		gotOrder = append(gotOrder, n.Label)
	}
	if diff := cmp.Diff(wantOrder, gotOrder); diff != "" {
		t.Errorf("Backpropogate() incorrectly sorted nodes (-want, +got):\n%v", diff)
	}
}
