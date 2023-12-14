package bst

import (
	"fmt"
	"strconv"
	"strings"
)

func (b bst) PrintValues() {
	sb := &strings.Builder{}
	b.inOrderTraversal(sb)
	
	fmt.Println(sb)
}

func (b bst) inOrderTraversal(sb *strings.Builder) {
	b.inOrderTraversalByNode(b.root, sb)
}

func (b bst) inOrderTraversalByNode(root *node, sb *strings.Builder) {
	if root == nil {
		return
	}

	b.inOrderTraversalByNode(root.left, sb)
	sb.WriteString(strconv.Itoa(root.Value()) + " ")
	b.inOrderTraversalByNode(root.right, sb)
}

