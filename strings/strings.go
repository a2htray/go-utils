package strings

import "regexp"

// RegexSplit is used to split string by regex
// https://stackoverflow.com/questions/4466091/split-string-using-regular-expression-in-go
func RegexSplit(s string, pattern string, n int) []string {
	reg := regexp.MustCompile(pattern)
	return reg.Split(s, n)
}

