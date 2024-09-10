package formatter

import (
	"go/format"
)

func FormatStruct(code string) (string, error) {
	formatted, err := format.Source([]byte(code))
	if err != nil {
		return "", err
	}
	return string(formatted), nil
}