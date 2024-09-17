package cmd

import (
	"fmt"
	"os"

	"github.com/Parutix/json2go/formatter"
	"github.com/Parutix/json2go/generator"
	"github.com/Parutix/json2go/parser"
)

// checkFileExists checks if a file exists
func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

// Run runs the CLI
func Run() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: json2go <input.json> <output.go> <structName>")
		return
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

	// Check if the output file exists, if not create it
	if !checkFileExists(os.Args[2]) {
		_, err := os.Create(os.Args[2])
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
	}
	outputFile := os.Args[2]
	
	// Write the formatted struct to the output file
	err = os.WriteFile(outputFile, []byte(str), 0644)
}