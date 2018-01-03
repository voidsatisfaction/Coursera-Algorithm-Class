package binarySearchTree

type Key interface {
	Less(key interface{}) bool
}

type node struct {
	key   Key
	value interface{}
	left  *node
	right *node
}

type BinarySearchTree struct {
	Root *node
}

func New() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (bst *BinarySearchTree) Get(key Key) interface{} {
	node := bst.Root
	for node != nil {
		if key.Less(node.key) {
			node = node.left
		} else if node.key.Less(key) {
			node = node.right
		} else {
			return node.value
		}
	}
	return nil
}

func (bst *BinarySearchTree) Include(key Key) bool {
	if bst.Get(key) != nil {
		return true
	}
	return false
}

func (bst *BinarySearchTree) Insert(key Key, value interface{}) {
	var insert func(n *node, key Key, value interface{}) *node
	insert = func(n *node, key Key, value interface{}) *node {
		if n == nil {
			newNode := &node{key, value, nil, nil}
			return newNode
		}
		if key.Less(n.key) {
			n.left = insert(n.left, key, value)
		} else if n.key.Less(key) {
			n.right = insert(n.right, key, value)
		} else {
			n.value = value
		}
		return n
	}

	n := bst.Root
	bst.Root = insert(n, key, value)
}
