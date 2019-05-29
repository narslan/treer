package main

import (
	_ "fmt"
	"github.com/narslan/node/node"
)

func main() {

	t1 := node.NewTree() //main tree
	t1.Insert("hello")
	t1.Insert("newroz")
	world := t1.Insert("world")

	t2 := node.NewTree() //main tree
	t3 := world.Insert(t2)
	t3.Insert("sub world1")
	t1.Walker()

}
