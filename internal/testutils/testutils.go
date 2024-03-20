// Package testutils contains utility methods for the testing.
package testutils

import (
	"reflect"
	"testing"
)

// AssertEquals compares actual value with expected value.
func AssertEquals[T any](t *testing.T, expected T, actual T) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\nExpected: %v\nActual: %v", expected, actual)
	}
}

// AssertNil checks if the value is nil.
func AssertNil(t *testing.T, value any) {
	t.Helper()
	if value != nil {
		t.Fatalf("\nExpected: nil\nActual: %v", value)
	}
}

// AssertNotNil checks if the value is not nil.
func AssertNotNil(t *testing.T, value any) {
	t.Helper()
	if value == nil {
		t.Fatalf("\nExpected: not nil\nActual: %v", value)
	}
}
