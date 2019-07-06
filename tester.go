package mtest

type TestInterface interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Helper()
}

type TestHelper struct {
}

func (s *TestHelper) Assert(t TestInterface, o interface{}, f func(interface{}) bool) {
	t.Helper()
	if !f(o) {
		t.Error("assertion failed")
		panic("assertion failed")
	}
	return
}

func (s *TestHelper) AssertEqual(t TestInterface, a, b interface{}) {
	t.Helper()
	if a != b {
		t.Error("assertion failed")
		panic("assertion failed")
	}
}

func (s *TestHelper) AssertTrue(t TestInterface, a bool) {
	t.Helper()
	if !a {
		t.Error("assertion failed")
		panic("assertion failed")
	}
}

func (s *TestHelper) AssertNoErr(t TestInterface, o error) {
	t.Helper()
	if o != nil {
		t.Error(o)
		panic("assertion failed")
	}
	return
}

func (s *TestHelper) OutAssert(o interface{}, f func(interface{}) bool) {
	if !f(o) {
		panic("assertion failed")
	}
	return
}

func (s *TestHelper) OutAssertNoErr(o error) {
	if o != nil {
		panic("assertion failed")
	}
	return
}

func IsNil(i interface{}) bool {
	return i == nil
}

func IsNotNil(i interface{}) bool {
	return i != nil
}

func IsTrue(i interface{}) bool {
	return i.(bool)
}

func IsFalse(i interface{}) bool {
	return i.(bool) == false
}

