package main

import (
	adventofcode "aoc"
	"strings"
	"testing"
)

var testinput = strings.NewReader(`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`)

func TestParse(t *testing.T) {
	rules, updates := parseInput(testinput)

	adventofcode.Assert(len(rules) == 6 && len(rules[47]) == 4)
	adventofcode.Assert(len(updates) == 6)
	adventofcode.Assert(updates.AtIndex(0).MiddlePageNumber() == 61)
	adventofcode.Assert(updates.AtIndex(4).MiddlePageNumber() == 13)
}

func TestValidate(t *testing.T) {
	rules, updates := parseInput(testinput)
	adventofcode.Assert(updates[0].IsValidAgainstRules(rules))
	adventofcode.Assert(updates[1].IsValidAgainstRules(rules))
	adventofcode.Assert(updates[2].IsValidAgainstRules(rules))
	adventofcode.Assert(!updates[3].IsValidAgainstRules(rules))
	adventofcode.Assert(!updates[4].IsValidAgainstRules(rules))
	adventofcode.Assert(!updates[5].IsValidAgainstRules(rules))
}

func TestFixUpdates(t *testing.T) {
	rules, updates := parseInput(testinput)
	adventofcode.Assert(updates[3].AlignWithRules(rules).IsValidAgainstRules(rules))
	adventofcode.Assert(updates[4].AlignWithRules(rules).IsValidAgainstRules(rules))
	adventofcode.Assert(updates[5].AlignWithRules(rules).IsValidAgainstRules(rules))
}
