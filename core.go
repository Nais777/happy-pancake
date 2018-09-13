package pancake

import (
	"regexp"
)

var caseRegex *regexp.Regexp

func init() {
	caseRegex = regexp.MustCompile(`[^\-\+]`)
}

func ValidCaseString(caseString string) bool {
	return len(caseString) >= 1 &&
		len(caseString) <= 100 &&
		!caseRegex.MatchString(caseString)
}

func GetGroupCount(caseString string) (grpCount int, lastUpsideDown bool) {
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
