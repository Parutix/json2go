package parser

import (
	"testing"
)

func TestParseJSON_Success(t *testing.T) {
	jsonData := []byte(
		`{
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

	result, err := ParseJSON(jsonData)
	if err != nil {
		t.Errorf("Expected error to be nil, got %v", err)
	}

	if result == nil {
		t.Error("Expected result to be not nil")
	}

	// Test Specific Fields
	if age, ok := result["age"].(float64); ok {
		if age != 30 {
			t.Errorf("Expected age to be 30, got %v", age)
		}
	} else {
		t.Errorf("Expected age to be float64, but got %T", result["age"])
	}
	
	if result["isActive"] != true {
		t.Errorf("Expected isActive to be true, got %v", result["isActive"])
	}
}