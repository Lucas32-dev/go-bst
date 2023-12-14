package main

import (
	"bst"
	"fmt"
)

type value struct {
	v int
}

func main() {
	b := bst.New()

	b.Insert(newNodeValue(10))
	b.Insert(newNodeValue(5))
	b.Insert(newNodeValue(8))
	b.Insert(newNodeValue(7))
	b.Insert(newNodeValue(6))
	b.Insert(newNodeValue(9))
	b.Insert(newNodeValue(4))

	b.PrintValues()

	fmt.Println(b.Search(3))
	fmt.Println(b.Search(7))
}

func (v value) Value() int {
	return v.v
}

func newNodeValue(v int) bst.NodeValue {
	var nv bst.NodeValue = &value{v}
	return nv
}
