package dummy

type TreeKey struct {
	Key int
}

func NewTreeKey(key int) TreeKey {
	return TreeKey{key}
}

func (n TreeKey) Less(a interface{}) bool {
	return n.Key < a.(TreeKey).Key
}
