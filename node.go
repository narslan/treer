package treer

import (
	_ "bytes"
	"fmt"
	_ "strconv"
	_ "strings"
)

// Tree is the representation of a single parsed template.
type Tree struct {
	Children []*Node
}

// New allocates a new parse tree.
func NewTree() *Tree {
	return &Tree{}
}

// A Node is an element in the parse tree.

type Node struct {
	Value string
	Tree  *Tree
}

// NewNode allocates a new node.
func newNode() *Node {
	return &Node{Value: ""}
}

func (t *Tree) Insert(v string) *Node {

	n := newNode()
	n.Value = v
	t.Children = append(t.Children, n)
	return n
}

func (n *Node) Insert(tr *Tree) *Tree {
	t := NewTree()
	n.Tree = t
	return t
}

func (t *Tree) Walker() error {

	for _, v := range t.Children {
		if v.Tree == nil {
			fmt.Printf(">%s", v.Value)
		} else {

			fmt.Printf("\n %s \t", v.Value)
			v.Tree.Walker()
		}
	}

	return nil
}
