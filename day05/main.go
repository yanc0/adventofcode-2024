package main

import (
	. "aoc"
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

func main() {
	rules, updates := parseInput(InputFile())

	total := 0
	totalFixed := 0
	for _, update := range updates {
		if update.IsValidAgainstRules(rules) {
			total += update.MiddlePageNumber()
			continue
		}
		totalFixed += update.AlignWithRules(rules).MiddlePageNumber()
	}
	fmt.Println("middle page sum of correctly ordered updates:", total)
	fmt.Println("middle page sum of fixed updates:", totalFixed)
}

func parseInput(input io.Reader) (Rules, Updates) {
	rules := make(Rules)
	updates := make(Updates, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// rules ends, jumps to updates
			break
		}
		splited := strings.Split(line, "|")
		Assert(len(splited) == 2, "rule must be separated by |")

		num, err := strconv.Atoi(splited[0])
		Assert(err == nil, "first num should be integer", err)
		before, err := strconv.Atoi(splited[1])
		Assert(err == nil, "second num should be integer", err)

		rules[num] = append(rules[num], before)
	}

	for scanner.Scan() {
		line := scanner.Text()
		splited := strings.Split(line, ",")
		updates = append(updates, ConvertToInts(splited))
	}
	return rules, updates
}

func ConvertToInts(strs []string) []int {
	ints := make([]int, len(strs))
	for i, str := range strs {
		num, err := strconv.Atoi(str)
		Assert(err == nil, "str should be integer convertible", err)
		ints[i] = num
	}
	return ints
}

type Rules map[int][]int

func (r Rules) IsPageAllowedAfter(page int, previouses []int) bool {
	rules, exists := r[page]
	if !exists { // no rules for page
		return true
	}
	for _, rule := range rules {
		for _, previous := range previouses {
			if rule == previous {
				return false
			}
		}
	}
	return true
}

type Updates []Update
type Update []int

func (updates Updates) AtIndex(index int) Update {
	return updates[index]
}

func (u Update) MiddlePageNumber() int {
	isOdd := len(u)%2 == 1
	containsAtLeastOnePage := len(u) > 0
	Assert(isOdd && containsAtLeastOnePage, "update must contain an odd number of pages")
	return u[len(u)/2]
}

func (u Update) IsValidAgainstRules(rules Rules) bool {
	for i, page := range u {
		if !rules.IsPageAllowedAfter(page, u[0:i]) {
			return false
		}
	}
	return true
}

func (u Update) AlignWithRules(rules Rules) Update {
	i := 0
	for {
		if !rules.IsPageAllowedAfter(u[i], u[0:i]) {
			// move page to first position
			u = slices.Insert(u, 0, u[i])
			u = slices.Delete(u, i+1, i+2)
			i = 0
		}
		if u.IsValidAgainstRules(rules) {
			break
		}
		i++
	}
	return u
}
