package nthElementUtils

func IntMin(a, b int) int{
	if a < b {
		return a
	}
	return b
}


func IntMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Stack []int

func (s Stack) Push(v int) Stack {
	return append(s, v)
}
func (s Stack) Pop() (Stack, int) {
	l := len(s)
	return s[:l-1], s[l-1]
}



