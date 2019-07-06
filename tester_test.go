package mtest

import "testing"

type MyTestHelper struct {
	TestHelper
}

var s MyTestHelper

func TestA(t *testing.T) {
	s.AssertEqual(t, 1, 1)
}

func TestB(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			return
		} else {
			t.Error("does not fetch error...")
		}
	}()
	s.AssertEqual(t, 1, 2)
}
