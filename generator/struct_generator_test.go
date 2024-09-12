package generator

import (
	"strings"
	"testing"

	"github.com/Parutix/json2go/formatter"
	"github.com/Parutix/json2go/parser"
)
		
func TestEmptyJSON(t *testing.T) {
	jsonData := []byte(`{}`)
	result, err := parser.ParseJSON(jsonData)
	if err != nil {
		t.Errorf("Expected error to be nil, got %v", err)
	}

	if result == nil {
		t.Error("Expected result to be not nil")
	}

	structName := "User"
	str := GenerateStruct(structName, result)
	expected := `type User struct {
	}`

	expected, _ = formatter.FormatStruct(expected)
	expected = strings.TrimSpace(expected)
	str, _ = formatter.FormatStruct(str)
	str = strings.TrimSpace(str)

	if str != expected {
		t.Errorf("Expected %s, got %s", expected, str)
	}
}
