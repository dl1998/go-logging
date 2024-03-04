package testutils

import (
	"reflect"
	"testing"
)

func AssertEquals[T any](t *testing.T, expected T, actual T) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\nExpected: %v\nActual: %v", expected, actual)
	}
}
