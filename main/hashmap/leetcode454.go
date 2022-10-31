package hashmap

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	twoSums := make(map[int]int)
	count := 0
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			twoSums[num1+num2]++
		}
	}
	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			if value, ok := twoSums[0-num3-num4]; ok {
				count = count + value
			}
		}
	}

	return count
}
