package hashmap

import "math"

func commonChars(words []string) []string {

	var result []string
	hashArrays := make([][26]int, 0)

	for _, word := range words {
		rowResult := [26]int{}
		for _, char := range word {
			rowResult[char-97] = rowResult[char-97] + 1
		}
		hashArrays = append(hashArrays, rowResult)
	}

	for i := 0; i < 26; i++ {
		counts := math.MaxInt
		for j := 0; j < len(hashArrays); j++ {
			if hashArrays[j][i] < counts {
				counts = hashArrays[j][i]
			}
		}
		for k := 0; k < counts; k++ {
			result = append(result, string(rune(97+i)))
		}
	}

	return result
}
