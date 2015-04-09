package main

import (
	"log"

	"github.com/adampresley/sample-ini-parser/services/parser"
)

func main() {
	sampleInput := `
		[section]
		key1=value1
		key2=value2
		[bob]
		key3=value3
		key4=This is a big test of a longer, more awesome line

	`

	output := parser.Parse("sample.ini", sampleInput)
	log.Println(output)
}
