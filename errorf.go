package errorf

type I interface {
	Pf(format string, args ...interface{}) I
	Pln(a ...interface{}) I
	Error() string
}
type Iwrap func(interface{}) I

type discard struct {
}

func NewDiscard() I {
	d := new(discard)
	return d
}
func (d discard) Pf(string, ...interface{}) I { return d }
func (d discard) Pln(...interface{}) I        { return d }
func (discard) Error() string                 { return "" }

func WrapDiscard(interface{}) I {
	return NewDiscard()
}
