package parser

// DetectType detects the type of a value
func DetectType(value interface{}) string {
	switch value.(type) {
	case string:
		return "string"
	case float64:
		return "float64"
	case bool:
		return "bool"
	case int:
		return "int"
	case map[string]interface{}:
		return "map[string]interface{}"
	case []interface{}:
		return "[]interface{}"
	default:
		return "unknown"
	}
}

// IsPrimitive checks if a value is a primitive type
func IsPrimitive(value interface{}) bool {
	switch value.(type) {
	case string, float64, bool, int:
		return true
	default:
		return false
	}
}