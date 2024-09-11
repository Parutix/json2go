package cmd

import (
	"fmt"
	"os"

	"github.com/Parutix/json2go/formatter"
	"github.com/Parutix/json2go/generator"
	"github.com/Parutix/json2go/parser"
)

func Run() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: json2go <input.json> <output.go> <structName>")
	}

	jsonFile := os.Args[1]
	structName := os.Args[3]

	jsonData, err := os.ReadFile(jsonFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	result, err := parser.ParseJSON(jsonData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	str := generator.GenerateStruct(structName, result)
	str, err = formatter.FormatStruct(str)
	if err != nil {
		fmt.Println("Error formatting struct:", err)
		return
	}

	outputFile := os.Args[2]
	err = os.WriteFile(outputFile, []byte(str), 0644)
}