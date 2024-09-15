package formatter

import (
	"go/format"
)

// FormatStruct formats a Go struct
func FormatStruct(code string) (string, error) {
	formatted, err := format.Source([]byte(code))
	if err != nil {
		return "", err
	}
	return string(formatted), nil
}

func CheckStartingDigit(s string) bool {
	if len(s) == 0 {
		return false
	}
	return s[0] >= '0' && s[0] <= '9'
}

func HandleStartingDigit(s string) string {
	return "_" + s
}