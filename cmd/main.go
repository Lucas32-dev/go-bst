package main

import (
	"bst"
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

	b.Delete(5)
	b.Delete(10)

	b.PrintValues()
}

func (v value) Value() int {
	return v.v
}

func newNodeValue(v int) bst.NodeValue {
	var nv bst.NodeValue = &value{v}
	return nv
}
