package hash

import (
	"reflect"
	"testing"
)

type stringHashKey string

// hashCode should not mod M which is array size
// because it is moded at hash implement code
func (shk stringHashKey) hashCode() int {
	str := string(shk)
	var hc int
	for i := 0; i < len(str); i++ {
		hc = 31*hc + int(str[i])
	}
	return hc & (1<<31 - 1)
}

func (shk stringHashKey) equals(hk interface{}) bool {
	hk, ok := hk.(stringHashKey)
	if !ok {
		panic("hk is not converted to stringHashKey type")
	}
	return shk == hk
}

func TestSeparateChaningHashTest(t *testing.T) {
	// test for get success
	successTests := []struct {
		key       stringHashKey
		expectOk  bool
		expectVal interface{}
	}{
		{stringHashKey("abc"), true, "hello"},
		{stringHashKey("ccccsdcasd"), true, []int{1, 2, 3}},
		{stringHashKey("222"), true, "hihihi"},
		{stringHashKey("wkfm4io3fj3432341231zcxv9fknalkdsfl"), true, true},
	}

	sch := NewSeparateChainingHash()

	sch.Put(stringHashKey("abc"), "hello")
	sch.Put(stringHashKey("ccccsdcasd"), []int{1, 2, 3})
	sch.Put(stringHashKey("222"), "hihihi")
	sch.Put(stringHashKey("wkfm4io3fj3432341231zcxv9fknalkdsfl"), true)

	for _, test := range successTests {
		key, expectOk, expectVal := test.key, test.expectOk, test.expectVal
		val, ok := sch.Get(key)
		if !reflect.DeepEqual(expectVal, val) {
			t.Errorf("value is not as expected")
			t.Errorf("expectVal: %+v, got: %+v", expectVal, val)
		}

		if ok != expectOk {
			t.Errorf("ok is not as expected")
			t.Errorf("expectOk: %+v, got: %+v", expectOk, ok)
		}
	}

	// test for get fail
	failTests := []struct {
		key      stringHashKey
		expectOk bool
	}{
		{stringHashKey("abc11"), false},
		{stringHashKey("ccccsdcasasdfadd"), false},
		{stringHashKey("222123"), false},
	}

	for _, test := range failTests {
		key, expectOk := test.key, test.expectOk
		_, ok := sch.Get(key)

		if ok != expectOk {
			t.Errorf("ok is not as expected")
			t.Errorf("expectOk: %+v, got: %+v", expectOk, ok)
		}
	}
}
