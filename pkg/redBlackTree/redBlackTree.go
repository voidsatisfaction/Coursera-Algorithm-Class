package redBlackTree

type Key interface {
	Less(key interface{}) bool
}

const (
	RED   = true
	BLACK = false
)

type node struct {
	key   Key
	value interface{}
	color bool
	left  *node
	right *node
}

func isRed(n *node) bool {
	if n == nil {
		return BLACK
	}
	return n.color == RED
}

func rotateLeft(n *node) *node {
	x := n.right
	n.right = x.left
	x.left = n
	x.color = n.color
	n.color = RED
	return x
}

func rotateRight(n *node) *node {
	x := n.left
	n.left = x.right
	x.right = n
	x.color = n.color
	n.color = RED
	return x
}

func flipColors(n *node) {
	n.color = RED
	n.left.color = BLACK
	n.right.color = BLACK
}

type RedBlackTree struct {
	Root *node
}

func New() *RedBlackTree {
	return &RedBlackTree{}
}

func (rbt *RedBlackTree) Get(key Key) interface{} {
	n := rbt.Root
	for n != nil {
		if key.Less(n.key) {
			n = n.left
		} else if n.key.Less(key) {
			n = n.right
		} else {
			return n.value
		}
	}
	return nil
}

func (rbt *RedBlackTree) Include(key Key) bool {
	if rbt.Get(key) != nil {
		return true
	}
	return false
}

func (rbt *RedBlackTree) Insert(key Key, value interface{}) {
	var insert func(n *node, key Key, value interface{}) *node
	insert = func(n *node, key Key, value interface{}) *node {
		if n == nil {
			return &node{key, value, BLACK, nil, nil}
		}
		if key.Less(n.key) {
			n.left = insert(n.left, key, value)
		} else if n.key.Less(key) {
			n.right = insert(n.right, key, value)
		} else {
			n.value = value
		}

		if isRed(n.right) && !isRed(n.left) {
			n = rotateLeft(n)
		} else if isRed(n.left) && isRed(n.left.left) {
			n = rotateRight(n)
		} else if isRed(n.left) && isRed(n.right) {
			flipColors(n)
		}

		return n
	}

	n := rbt.Root
	rbt.Root = insert(n, key, value)
}
