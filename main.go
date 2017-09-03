package main

import "fmt"

func main() {
	cases := getCases()
	if cases == nil {
		return
	}

	for i, c := range cases {
		flips := getFlipCount(c)

		fmt.Printf("Case #%d: %d\n", i, flips)
	}
}
