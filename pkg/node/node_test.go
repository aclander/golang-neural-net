package node

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestAdd(t *testing.T) {
	node := NewRoot(&Args{Label: "a", Val: 1.0})
	outGrad := 5.0
	tests := []struct {
		name          string
		other         *Node
		label         string
		wantNodeGrad  float64
		wantOtherGrad float64
	}{
		{
			name:          "add two nodes",
			other:         NewRoot(&Args{Label: "b", Val: 2.0}),
			label:         "c",
			wantNodeGrad:  outGrad,
			wantOtherGrad: outGrad,
		},
		{
			name:          "add node to self",
			other:         node,
			label:         "b",
			wantNodeGrad:  10.0,
			wantOtherGrad: 10.0,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			node.Grad = 0
			want := &Node{
				Label:   tc.label,
				val:     node.val + tc.other.val,
				Parents: []*Node{node, tc.other},
				formula: fmt.Sprintf("%v + %v", node.Label, tc.other.Label),
			}
			got := node.Add(tc.other, tc.label)
			if diff := cmp.Diff(want, got, cmp.AllowUnexported(Node{}), cmpopts.IgnoreFields(Node{}, "Backwards")); diff != "" {
				t.Fatalf("Add() returned unexpected diff (-want, +got):\n%v", diff)
			}
			got.Grad = outGrad
			got.Backwards()
			if (node.Grad != tc.wantNodeGrad) || (tc.other.Grad != tc.wantOtherGrad) {
				t.Errorf("got.Backwards() set unexpected grad values for parents:\nnode.Grad = %.2f, want = %.2f\nother.Grad = %.2f, want =%.2f", node.Grad, tc.wantNodeGrad, tc.other.Grad, tc.wantOtherGrad)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	node := NewRoot(&Args{Label: "a", Val: 2.0})
	outGrad := 5.0
	tests := []struct {
		name          string
		other         *Node
		label         string
		wantNodeGrad  float64
		wantOtherGrad float64
	}{
		{
			name:          "multiply two nodes",
			other:         NewRoot(&Args{Label: "b", Val: 3.0}),
			label:         "c",
			wantNodeGrad:  15.0,
			wantOtherGrad: 10.0,
		},
		{
			name:          "multiply same node",
			other:         node,
			label:         "b",
			wantNodeGrad:  20.0,
			wantOtherGrad: 20.0,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			node.Grad = 0
			want := &Node{
				Label:   tc.label,
				val:     node.val * tc.other.val,
				Parents: []*Node{node, tc.other},
				formula: fmt.Sprintf("%v * %v", node.Label, tc.other.Label),
			}
			got := node.Multiply(tc.other, tc.label)
			if diff := cmp.Diff(want, got, cmp.AllowUnexported(Node{}), cmpopts.IgnoreFields(Node{}, "Backwards")); diff != "" {
				t.Fatalf("Multiply() returned unexpected diff (-want, +got):\n%v", diff)
			}
			got.Grad = outGrad
			got.Backwards()
			if (node.Grad != tc.wantNodeGrad) || (tc.other.Grad != tc.wantOtherGrad) {
				t.Errorf("got.Backwards() set unexpected grad values for parents:\nnode.Grad = %.2f, want = %.2f\nother.Grad = %.2f, want =%.2f", node.Grad, tc.wantNodeGrad, tc.other.Grad, tc.wantOtherGrad)
			}
		})
	}
}
