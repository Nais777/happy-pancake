package main

import "fmt"

func main() {
	cases := getCases()
	if cases == nil {
		return
	}

	for i, c := range cases {
		flipCount := c.grpCount

		if !c.lastUpDown {
			flipCount--
		}

		fmt.Printf("Case #%d: %d\n", i, flipCount)
	}
}
