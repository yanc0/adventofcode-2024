package main

import (
	"embed"
	"fmt"
	"log"
	"os"
)

//go:embed input.txt
var input embed.FS

func main() {
	f, err := input.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	parser := NewParser(NewScanner(f))

	total := int64(0)
	for {
		num, err := parser.Parse()
		if err != nil {
			if err.Error() == "eof" {
				fmt.Println("total:", total)
				os.Exit(0)
			}
			continue
		}
		total += int64(num)
	}
}
