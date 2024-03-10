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
