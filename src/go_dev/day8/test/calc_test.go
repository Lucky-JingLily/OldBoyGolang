package test

import "testing"

func TestAdd(t *testing.T) {
	r := add(2, 4)
	if r != 6 {
		t.Fatalf("test 2+4 err,expect:6,act:%d", r)
	}
	t.Logf("test add success")
}
