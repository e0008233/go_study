package stack_and_queue

func RemoveDuplicates(s string) string {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			s = s[:i] + s[i+2:]
			i = i - 2
			if i < 0 {
				i = -1
			}
		}
	}
	return s
}
