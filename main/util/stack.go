package util

type stack []interface{}

func NewStack() *stack {
	return &stack{}
}

func (s *stack) Push(v interface{}) {
	*s = append(*s, v)
}

func (s *stack) Pop() interface{} {
	h := *s
	var el interface{}
	l := len(h)
	if l == 0 {
		return nil
	}
	el, *s = h[l-1], h[0:l-1]
	return el
}
