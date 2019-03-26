package common


/* a function to compare []string */
func StringSliceEqual(s1 []string, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	} else if (s1 == nil) != (s2 == nil) {
		return false
	}

	for idx, _ := range s1 {
		if s1[idx] != s2[idx] {
			return false
		}
	}

	return true
}


/* a function to compare [][]string */
func StringSSEqual(res [][]string, expect [][]string) bool {
	if len(res) != len(expect) {
		return false
	}
	for idx, _ := range expect {
		if StringSliceEqual(res[idx], expect[idx]) == false {
			return false
		}
	}
	return true
}