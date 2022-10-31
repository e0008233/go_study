package hashmap

import "sort"

func ThreeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		twoNumSum := 0 - nums[i]
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[left]+nums[right] == twoNumSum {
				toAdd := []int{nums[i], nums[left], nums[right]}
				result = append(result, toAdd)
				tempLeft := left
				for left < right && nums[left] == nums[tempLeft] {
					left++
				}
			} else {
				if nums[left]+nums[right] > twoNumSum {
					right--
				} else {
					left++
				}
			}

		}
	}
	return result
}
