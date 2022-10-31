package hashmap

func CanConstruct(ransomNote string, magazine string) bool {
	dict := make(map[int32]int)

	for _, value := range magazine {
		dict[value]++
	}

	for _, value := range ransomNote {
		if _, ok := dict[value]; ok {
			dict[value]--
			if dict[value] < 0 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
