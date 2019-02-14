package FloydRivest


func min (a, b int) int{
	if a < b {
		return a
	}
	return b
}


func max (a, b int) int {
	if a > b {
		return a
	}
	return b
}

type stack []int

func (s stack) push (v int) stack {
	return append(s, v)
}
func (s stack) pop () (stack, int) {
	l := len(s)
	return s[:l-1], s[l-1]
}


