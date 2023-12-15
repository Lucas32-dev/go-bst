package bst

func (b *bst) Search(value int) (NodeValue, bool) {
	b.mux.RLock()
	defer b.mux.RUnlock()

	node := b.searchNode(value)
	if node == nil {
		return nil, false
	}

	return node.value, true
}

func (b *bst) searchNode(value int) *node {
	var curValue int
	curNode := b.root

	for curNode != nil {
		curValue = curNode.Value()

		if value > curValue {
			curNode = curNode.right
		} else if value < curValue {
			curNode = curNode.left
		} else {
			return curNode
		}
	}

	return nil
}
