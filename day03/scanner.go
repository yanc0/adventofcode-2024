package main

import (
	"bufio"
	"io"
	"strings"
)

type Token int

const (
	IDENT Token = iota
	COMMA
	OPENING_PARENTHESIS
	CLOSING_PARENTHESIS
	UNKNOWN
	EOF
)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		r: bufio.NewReader(r),
	}
}

func isLetter(rn rune) bool {
	return rn >= 'a' && rn <= 'z' || rn >= 'A' && rn <= 'Z' || rn == '\''
}

func isNumber(rn rune) bool {
	return rn >= '0' && rn <= '9'
}

func isSpecial(rn rune) bool {
	specials := []rune{'(', ')', ','}
	for _, special := range specials {
		if rn == special {
			return true
		}
	}
	return false
}

func (s *Scanner) ScanIdent() string {
	ident := ""
	for {
		rn, _, err := s.r.ReadRune()
		if err != nil {
			return strings.ToUpper(ident)
		}
		if !isLetter(rn) && !isNumber(rn) {
			s.unread()
			return strings.ToUpper(ident)
		}

		ident += string(rn)
	}
}

func (s *Scanner) unread() {
	_ = s.r.UnreadRune()

}

func (s *Scanner) NextToken() (Token, string) {
	tkn := ""

	rn, _, err := s.r.ReadRune()
	if err != nil {
		return EOF, ""
	}
	if isLetter(rn) || isNumber(rn) {
		s.unread()
		tkn = s.ScanIdent()
	}
	if isSpecial(rn) {
		tkn = string(rn)
	}

	switch tkn {
	case "":
		return UNKNOWN, ""
	case "(":
		return OPENING_PARENTHESIS, ""
	case ")":
		return CLOSING_PARENTHESIS, ""
	case ",":
		return COMMA, ""
	default:
		return IDENT, tkn
	}
}
