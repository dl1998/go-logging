package utils

import (
	"fmt"
	"github.com/dl1998/go-logging/internal/testutils"
	"net/http"
	"testing"
)

var (
	testStruct = struct {
		private      string
		Int          int
		Int8         int8
		Int16        int16
		Int32        int32
		Int64        int64
		Uint         uint
		Uint8        uint8
		Uint16       uint16
		Uint32       uint32
		Uint64       uint64
		Float32      float32
		Float64      float64
		Bool         bool
		String       string
		Header       http.Header
		CustomStruct CustomStruct
	}{
		private:      "private",
		Int:          -1,
		Int8:         -8,
		Int16:        -16,
		Int32:        -32,
		Int64:        -64,
		Uint:         1,
		Uint8:        8,
		Uint16:       16,
		Uint32:       32,
		Uint64:       64,
		Float32:      32.00,
		Float64:      64.00,
		Bool:         true,
		String:       "string",
		Header:       http.Header{"key": []string{"value"}},
		CustomStruct: CustomStruct{Field: "custom field"},
	}
)

type CustomStruct struct {
	Field string
}

func (customStruct *CustomStruct) String() string {
	return customStruct.Field
}

// TestStructToMap tests that StructToMap converts a struct to a map.
func TestStructToMap(t *testing.T) {
	expected := map[string]interface{}{
		"Int":          testStruct.Int,
		"Int8":         testStruct.Int8,
		"Int16":        testStruct.Int16,
		"Int32":        testStruct.Int32,
		"Int64":        testStruct.Int64,
		"Uint":         testStruct.Uint,
		"Uint8":        testStruct.Uint8,
		"Uint16":       testStruct.Uint16,
		"Uint32":       testStruct.Uint32,
		"Uint64":       testStruct.Uint64,
		"Float32":      testStruct.Float32,
		"Float64":      testStruct.Float64,
		"Bool":         testStruct.Bool,
		"String":       testStruct.String,
		"Header.key":   []string{"value"},
		"CustomStruct": fmt.Sprintf("{%s}", testStruct.CustomStruct.Field),
	}

	actual := StructToMap(testStruct)

	testutils.AssertNotNil(t, actual)
	testutils.AssertEquals(t, expected, actual)
}

// BenchmarkStructToMap benchmarks the StructToMap function.
func BenchmarkStructToMap(b *testing.B) {
	for index := 0; index < b.N; index++ {
		_ = StructToMap(testStruct)
	}
}
