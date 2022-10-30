package hashmap

func intersection(nums1 []int, nums2 []int) []int {
	hashmap := make(map[int]int)
	var result []int
	resultMap := make(map[int]int)
	for _, value := range nums1 {
		hashmap[value] = 1
	}

	for _, value := range nums2 {
		if _, ok := hashmap[value]; ok {
			resultMap[value] = 1
		}
	}
	for key, _ := range resultMap {
		result = append(result, key)
	}
	return result
}
