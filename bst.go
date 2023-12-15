package bst

import "sync"

type BinarySearchTree interface {
	// Insert adds a new node to the
	// binary search tree. Returns an
	// error if node already exists.
	Insert(NodeValue) error
	// Search searchs the tree for the
	// the value. Returns the node value
	// and it's existence.
	Search(int) (NodeValue, bool)
	// Delete deletes an node from the
	// tree. Returns true if deleted
	// and false if node not found.
	Delete(int) bool
	// Length returns the number of
	// nodes in the tree.
	Length() int
	// PrintValues prints the bst values
	// in traversal order.
	PrintValues()
}

type bst struct {
	// root is the tree's root node.
	root *node
	// mux is the read/write mutext
	mux sync.RWMutex
	// len is the number of nodes
	// in the binary search tree.
	len int
}

// New returns an empty binary search tree.
func New() BinarySearchTree {
	return &bst{}
}

func (b *bst) Length() int {
	b.mux.RLock()
	defer b.mux.RUnlock()

	return b.len
}
