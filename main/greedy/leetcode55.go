package greedy

func CanJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	maxStep := 0
	for i := 0; i <= maxStep; i++ {
		currentStep := i + nums[i]
		if currentStep > maxStep {
			maxStep = currentStep
		}
		if maxStep >= len(nums)-1 {
			return true
		}
	}

	return false
}
