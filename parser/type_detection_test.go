package parser

import (
	"testing"
)

func TestDetectType_Succes(t *testing.T) {
	jsonData := []byte(`{
		"name": "John Doe",
		"age": 30,
		"email": "john.doe@example.com",
		"address": {
			"street": "123 Main St",
			"city": "New York",
			"zipcode": "10001"
		},
		"phoneNumbers": [
			{
				"type": "home",
				"number": "212-555-1234"
			},
			{
				"type": "work",
				"number": "646-555-4567"
			}
		],
		"isActive": true,
		"preferences": {
			"newsletter": false,
			"notifications": ["email", "sms"]
		},
		"projects": [
			{
				"name": "Project A",
				"budget": 5000.75,
				"completed": false
			},
			{
				"name": "Project B",
				"budget": 12000.50,
				"completed": true
			}
		]
	}`)
		result, _ := ParseJSON(jsonData)
		for key, value := range result {
			if key == "" {
				t.Error("Expected key to be not empty")
			}
			if value == nil {
				t.Error("Expected value to be not nil")
			}
			if DetectType(value) == "unknown" {
				t.Error("Expected type to be known")
			}
		}
}