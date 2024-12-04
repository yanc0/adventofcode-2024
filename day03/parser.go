package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Parser struct {
	s         *Scanner
	activated bool
}

func NewParser(s *Scanner) *Parser {
	return &Parser{
		s:         s,
		activated: true,
	}
}

func (p *Parser) parseMul() (mul int, err error) {
	var a, b int
	if t, _ := p.s.NextToken(); t != OPENING_PARENTHESIS {
		return 0, fmt.Errorf("expected opening parenthesis")
	}

	t, lit := p.s.NextToken()
	if t != IDENT {
		return 0, fmt.Errorf("expected ident number")
	}
	a, err = strconv.Atoi(lit)
	if err != nil {
		return 0, err
	}
	if t, _ := p.s.NextToken(); t != COMMA {
		return 0, fmt.Errorf("parse error, expected comma")
	}
	t, lit = p.s.NextToken()
	if t != IDENT {
		return 0, fmt.Errorf("expected ident number")
	}
	b, err = strconv.Atoi(lit)
	if err != nil {
		return 0, err
	}
	if t, _ := p.s.NextToken(); t != CLOSING_PARENTHESIS {
		return 0, fmt.Errorf("expected closing parenthesis")
	}

	return a * b, nil
}

func (p *Parser) parseDo() {
	if t, _ := p.s.NextToken(); t != OPENING_PARENTHESIS {
		return
	}
	if t, _ := p.s.NextToken(); t != CLOSING_PARENTHESIS {
		return
	}
	p.activated = true
}

func (p *Parser) parseDont() {
	if t, _ := p.s.NextToken(); t != OPENING_PARENTHESIS {
		return
	}
	if t, _ := p.s.NextToken(); t != CLOSING_PARENTHESIS {
		return
	}
	p.activated = false
}

func (p *Parser) Parse() (mul int, err error) {
	for {
		t, lit := p.s.NextToken()
		if t == EOF {
			return 0, fmt.Errorf("eof")
		}
		if t == IDENT && strings.HasSuffix(lit, "MUL") {
			if !p.activated {
				break
			}
			return p.parseMul()
		}
		if t == IDENT && strings.HasSuffix(lit, "DO") {
			p.parseDo()
			break
		}
		if t == IDENT && strings.HasSuffix(lit, "DON'T") {
			p.parseDont()
			break
		}
	}
	return 0, fmt.Errorf("no action")
}
