package arrays

func Search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
		if left >= right {
			if left == right && nums[left] == target {
				return left
			}
			break
		}
	}
	return -1
}
