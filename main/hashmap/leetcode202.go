package hashmap

func isHappy(n int) bool {
	cache := make(map[int]bool)

	return isHappyHelper(n, cache)
}

func isHappyHelper(n int, cache map[int]bool) bool {
	temp := n
	if n == 1 {
		return true
	}
	if n == 0 {
		return false
	}
	if _, ok := cache[n]; ok {
		return false
	}
	sum := 0
	for n > 0 {
		m := n % 10
		sum = sum + m*m
		n = n / 10
	}
	if sum == 1 {
		return true
	} else {
		cache[temp] = true
		return isHappyHelper(sum, cache)
	}
}
