package bst

import "strconv"

type NodeValue interface {
	Value() int
}

type node struct {
	value NodeValue
	left *node
	right *node
}

func (n node) Value() int {
	return n.value.Value()
}

func (n node) String() string {
	return strconv.Itoa(n.Value())
}