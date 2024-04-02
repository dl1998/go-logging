package parser

import (
	"encoding/json"
	"fmt"
	"github.com/dl1998/go-logging/internal/testutils"
	"testing"
)

var (
	testFile = "test.json"
)

// TestReadFromFile tests that readFromFile reads the file correctly.
func TestReadFromFile(t *testing.T) {
	readFile = func(_ string) ([]byte, error) {
		return []byte(`{"loggers": [{"name": "test"}]}`), nil
	}

	expected := &Configuration{
		Loggers: []LoggerConfiguration{
			{
				Name: "test",
			},
		},
	}

	configuration, err := readFromFile("test.json", json.Unmarshal)

	testutils.AssertNil(t, err)
	testutils.AssertEquals(t, expected, configuration)
}

// TestReadFromFile_Error tests that readFromFile return nil and error, when
// error occurs while reading the file or unmarshalling the data from the file to
// the configuration struct.
func TestReadFromFile_Error(t *testing.T) {
	tests := map[string]struct {
		readFile  func(string) ([]byte, error)
		unmarshal func([]byte, any) error
	}{
		"Failed To Read": {
			readFile: func(_ string) ([]byte, error) {
				return nil, fmt.Errorf("test error")
			},
			unmarshal: json.Unmarshal,
		},
		"Failed To Unmarshal": {
			readFile: func(_ string) ([]byte, error) {
				return []byte(`{"loggers": [{"name": "test"}]}`), nil
			},
			unmarshal: func(_ []byte, _ any) error {
				return fmt.Errorf("test error")
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			readFile = test.readFile

			configuration, err := readFromFile("test.json", test.unmarshal)

			testutils.AssertNotNil(t, err)
			testutils.AssertNil(t, configuration)
		})
	}
}

// BenchmarkReadFromFile benchmarks the readFromFile function.
func BenchmarkReadFromFile(b *testing.B) {
	readFile = func(_ string) ([]byte, error) {
		return []byte(`{"loggers": [{"name": "test"}]}`), nil
	}

	for index := 0; index < b.N; index++ {
		_, _ = readFromFile(testFile, json.Unmarshal)
	}
}

// TestReadFromJSON tests that ReadFromJSON reads the JSON file correctly.
func TestReadFromJSON(t *testing.T) {
	readFile = func(_ string) ([]byte, error) {
		return []byte(`{"loggers": [{"name": "test"}]}`), nil
	}

	expected := &Configuration{
		Loggers: []LoggerConfiguration{
			{
				Name: "test",
			},
		},
	}

	configuration, err := ReadFromJSON("test.json")

	testutils.AssertNil(t, err)
	testutils.AssertEquals(t, expected, configuration)
}

// BenchmarkReadFromJSON benchmarks the ReadFromJSON function.
func BenchmarkReadFromJSON(b *testing.B) {
	readFile = func(_ string) ([]byte, error) {
		return []byte(`{"loggers": [{"name": "test"}]}`), nil
	}

	for index := 0; index < b.N; index++ {
		_, _ = ReadFromJSON(testFile)
	}
}

// TestReadFromYAML tests that ReadFromYAML reads the YAML file correctly.
func TestReadFromYAML(t *testing.T) {
	readFile = func(_ string) ([]byte, error) {
		return []byte("loggers:\n- name: test"), nil
	}

	expected := &Configuration{
		Loggers: []LoggerConfiguration{
			{
				Name: "test",
			},
		},
	}

	configuration, err := ReadFromYAML("test.yaml")

	testutils.AssertNil(t, err)
	testutils.AssertEquals(t, expected, configuration)
}

// BenchmarkReadFromYAML benchmarks the ReadFromYAML function.
func BenchmarkReadFromYAML(b *testing.B) {
	readFile = func(_ string) ([]byte, error) {
		return []byte("loggers:\n- name: test"), nil
	}

	for index := 0; index < b.N; index++ {
		_, _ = ReadFromYAML(testFile)
	}
}

// TestReadFromXML tests that ReadFromXML reads the XML file correctly.
func TestReadFromXML(t *testing.T) {
	readFile = func(_ string) ([]byte, error) {
		return []byte("<root><loggers><logger><name>test</name></logger></loggers></root>"), nil
	}

	expected := &Configuration{
		Loggers: []LoggerConfiguration{
			{
				Name: "test",
			},
		},
	}

	configuration, err := ReadFromXML("test.xml")

	testutils.AssertNil(t, err)
	testutils.AssertEquals(t, expected, configuration)
}

// BenchmarkReadFromXML benchmarks the ReadFromXML function.
func BenchmarkReadFromXML(b *testing.B) {
	readFile = func(_ string) ([]byte, error) {
		return []byte("<loggers><logger><name>test</name></logger></loggers>"), nil
	}

	for index := 0; index < b.N; index++ {
		_, _ = ReadFromXML(testFile)
	}
}
