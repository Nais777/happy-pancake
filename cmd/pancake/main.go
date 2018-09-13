package main

import (
	"fmt"

	pancake "github.com/Nais777/happy-pancake"
)

func main() {
	cases := pancake.GetCaseStrings()
	if cases == nil {
		return
	}

	for i, s := range cases {
		flipCount, lastUpDown := pancake.GetGroupCount(s)
		if !lastUpDown {
			flipCount--
		}

		fmt.Printf("Case #%d: %d\n", i+1, flipCount)
	}
}
