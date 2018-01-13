package hash

type stringHashKey string

func (shk stringHashKey) hashCode() int {
	str := string(shk)
	var hc int
	for _, cp := range str {
		hc = 31*hc + int(cp)
	}

	return hc & (1<<31 - 1)
}

func (shk stringHashKey) equals(hk interface{}) bool {
	return string(shk) == string(hk.(stringHashKey))
}
