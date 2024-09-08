package parser

import (
	"encoding/json"
	"errors"
	"fmt"
)

func ParseJSON(jsonData []byte) (map[string]interface{}, error){
	var result map[string]interface{}
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return nil, errors.New(fmt.Sprintf("Error parsing JSON: %s", err.Error()))
	}
	return result, nil
}