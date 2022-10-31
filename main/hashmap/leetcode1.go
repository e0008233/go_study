package hashmap

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for index, value := range nums {
		if num, ok := numMap[target-value]; ok {
			return []int{index, num}
		} else {
			numMap[value] = index
		}
	}

	return []int{0, 0}
}
