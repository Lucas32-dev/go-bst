package bst

import "errors"

var (
	ErrDuplicatedNodeValue = errors.New("invalid node value, node value already exists")
)

func (b *bst) Insert(value NodeValue) error {
	err := b.insertNode(value)
	if err != nil {
		return err
	}

	b.len++
	return nil
}

func (b *bst) insertNode(value NodeValue) error {
	newNode := &node{
		value: value,
	}

	if b.root == nil {
		b.root = newNode
		return nil
	}

	v := value.Value()
	n := b.root

	var nv int
	for {
		nv = n.Value()
		if v > nv {
			if n.right == nil {
				n.right = newNode
				break
			}

			n = n.right
		} else if v < nv {
			if n.left == nil {
				n.left = newNode
				break
			}

			n = n.left
		} else {
			return ErrDuplicatedNodeValue
		}
	}

	return nil
}
