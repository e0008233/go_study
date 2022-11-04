package string

func ReverseStr(s string, k int) string {
	ss := []byte(s)
	length := len(ss)
	needReverse := true
	for i := 0; i < length; i = i + k {
		if needReverse {
			if i+k-1 <= length-1 {
				reverseStringByPosition(ss, i, i+k-1)
			} else {
				reverseStringByPosition(ss, i, length-1)
			}
		}
		needReverse = !needReverse
	}
	return string(ss)
}

func reverseStringByPosition(s []byte, start, end int) {
	for i := 0; i < (end-start+1)/2; i++ {
		temp := s[start+i]
		s[start+i] = s[end-i]
		s[end-i] = temp
	}
}
