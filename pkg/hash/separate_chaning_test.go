package hash

import (
	"reflect"
	"testing"
)

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

func TestSeparateChaningHashDel(t *testing.T) {
	successTests := []struct {
		key      stringHashKey
		expectOk bool
	}{
		{stringHashKey("abc"), true},
		{stringHashKey("wlefkmwelmkl"), false},
		{stringHashKey("elkfmwlmf2349u98sufdjs"), false},
		{stringHashKey("exsit"), true},
		{stringHashKey("happy-birth-day...!!!"), true},
	}

	sch := NewSeparateChainingHash()
	sch.Put(stringHashKey("abc"), "abc")
	sch.Put(stringHashKey("wlefkmwelmkl"), "abc")
	sch.Put(stringHashKey("elkfmwlmf2349u98sufdjs"), "abc")
	sch.Put(stringHashKey("exsit"), "abc")
	sch.Put(stringHashKey("happy-birth-day...!!!"), "abc")

	sch.Del(stringHashKey("abc"))
	sch.Del(stringHashKey("wlefkmwelmkl"))
	sch.Del(stringHashKey("elkfmwlmf2349u98sufdjs"))

	sch.Put(stringHashKey("abc"), "abc")

	for _, test := range successTests {
		key, expectOk := test.key, test.expectOk
		if _, ok := sch.Get(key); ok != expectOk {
			t.Errorf("hash delete error")
			t.Errorf("Expect: %+v, got: %+v", expectOk, ok)
		}
	}
}
