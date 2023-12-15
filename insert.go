package bst

import (
	"errors"
	"fmt"
)

var (
	ErrDuplicatedNodeValue = errors.New("duplicated node value")
)

func (b *bst) Insert(value NodeValue) error {
	b.mux.Lock()
	defer b.mux.Unlock()

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

	curNode := b.root
	newValue := value.Value()

	for {
		curValue := curNode.Value()

		if newValue > curValue {
			if curNode.right == nil {
				curNode.right = newNode
				break
			}

			curNode = curNode.right
		} else if newValue < curValue {
			if curNode.left == nil {
				curNode.left = newNode
				break
			}

			curNode = curNode.left
		} else {
			return fmt.Errorf("%w: %d", ErrDuplicatedNodeValue, newValue)
		}
	}

	return nil
}
