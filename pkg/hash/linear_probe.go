package hash

type Keys []Key

type Values []interface{}

// LPHash is hash with Linear probe approach
type LPHash struct {
	keys Keys
	vals Values
}

// NewLinearProbeHash makes new hash with Linear probe approach
func NewLinearProbeHash() *LPHash {
	return &LPHash{
		keys: make(Keys, M),
		vals: make(Values, M),
	}
}

func (lph *LPHash) Put(key Key, val interface{}) {
	var i int
	for i = key.hashCode() % M; lph.keys[i] != nil; i = (i + 1) % M {
		if lph.keys[i].equals(key) {
			lph.vals[i] = val
			return
		}
	}
	lph.keys[i] = key
	lph.vals[i] = val
}

func (lph *LPHash) Get(key Key) (interface{}, bool) {
	var i int
	for i = key.hashCode() % M; lph.keys[i] != nil; i = (i + 1) % M {
		if lph.keys[i].equals(key) {
			return lph.vals[i], true
		}
	}
	return nil, false
}
