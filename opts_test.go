package opts

import "testing"

type defaulter struct {
	Str   string
	Int   int
	Float float32
	Bool  bool
}

func (d defaulter) DefaultOptions() defaulter {
	return defaulter{
		Str:   "1",
		Int:   2,
		Float: 3.0,
	}
}

func withStr(s string) Opt[defaulter] {
	return func(o *defaulter) {
		o.Str = s
	}
}

func withBool(b bool) Opt[defaulter] {
	return func(o *defaulter) {
		o.Bool = b
	}
}

func withInt(i int) Opt[defaulter] {
	return func(o *defaulter) {
		o.Int = i
	}
}

func withFloat(f float32) Opt[defaulter] {
	return func(o *defaulter) {
		o.Float = f
	}
}

func TestDefaultApply(t *testing.T) {
	expected := defaulter{
		Str:   "1",
		Int:   5,
		Float: 2.0,
		Bool:  true,
	}
	actual := DefaultApply(withFloat(expected.Float), withInt(expected.Int), withBool(true))

	if expected != actual {
		t.Errorf("Expected %v but got %v", expected, actual)
		t.Fail()
	}
}
