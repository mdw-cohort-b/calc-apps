package should

import (
	"errors"
	"fmt"
	"reflect"
)

var errAssertionFailure = errors.New("assertion failure")

type testingT interface {
	Helper()
	Error(...any)
}

type Assertion func(actual any, expected ...any) error

func So(t testingT, actual any, assert Assertion, expected ...any) bool {
	t.Helper()
	err := assert(actual, expected...)
	if err != nil {
		t.Error(err)
	}
	return err == nil
}

func Equal(actual any, EXPECTED ...any) error {
	if reflect.DeepEqual(actual, EXPECTED[0]) {
		return nil
	}
	return fmt.Errorf("%w: got [%s] want [%s]", errAssertionFailure, actual, EXPECTED[0])
}
func BeTrue(actual any, _ ...any) error  { return Equal(actual, true) }
func BeFalse(actual any, _ ...any) error { return Equal(actual, false) }
func BeNil(actual any, _ ...any) error   { return Equal(actual, nil) }

type negated struct{}

var NOT negated

func (negated) Equal(actual any, expected ...any) error {
	err := Equal(actual, expected...)
	if err == nil {
		return fmt.Errorf("%w: expected values to not equal, but they did", errAssertionFailure)
	}
	return nil
}
func (negated) BeNil(actual any, _ ...any) error {
	err := BeNil(actual)
	if err == nil {
		return fmt.Errorf("%w: expected nil, got: %v", errAssertionFailure, actual)
	}
	return nil
}
