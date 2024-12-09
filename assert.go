package adventofcode

import (
	"fmt"
	"os"
	"strconv"
)

func Assert(assertion bool, message ...any) {
	if !assertion {
		fmt.Println(message...)
		os.Exit(1)
	}
}

func Int(str string) int {
	num, err := strconv.Atoi(str)
	Assert(err == nil, str, "is not an int:", err)
	return num
}

func Ints(strs []string) []int {
	ints := make([]int, len(strs))
	for i, str := range strs {
		ints[i] = Int(str)
	}
	return ints
}
