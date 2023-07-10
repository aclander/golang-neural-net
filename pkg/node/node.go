package node

import "fmt"

type Args struct {
	Label string
	Val   float64
	Grad  float64
}

type Node struct {
	Label     string
	val       float64
	Grad      float64
	formula   string
	Backwards func()
	Parents   []*Node
}

func NewRoot(args *Args) *Node {
	return &Node{
		Label:     args.Label,
		val:       args.Val,
		formula:   args.Label,
		Grad:      args.Grad,
		Backwards: func() {},
	}
}

func (n *Node) Print() {
	fmt.Printf("Label: %v\n", n.Label)
	fmt.Printf("Formula: %v\n", n.formula)
	fmt.Printf("Value: %.2f\n", n.val)
	fmt.Printf("Gradient: %.2f\n", n.Grad)
	fmt.Println("----------------------------")
}

func (n *Node) Add(other *Node, label string) *Node {
	out := &Node{
		Label:   label,
		val:     n.val + other.val,
		Parents: []*Node{n, other},
		formula: fmt.Sprintf("%v + %v", n.Label, other.Label),
	}
	out.Backwards = func() {
		n.Grad += out.Grad
		other.Grad += out.Grad
	}
	return out
}

func (n *Node) Multiply(other *Node, label string) *Node {
	out := &Node{
		Label:   label,
		val:     n.val * other.val,
		Parents: []*Node{n, other},
		formula: fmt.Sprintf("%v * %v", n.Label, other.Label),
	}
	out.Backwards = func() {
		n.Grad += other.val * out.Grad
		other.Grad += n.val * out.Grad
	}
	return out
}
