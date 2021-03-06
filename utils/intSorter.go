package nthElementUtils

type IntSorter []int


func (s IntSorter) Len() int {
	return len(s)
}

func (s IntSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s IntSorter) Less (i, j int) bool {
	return s[i] < s[j]
}
