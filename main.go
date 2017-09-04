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

func getGroupCount(caseString string) (grpCount int, lastUpsideDown bool) {
	lastUpsideDown = caseString[len(caseString)-1] == '-'
	grpCount = 1

	prior := rune(caseString[0])
	for _, c := range caseString[1:] {
		if c == prior {
			continue
		}

		grpCount++
		prior = c
	}

	return
}
