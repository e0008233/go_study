package string

func RepeatedSubstringPattern(s string) bool {
	ss := []byte(s)
	for end, _ := range ss {
		if (end+1)*2 > len(ss) {
			return false
		}
		if len(ss)%(end+1) == 0 {
			cycle := len(ss) / (end + 1)
			baseString := s[0 : end+1]
			isFound := true
			for j := 1; j < cycle; j++ {
				toCompare := s[len(baseString)*j : len(baseString)*(j+1)]
				if baseString != toCompare {
					isFound = false
					break
				}

			}
			if isFound {
				return true
			}
		}

	}
	return false
}
