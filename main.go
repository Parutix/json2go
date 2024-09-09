package main

import (
	"fmt"

	"github.com/Parutix/json2go/generator"
	"github.com/Parutix/json2go/parser"
)

func main() {
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
	result, err := parser.ParseJSON(jsonData)
	if err != nil {
		panic(err)
	}
	for key,value := range result {
		fmt.Printf("Key: %s, Value: %v, Type: %s\n", key, value, parser.DetectType(value))
	}

  str := generator.GenerateStruct("Person", result)
  fmt.Println(str)
}