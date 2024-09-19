package cmd

import (
	"fmt"
	"log"
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
	if err := validateArgs(); err != nil {
		log.Println(err)
		return
	}

	jsonData, err := readFile(os.Args[1])
	if err != nil {
		log.Printf("Error reading file: %v", err)
		return
	}

	result, err := parser.ParseJSON(jsonData)
	if err != nil {
		log.Printf("Error parsing JSON: %v", err)
		return
	}

	str, err := generateAndFormatStruct(os.Args[3], result)
	if err != nil {
		log.Printf("Error formatting struct: %v", err)
		return
	}

	if err := writeFile(os.Args[2], str); err != nil {
		log.Printf("Error writing file: %v", err)
	}
}

func validateArgs() error {
	if len(os.Args) != 4 {
		return fmt.Errorf("Usage: json2go <input.json> <output.go> <structName>")
	}
	return nil
}

func readFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

func writeFile(filePath, content string) error {
	if !checkFileExists(filePath) {
		if _, err := os.Create(filePath); err != nil {
			return fmt.Errorf("Error creating file: %w", err)
		}
	}
	return os.WriteFile(filePath, []byte(content), 0644)
}

func generateAndFormatStruct(structName string, data map[string]interface{}) (string, error) {
	str := generator.GenerateStruct(structName, data)
	return formatter.FormatStruct(str)
}