package generator

import (
	"fmt"
	"strings"

	"github.com/Parutix/json2go/parser"
)

func GenerateStruct(structName string, jsonData map[string]interface{}) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("type %s struct {\n", structName))
	
	for key, value := range jsonData {
		fieldName, err := UpperFirst(key)
		if err != nil {
			fieldName = key
		}

		fieldType := parser.DetectType(value)

		sb.WriteString(fmt.Sprintf("\t%s %s `json:\"%s\"`\n", fieldName, fieldType, key))
		
	}
	sb.WriteString("}\n")
	return sb.String()
}

func UpperFirst(s string) (string, error) {
	if len(s) == 0 {
		return "", fmt.Errorf("String is empty")
	}
	if s[0] >= 'a' && s[0] <= 'z' {
		return string(s[0] - 32) + s[1:], nil
	}

	return "", fmt.Errorf("First character is not a lowercase letter")
}