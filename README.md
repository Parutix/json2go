# json2go

json2go is a command-line tool that simplifies the process of converting JSON data into Go structs. It automatically generates Go struct definitions based on your JSON input, allowing you to seamlessly integrate and work with JSON data in your Go applications.

## Motivation

In the development of Go applications, working with JSON data often involves creating and maintaining complex struct definitions manually. This process can be time-consuming and error-prone, especially when dealing with large or evolving JSON schemas. json2go was created to address this challenge by automating the conversion of JSON data into Go structs.

### Goal

The primary goal of json2go is to streamline the process of converting JSON data into Go struct definitions. By automating this conversion, json2go aims to eliminate manual, error-prone tasks and accelerate development. 

The application seeks to enhance productivity and code quality for Go developers by providing a reliable tool that generates accurate, well-formatted Go structs from JSON input, thus facilitating easier integration and manipulation of JSON data within Go applications.

## ⚙️ Installation

Open a terminal and write:

go install github.com/Parutix/json2go

## Usage

json2go <input.json> <output.go> <structName>

