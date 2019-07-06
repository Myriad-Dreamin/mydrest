package mtest

type testcommon interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Helper()
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Name() string
	Skip(args ...interface{})
	SkipNow()
	Skipf(format string, args ...interface{})
	Skipped() bool
}

type TestHelper struct {
}

func (s *TestHelper) Assert(t testcommon, o interface{}, f func(interface{}) bool) {
	t.Helper()
	if !f(o) {
		t.Error("assertion failed")
		panic("assertion failed")
	}
	return
}

func (s *TestHelper) AssertEqual(t testcommon, a, b interface{}) {
	t.Helper()
	if a != b {
		t.Error("assertion failed")
		panic("assertion failed")
	}
}

func (s *TestHelper) AssertNoErr(t testcommon, o error) {
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
