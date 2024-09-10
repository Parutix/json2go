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
		fieldType := ""
		fieldName, _ := UpperFirst(key)

		if parser.IsPrimitive(value) {
			fieldType = parser.DetectType(value)
			sb.WriteString(fmt.Sprintf("\t%s %s `json:\"%s\"`\n", fieldName, fieldType, key))
		} else {
			switch v := value.(type) {
			case map[string]interface{}:
				sb.WriteString(fmt.Sprintf("\t%s struct {\n", fieldName))
				ContinueStruct(&sb, v)
				sb.WriteString(fmt.Sprintf("\t} `json:\"%s\"`\n", key))
			case []interface{}:
				if len(v) > 0 {
					if firstElem, ok := v[0].(map[string]interface{}); ok {
						sb.WriteString(fmt.Sprintf("\t%s []struct {\n", fieldName))
						ContinueStruct(&sb, firstElem)
						sb.WriteString(fmt.Sprintf("\t} `json:\"%s\"`\n", key))
					} else {
						fieldType = parser.DetectType(v[0])
						sb.WriteString(fmt.Sprintf("\t%s []%s `json:\"%s\"`\n", fieldName, fieldType, key))
					}
				}
			default:
				fmt.Printf("Unhandled type for key: %s, value: %v\n", key, value)
			}
		}
	}
	sb.WriteString("}\n")
	return sb.String()
}

func ContinueStruct(sb *strings.Builder, jsonData map[string]interface{}) {
	for key, value := range jsonData {
		fieldType := ""
		fieldName, _ := UpperFirst(key)

		if parser.IsPrimitive(value) {
			fieldType = parser.DetectType(value)
			sb.WriteString(fmt.Sprintf("\t%s %s `json:\"%s\"`\n", fieldName, fieldType, key))
		} else {
			switch v := value.(type) {
			case map[string]interface{}:
				sb.WriteString(fmt.Sprintf("\t%s struct {\n", fieldName))
				ContinueStruct(sb, v)
				sb.WriteString(fmt.Sprintf("\t} `json:\"%s\"`\n", key))
			case []interface{}:
				if len(v) > 0 {
					if firstElem, ok := v[0].(map[string]interface{}); ok {
						sb.WriteString(fmt.Sprintf("\t%s []struct {\n", fieldName))
						ContinueStruct(sb, firstElem)
						sb.WriteString(fmt.Sprintf("\t} `json:\"%s\"`\n", key))
					} else {
						fieldType = parser.DetectType(v[0])
						sb.WriteString(fmt.Sprintf("\t%s []%s `json:\"%s\"`\n", fieldName, fieldType, key))
					}
				}
			default:
				fmt.Printf("Unhandled type for key: %s, value: %v\n", key, value)
			}
		}
	}
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