package main

import (
	"fmt"
	"regexp"
)

var caseRegex *regexp.Regexp

func init() {
	caseRegex = regexp.MustCompile(`[^\-\+]`)
}

func main() {
	cases := getCaseStrings()
	if cases == nil {
		return
	}

	for i, s := range cases {
		flipCount, lastUpDown := getGroupCount(s)
		if !lastUpDown {
			flipCount--
		}

		fmt.Printf("Case #%d: %d\n", i+1, flipCount)
	}
}

func validCaseString(caseString string) bool {
	return len(caseString) >= 1 &&
		len(caseString) <= 100 &&
		!caseRegex.MatchString(caseString)
}

func isUp(c rune) bool {
	return c == '+'
}

func getGroupCount(caseString string) (grpCount int, lastUpDown bool) {
	lastUpDown = !isUp(rune(caseString[len(caseString)-1]))
	grpCount = 1

	prior := isUp(rune(caseString[0]))
	for _, c := range caseString[1:] {
		up := isUp(c)

		if up == prior {
			continue
		}

		grpCount++
		prior = up
	}

	return
}
