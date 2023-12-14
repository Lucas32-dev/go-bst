package bst

import "errors"

var (
	ErrDuplicatedNodeValue = errors.New("invalid node value, node value already exists")
)

func (b *bst) Insert(value NodeValue) error {
	root, err := b.insertByNode(b.root, value) 
	if err != nil {
		return err
	}

	b.root = root
	b.len++
	return nil
}

func (b *bst) insertByNode(root *node, value NodeValue) (*node, error) {
	if root == nil {
		return &node{
			value: value,
		}, nil
	}

	v := value.Value()
	rv := root.Value()

	if v > rv {
		n, err := b.insertByNode(root.right, value)
		if err != nil {
			return nil, err
		}

		root.right = n
	} else if v < rv {
		n, err := b.insertByNode(root.left, value)
		if err != nil {
			return nil, err
		}

		root.left = n
	} else {
		return nil, ErrDuplicatedNodeValue
	}

	return root, nil
}
