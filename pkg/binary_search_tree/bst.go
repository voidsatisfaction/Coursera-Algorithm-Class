package binarySearchTree

type node struct {
	key   Key
	value interface{}
	left  *node
	right *node
}

type Key interface {
	Less(key interface{}) bool
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
	n := bst.Root
	newNode := &node{key, value, nil, nil}

	if n == nil {
		bst.Root = newNode
		return
	}

	for {
		if key.Less(n.key) {
			if n.left == nil {
				n.left = newNode
				return
			}
			n = n.left
		} else if n.key.Less(key) {
			if n.right == nil {
				n.right = newNode
				return
			}
			n = n.right
		} else {
			break
		}
	}
}
