package errorft

import (
	"fmt"
	"testing"

	"github.com/WhisperingChaos/errorf"
)

type T struct {
	t   *testing.T
	msg string
}

func New(t *testing.T) errorf.I {
	return newt(t)
}
func newt(t *testing.T) (e *T) {
	e = new(T)
	e.t = t
	return
}
func (t T) Pf(format string, args ...interface{}) errorf.I {
	// create race safe instance that will hold message so
	// Error will work
	tn := newt(t.t)
	tn.msg = fmt.Sprintf(format, args...)
	t.t.Errorf(format, args...)
	return tn
}
func (t T) Pln(a ...interface{}) errorf.I {
	tn := newt(t.t)
	tn.msg = fmt.Sprintln(a...)
	t.t.Error(a...)
	return tn
}
func (t T) Error() string {
	return t.msg
}
func Wrap(t *testing.T) errorf.Iwrap {
	return func(interface{}) errorf.I {
		return New(t)
	}
}
