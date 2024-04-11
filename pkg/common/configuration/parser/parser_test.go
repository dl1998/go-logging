package parser

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/dl1998/go-logging/internal/testutils"
	"testing"
)

var (
	testFile  = "test.json"
	xmlString = "<key>value</key>"
)

// TestEscapedString_escapeString tests that EscapedString.escapeString unescapes
// the string.
func TestEscapedString_escapeString(t *testing.T) {
	var escapedString EscapedString
	value := "\\t"

	expected := "\t"
	actual := escapedString.escapeString(value)

	testutils.AssertEquals(t, expected, actual)
}

// BenchmarkEscapedString_escapeString benchmarks the EscapedString.escapeString
// function.
func BenchmarkEscapedString_escapeString(b *testing.B) {
	escapedString := EscapedString("\\t")

	b.ResetTimer()

	for index := 0; index < b.N; index++ {
		_ = escapedString.escapeString("\\t")
	}
}

// TestEscapedString_UnmarshalXML tests that EscapedString.UnmarshalXML unmarshal
// the XML data correctly for escaped sequences.
func TestEscapedString_UnmarshalXML(t *testing.T) {
	var actual EscapedString

	err := xml.Unmarshal([]byte("<string>\\t</string>"), &actual)

	expected := EscapedString("\t")

	testutils.AssertNil(t, err)
	testutils.AssertEquals(t, expected, actual)
}

// BenchmarkEscapedString_UnmarshalXML benchmarks the EscapedString.UnmarshalXML
// function.
func BenchmarkEscapedString_UnmarshalXML(b *testing.B) {
	var escapedString EscapedString

	for index := 0; index < b.N; index++ {
		_ = xml.Unmarshal([]byte("<string>\\t</string>"), &escapedString)
	}
}

// TestKeyValue_UnmarshalXML tests that KeyValue.UnmarshalXML unmarshal the XML
// data correctly.
func TestKeyValue_UnmarshalXML(t *testing.T) {
	var expected = KeyValue{
		"key": "value",
	}
	var actual KeyValue

	decoder := xml.NewDecoder(bytes.NewReader([]byte(xmlString)))
	err := actual.UnmarshalXML(decoder, xml.StartElement{})

	testutils.AssertNil(t, err)
	testutils.AssertEquals(t, expected, actual)
}

// TestKeyValue_UnmarshalXML_Error tests that KeyValue.UnmarshalXML returns an
// error when it fails to unmarshal the XML data.
func TestKeyValue_UnmarshalXML_Error(t *testing.T) {
	decoder := xml.NewDecoder(bytes.NewReader([]byte("<start>broken element")))
	var keyValue KeyValue

	err := keyValue.UnmarshalXML(decoder, xml.StartElement{})

	testutils.AssertNotNil(t, err)
}

// BenchmarkKeyValue_UnmarshalXML benchmarks the KeyValue.UnmarshalXML function.
func BenchmarkKeyValue_UnmarshalXML(b *testing.B) {
	var keyValue KeyValue
	decoder := xml.NewDecoder(bytes.NewReader([]byte(xmlString)))

	for index := 0; index < b.N; index++ {
		_ = keyValue.UnmarshalXML(decoder, xml.StartElement{})
	}
}

// TestKeyValue_MarshalXML tests that KeyValue.MarshalXML marshals the XML data
// correctly.
func TestKeyValue_MarshalXML(t *testing.T) {
	buffer := bytes.Buffer{}
	encoder := xml.NewEncoder(&buffer)

	keyValue := &KeyValue{
		"key": "value",
	}

	start := xml.StartElement{Name: xml.Name{
		Space: "",
		Local: "KeyValue",
	}}

	err := keyValue.MarshalXML(encoder, start)

	testutils.AssertNil(t, err)

	_ = encoder.Close()

	testutils.AssertEquals(t, "<KeyValue><key>value</key></KeyValue>", buffer.String())
}

// TestKeyValue_MarshalXML_Error tests that KeyValue.MarshalXML returns an error
// when it fails to marshal the XML data.
func TestKeyValue_MarshalXML_Error(t *testing.T) {
	buffer := bytes.Buffer{}
	encoder := xml.NewEncoder(&buffer)

	keyValue := &KeyValue{
		"key": "value",
	}

	err := keyValue.MarshalXML(encoder, xml.StartElement{})

	_ = encoder.Close()

	// Assert that an error was returned
	testutils.AssertNotNil(t, err)
}

// TestKeyValue_MarshalXML_Error_Close tests that KeyValue.MarshalXML returns an
// error when it fails to close the encoder.
func TestKeyValue_MarshalXML_Error_Close(t *testing.T) {
	buffer := bytes.Buffer{}
	encoder := xml.NewEncoder(&buffer)

	keyValue := &KeyValue{
		"key": "value",
	}

	_ = encoder.Close()

	err := keyValue.MarshalXML(encoder, xml.StartElement{})

	testutils.AssertNotNil(t, err)
}

// BenchmarkKeyValue_MarshalXML benchmarks the KeyValue.MarshalXML function.
func BenchmarkKeyValue_MarshalXML(b *testing.B) {
	buffer := bytes.Buffer{}
	encoder := xml.NewEncoder(&buffer)

	keyValue := &KeyValue{
		"key": "value",
	}

	for index := 0; index < b.N; index++ {
		_ = keyValue.MarshalXML(encoder, xml.StartElement{})
	}
}

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
