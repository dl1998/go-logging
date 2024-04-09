// Package testutils contains utility methods for the testing.
package testutils

import (
	"path"
	"reflect"
	"runtime"
	"testing"
)

var (
	TEST_DATA_PATH string
)

func init() {
	_, fileName, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	currentPath := path.Dir(fileName)

	projectRoot := path.Dir(path.Dir(currentPath))

	TEST_DATA_PATH = path.Join(projectRoot, "testdata")
}

// AssertEquals compares actual value with expected value.
func AssertEquals[T any](t *testing.T, expected T, actual T) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\nExpected: %v\nActual: %v", expected, actual)
	}
}

func isNil(value interface{}) bool {
	if value == nil {
		return true
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
		return reflect.ValueOf(value).IsNil()
	default:
		return false
	}
}

// AssertNil checks if the value is nil.
func AssertNil(t *testing.T, value any) {
	t.Helper()
	if !isNil(value) {
		t.Fatalf("\nExpected: nil\nActual: %v", value)
	}
}

// AssertNotNil checks if the value is not nil.
func AssertNotNil(t *testing.T, value any) {
	t.Helper()
	if isNil(value) {
		t.Fatalf("\nExpected: not nil\nActual: %v", value)
	}
}
