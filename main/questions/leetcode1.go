package questions

func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for index, value := range nums {
		if value, ok := numMap[target-value]; ok {
			return []int{index, numMap[target-value]}
		} else {
			numMap[value] = index
		}
	}

	return []int{0, 0}
}
