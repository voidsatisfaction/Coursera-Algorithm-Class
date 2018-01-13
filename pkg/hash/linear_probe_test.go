package hash

import (
	"reflect"
	"testing"
)

func TestLinearProbeHashTest(t *testing.T) {
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

	lph := NewLinearProbeHash()

	lph.Put(stringHashKey("abc"), "hello")
	lph.Put(stringHashKey("ccccsdcasd"), []int{1, 2, 3})
	lph.Put(stringHashKey("222"), "hihihi")
	lph.Put(stringHashKey("wkfm4io3fj3432341231zcxv9fknalkdsfl"), true)

	for _, test := range successTests {
		key, expectOk, expectVal := test.key, test.expectOk, test.expectVal
		val, ok := lph.Get(key)
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
		_, ok := lph.Get(key)

		if ok != expectOk {
			t.Errorf("ok is not as expected")
			t.Errorf("expectOk: %+v, got: %+v", expectOk, ok)
		}
	}
}
