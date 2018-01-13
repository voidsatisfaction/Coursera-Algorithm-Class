package hash

type HashKey interface {
	hashCode() int
	equals(hk interface{}) bool
}

type node struct {
	key HashKey

	val interface{}
}

const M int = 30001
