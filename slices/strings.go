package slices

func StringIn(ss []string, s string) bool {
	for _, find := range ss {
		if find == s {
			return true
		}
	}
	return false
}

func StringMap(ss []string, cb func(s string) interface{}) []interface{} {
	ret := make([]interface{}, 0, len(ss))
	for _, s := range ss {
		ret = append(ret, cb(s))
	}
	return ret
}


