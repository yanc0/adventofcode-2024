package main

import (
	. "aoc"
	"strings"
	"testing"
)

var testinput = `47|53
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
97,13,75,29,47`

func TestParse(t *testing.T) {
	rules, updates := parseInput(strings.NewReader(testinput))

	Assert(len(rules) == 6 && len(rules[47]) == 4)
	Assert(len(updates) == 6)
	Assert(updates.AtIndex(0).MiddlePageNumber() == 61)
	Assert(updates.AtIndex(4).MiddlePageNumber() == 13)
}

func TestValidate(t *testing.T) {
	rules, updates := parseInput(strings.NewReader(testinput))
	Assert(updates[0].IsValidAgainstRules(rules))
	Assert(updates[1].IsValidAgainstRules(rules))
	Assert(updates[2].IsValidAgainstRules(rules))
	Assert(!updates[3].IsValidAgainstRules(rules))
	Assert(!updates[4].IsValidAgainstRules(rules))
	Assert(!updates[5].IsValidAgainstRules(rules))
}

func TestFixUpdates(t *testing.T) {
	rules, updates := parseInput(strings.NewReader(testinput))
	Assert(updates[3].AlignWithRules(rules).IsValidAgainstRules(rules))
	Assert(updates[4].AlignWithRules(rules).IsValidAgainstRules(rules))
	Assert(updates[5].AlignWithRules(rules).IsValidAgainstRules(rules))
}
