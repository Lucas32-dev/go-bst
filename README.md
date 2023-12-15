# Golang Binary Search Tree

### Overview

The Binary Search Tree (BST) is a node-based data structure for fast insertion, deletion, and search operations. The binary tree follows three rules:

- Each node has at most 2 children.
- The left child node must be lower than the parent node.
- The right child node must be greater than the parent node.

![Tree](https://github.com/Lucas32-dev/go-bst/blob/main/image.jpeg?raw=true)

### Time complexity in big O notation
| Operation | Average | Worst |
| --------- | ------- | ----- |
| Search    | O(log n)| O(n)  |
| Insert    | O(log n)| O(n)  |
| Delete    | O(log n)| O(n)  |

## Getting started

Initialize a new empty tree:

```go
b := bst.New()
```
Insert node:

```go
type MyBSTValue struct {
    A int
}

func (v MyValue) Value() int {
    return v.A
}

b.Insert(MyValue{42})
```

Delete node:

```go
b.Delete(42)
```

Search for node:

```go
value, ok := b.Search(42)
if ok {
    fmt.Println("Found: ", value)
} else {
    fmt.Println("Not found")
}
```