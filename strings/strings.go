package strings

import "regexp"

// RegexSplit is used to split string by regex
func RegexSplit(s string, pattern string, n int) []string {
	reg := regexp.MustCompile(pattern)
	return reg.Split(s, n)
}

