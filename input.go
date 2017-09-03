package main

import (
	"fmt"
	"regexp"
)

var caseRegex *regexp.Regexp

func init() {
	caseRegex = regexp.MustCompile(`[^\-\+]`)
}

func getCaseStrings() []string {
	var caseCount int
	if _, err := fmt.Scanf("%d", &caseCount); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	if caseCount < 1 || caseCount > 100 {
		fmt.Println("Number of test cases must be between 1 and 100 inclusive.")
		return nil
	}

	cases := make([]string, caseCount, caseCount)

	var line string
	for i := 0; i < caseCount; i++ {
		if _, err := fmt.Scanln(&line); err != nil {
			fmt.Println(err)
			return nil
		} else if !validCaseString(line) {
			fmt.Println("Invalid case string. Must contain only - and + characters, be >= 1 character, and be <= 100 characters")
			return nil
		}

		cases[i] = line
	}

	return cases
}
