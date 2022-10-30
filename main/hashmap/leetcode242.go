package hashmap

func isAnagram(s string, t string) bool {
	dictionary := make(map[string]int)
	for _, char := range s {
		if _, ok := dictionary[string(char)]; ok {
			dictionary[string(char)] = dictionary[string(char)] + 1
		} else {
			dictionary[string(char)] = 1
		}
	}

	for _, char := range t {
		if _, ok := dictionary[string(char)]; ok {
			dictionary[string(char)] = dictionary[string(char)] - 1
		} else {
			return false
		}
	}

	for _, value := range dictionary {
		if value != 0 {
			return false
		}
	}
	return true
}
