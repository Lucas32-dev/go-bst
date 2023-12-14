package bst

func (b bst) Search(value int) (NodeValue, bool) {
	node :=  b.searchByNode(b.root, value)
	if node == nil {
		return nil, false
	}

	return node.value, true
}

func (b bst) searchByNode(root *node, value int) (*node) {
	if root == nil {
		return nil
	}

	rv := root.Value()
	if value > rv {
		return b.searchByNode(root.right, value)
	} else if value < rv {
		return b.searchByNode(root.left, value)
	} else {
		return root
	}
}