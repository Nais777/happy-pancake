package main

import (
	"fmt"
	"regexp"
)

var caseRegex *regexp.Regexp

func init() {
	caseRegex = regexp.MustCompile(`[^\-\+]`)
}

type caseDef struct {
	grpCount   int
	lastUpDown bool
}

func getCases() []*caseDef {
	var caseCount int
	if _, err := fmt.Scanf("%d", &caseCount); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	cases := make([]*caseDef, caseCount, caseCount)

	var line string
	for i := 0; i < caseCount; i++ {
		if _, err := fmt.Scanln(&line); err != nil {
			fmt.Println(err)
			return nil
		} else if !validCaseString(line) {
			fmt.Println("Invalid case string. Must contain only - and + characters, be >= 1 character, and be <= 100 characters")
			return nil
		}

		c, upDown := getGroupCount(line)
		cases[i] = &caseDef{grpCount: c, lastUpDown: upDown}
	}

	return cases
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
