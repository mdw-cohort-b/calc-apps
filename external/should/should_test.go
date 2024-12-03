package should_test

import (
	"testing"
	"time"

	"github.com/mdw-cohort-b/calc-apps/external/should"
)

func pass(t *testing.T, actual any, assertion should.Assertion, expected ...any) {
	t.Helper()
	f := NewFakeT()
	should.So(f, actual, assertion, expected...)
	if f.failure != nil {
		t.Error("expected passing Assertion, but if failed:", f.failure)
	}
}
func fail(t *testing.T, actual any, assertion should.Assertion, expected ...any) {
	t.Helper()
	f := NewFakeT()
	should.So(f, actual, assertion, expected...)
	if f.failure == nil {
		t.Error("expected failing Assertion, but if passed!")
	}
}
func TestShouldEqual(t *testing.T) {
	pass(t, 1, should.Equal, 1)
	pass(t, "", should.Equal, "")
	pass(t, "a", should.Equal, "a")
	//pass(t, uint8(1), should.Equal, uint16(1)) // bonus points...

	fail(t, "a", should.Equal, "a ")
	fail(t, 1, should.Equal, 2)
}
func TestShouldBeTrue(t *testing.T) {
	pass(t, true, should.BeTrue)
	fail(t, false, should.BeTrue)
}
func TestShouldBeFalse(t *testing.T) {
	fail(t, true, should.BeFalse)
	pass(t, false, should.BeFalse)
}
func TestShouldBeNil(t *testing.T) {
	fail(t, true, should.BeNil)
	fail(t, false, should.BeNil)
	fail(t, 1, should.BeNil)
	pass(t, nil, should.BeNil)
	pass(t, (*time.Time)(nil), should.BeNil) // bonus points...
}
func TestShouldNotEqual(t *testing.T) {
	pass(t, 1, should.NOT.Equal, 2)
	fail(t, 1, should.NOT.Equal, 1)
}
func TestShouldNotBeNil(t *testing.T) {
	pass(t, 1, should.NOT.BeNil)
	fail(t, nil, should.NOT.BeNil)
}

/////////////////////////////////////////////////////////////

type FakeT struct{ failure error }

func NewFakeT() *FakeT { return &FakeT{} }

func (this *FakeT) Helper() {}
func (this *FakeT) Error(a ...any) {
	this.failure = a[0].(error)
}
