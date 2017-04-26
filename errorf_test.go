package errorf

import "testing"

func TestDiscard(t *testing.T) {
	d := NewDiscard()
	d.Pf("hello %s", "there")
	d = WrapDiscard(nil)
	d.Pf("hello %s", "there").Pln("Do This")
}
