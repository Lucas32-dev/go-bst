package bst

func (b bst) Search(value int) (NodeValue, bool) {
	node := b.searchNode(value)
	if node == nil {
		return nil, false
	}

	return node.value, true
}

func (b bst) searchNode(value int) *node {
	var nv int
	n := b.root

	for n != nil {
		nv = n.Value()

		if value > nv {
			n = n.right
		} else if value < nv {
			n = n.left
		} else {
			return n
		}
	}

	return nil
}
