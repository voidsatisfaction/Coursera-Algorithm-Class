package hash

// M is size of hash table
const M int = 30001

// Key has to implement hashCode, equals metods
type Key interface {
	hashCode() int
	equals(hk interface{}) bool
}
