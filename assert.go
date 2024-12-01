package adventofcode

import (
	"fmt"
	"os"
)

func Assert(assertion bool, message ...any) {
	if !assertion {
		fmt.Println(message...)
		os.Exit(1)
	}
}
