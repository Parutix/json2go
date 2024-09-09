package parser

import (
	"encoding/json"
	"fmt"
)

func ParseJSON(jsonData []byte) (map[string]interface{}, error){
	var result map[string]interface{}
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return nil, fmt.Errorf("Error parsing JSON: %s", err.Error())
	}
	return result, nil
}