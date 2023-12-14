package bst

func (b *bst) Delete(value int) bool {
	if ok := b.deleteNode(value); ok {
		b.len--
		return true
	}
	return false
}

func (b *bst) deleteNode(value int) bool {
	var parent *node
	n := b.root

	for n != nil {
		if n.Value() == value {
			break
		}

		parent = n
		if value > n.Value() {
			n = n.right
		} else {
			n = n.left
		}
	}

	if n == nil {
		return false
	}

	if n.left == nil && n.right == nil {
		b.replaceNode(parent, n, nil)
	} else if n.left != nil && n.right == nil {
		b.replaceNode(parent, n, n.left)
	} else if n.left == nil && n.right != nil {
		b.replaceNode(parent, n, n.right)
	} else {
		minParent, minNode := b.getMinNode(n.right)

		b.replaceNode(minParent, minNode, minNode.right)
		b.replaceNode(parent, n, minNode)

		minNode.left = n.left
		minNode.right = n.right
	}

	return true
}

func (b *bst) getMinNode(startNode *node) (minParent *node, minNode *node) {
	minParent = startNode
	minNode = startNode.left

	for minNode.left != nil {
		minParent = minNode
		minNode = minNode.left
	}

	return minParent, minNode
}

func (b *bst) replaceNode(parent *node, oldNode *node, newNode *node) {
	if parent == nil {
		b.root = newNode
	} else if parent.left.Value() == oldNode.Value() {
		parent.left = newNode
	} else {
		parent.right = newNode
	}
}
