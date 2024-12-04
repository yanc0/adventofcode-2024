package main

import (
	adventofcode "aoc"
	"strings"
	"testing"
)

func TestScanner(t *testing.T) {
	code := strings.NewReader("from()why()?mul(603,692)({select()}] )]-(mul(387,685)")

	scanner := NewScanner(code)
	for {
		t, _ := scanner.NextToken()
		if t == EOF {
			return
		}
	}
}

func TestParse(t *testing.T) {
	code := strings.NewReader(`from()why()?mul(603,692)({select()}] )]
	-(mul(387,685)`)

	parser := NewParser(NewScanner(code))
	n, err := parser.Parse()
	if err != nil {
		t.Fatal(err)
	}
	adventofcode.Assert(n == 417276, "n != 417276")

	n, err = parser.Parse()
	if err != nil {
		t.Fatal(err)
	}
	adventofcode.Assert(n == 265095, "n != 265095")

	_, err = parser.Parse()
	if err != nil {
		adventofcode.Assert(err.Error() == "eof", "expecting eof")
	}
}
