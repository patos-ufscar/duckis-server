package common

import (
	"regexp"
)

var (
	wildCardPattern string = `\*`
	matchWildCardPattern *regexp.Regexp = regexp.MustCompile(wildCardPattern)
)

func ConvertWildCardsToPattern(s string) string {
	if !matchWildCardPattern.MatchString(s) {
		return s
	}

	return matchWildCardPattern.ReplaceAllString(s, wildCardPattern)
}
