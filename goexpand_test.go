package goexpand

import (
	"sort"
	"testing"
)

func testEq(t *testing.T, actual, expected []string) {
	if len(expected) != len(actual) {
		t.Error("expected", expected, "actual", actual)
		return
	}
	sort.Slice(actual, func(i, j int) bool {
		return actual[i] < expected[i]
	})
	sort.Slice(expected, func(i, j int) bool {
		return expected[i] < expected[i]
	})
	for i := range expected {
		if expected[i] != actual[i] {
			t.Error("expected", expected, "actual", actual)
			return
		}
	}
}

func TestExpand(t *testing.T) {
	testEq(t, Expand("foo"), []string{"foo"})
	testEq(t, Expand("foo,bar"), []string{"foo", "bar"})
	testEq(t, Expand("foo[0:2]bar"), []string{"foo0bar", "foo1bar", "foo2bar"})
	testEq(t, Expand("foo[00:02]"), []string{"foo00", "foo01", "foo02"})
	testEq(t, Expand("foo[0:1][0:1]"), []string{"foo00", "foo01", "foo10", "foo11"})
	testEq(t, Expand("foo[0:1][0:1],bar[00:01]"), []string{"foo00", "foo01", "foo10", "foo11", "bar00", "bar01"})
}
