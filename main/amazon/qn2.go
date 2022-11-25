package amazon

func CountFamilyLogins(logins []string) int32 {
	count := 0
	for i := 0; i < len(logins)-1; i++ {

		for j := i + 1; j < len(logins); j++ {
			if logins[i] == logins[j] {
				continue
			}
			isChecked := true
			isPlus := true
			if getDifference(logins[i][0], logins[j][0]) == 1 {
				isPlus = false
			} else if getDifference(logins[j][0], logins[i][0]) == 1 {
				isPlus = true
			} else {
				continue
			}
			for k := 0; k < len(logins[i]); k++ {
				if isPlus {
					if getDifference(logins[j][k], logins[i][k]) != 1 {
						isChecked = false
						break
					}
				} else {
					if getDifference(logins[i][k], logins[j][k]) != 1 {
						isChecked = false
						break
					}
				}
			}

			if isChecked {
				count++
			}
		}
	}

	return int32(count)
}

func getDifference(u uint8, u2 uint8) int {
	if u == 'z' && u2 == 'a' {
		return 1
	} else {
		return int(u - u2)
	}
}
