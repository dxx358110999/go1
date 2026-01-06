package my_util

import (
	"regexp"
)

var compiledPatterns []*regexp.Regexp

func init() {
	var patterns = []string{
		`.*\d+.*`,
		`.*[a-z]+.*`,
		`.*[A-Z]+.*`,
		`.*\W+.*`,
	}
	for _, p := range patterns {
		compiledPatterns = append(compiledPatterns, regexp.MustCompile(p))
	}
}

func IsValidPassword(password string, targetScore int) bool {

	score := 0

	for _, cp := range compiledPatterns {
		ok := cp.MatchString(password)
		if ok {
			score += 1
		}
	}

	return score >= targetScore
}
