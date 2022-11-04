package hashmap

import "sort"

func FourSum(nums []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for j := 0; j < len(nums)-3; j++ {
		if j != 0 && nums[j] == nums[j-1] {
			continue
		}
		for i := j + 1; i < len(nums)-2; i++ {
			if i != j+1 && nums[i] == nums[i-1] {
				continue
			}
			twoNumSum := target - nums[i] - nums[j]
			left := i + 1
			right := len(nums) - 1
			for left < right {
				if nums[left]+nums[right] == twoNumSum {
					toAdd := []int{nums[i], nums[left], nums[right], nums[j]}
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
	}
	return result
}
