package bst

func (b *bst) Delete(value int) bool {
	b.mux.Lock()
	defer b.mux.Unlock()

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

	if n.left == nil && n.right == nil { // no child, replace node with nil node
		b.replaceNode(parent, n, nil)
	} else if n.left != nil && n.right == nil { // have only left child, replace node with left node
		b.replaceNode(parent, n, n.left)
	} else if n.left == nil && n.right != nil { // have only right child, replace node with right node
		b.replaceNode(parent, n, n.right)
	} else { // have two childs, replace node with min node from right tree
		rightNode := n.right

		if rightNode.left == nil { // check if right node is the min node
			b.replaceNode(parent, n, rightNode)
			rightNode.left = n.left
		} else { // find min node in subtree
			minParent, minNode := b.getMinNode(rightNode)

			b.replaceNode(minParent, minNode, minNode.right)
			b.replaceNode(parent, n, minNode)

			minNode.left = n.left
			minNode.right = n.right
		}
	}

	return true
}

func (b *bst) getMinNode(startNode *node) (minParent *node, minNode *node) {
	minParent = startNode
	minNode = startNode.left

	// if left is nil, there is no min node
	// for the start node
	if minNode == nil {
		return minParent, nil
	}

	for minNode.left != nil {
		minParent = minNode
		minNode = minNode.left
	}

	return minParent, minNode
}

func (b *bst) replaceNode(parent *node, oldNode *node, newNode *node) {
	if parent == nil {
		b.root = newNode
	} else if parent.left != nil && parent.left.Value() == oldNode.Value() {
		parent.left = newNode
	} else {
		parent.right = newNode
	}
}
