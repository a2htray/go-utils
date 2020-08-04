package slices

// StringIn is to check whether a string element is in a slice
func StringIn(ss []string, s string) bool {
	for _, find := range ss {
		if find == s {
			return true
		}
	}
	return false
}

// StringMap ...
func StringMap(ss []string, cb func(s string) string) string {
	ret := make([]string, 0, len(ss))
	for _, s := range ss {
		ret = append(ret, cb(s))
	}
	return ret
}

// StringForeach ...
func StringForeach(ss []string, cb func(s string)) {
	for _, s := range ss {
		cb(s)
	}
}

// StringUnique
func StringUnique(ss []string) []string {
	m := map[string]bool{}
	for _, s := range ss {
		m[s] = true
	}
	ret := make([]string, 0, len(m))
	for k := range m {
		ret = append(ret, k)
	}
	return ret
}

// StringExtend
func StringExtend(ss []string, extends... []string) []string {
	ret := ss
	for _, extend := range extends {
		ret = append(ret, extend...)
	}
	return ret
}

