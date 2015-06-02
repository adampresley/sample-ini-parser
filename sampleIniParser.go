package main

import (
	"encoding/json"
	"log"

	"github.com/adampresley/sample-ini-parser/services/parser"
)

func main() {
	sampleInput := `
		key=abcdefg

		[User]
		userName=adampresley
		keyFile=~/path/to/keyfile

		[Servers]
		server1=localhost:8080
	`

	parsedINIFile := parser.Parse("sample.ini", sampleInput)
	prettyJSON, err := json.MarshalIndent(parsedINIFile, "", "   ")

	if err != nil {
		log.Println("Error marshalling JSON:", err.Error())
		return
	}

	log.Println(string(prettyJSON))
}
