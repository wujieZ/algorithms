package stack

type Stack []int

func (s *Stack) Push(item...int) int{
	*s = append(*s, item...)
	return len(*s)
}

func (s *Stack) Pop() (int, bool){
	sLen := len(*s)
	if sLen <= 0 {
		return 0, false
	}
	popItem := (*s)[sLen - 1]
	*s = (*s)[0:sLen - 1]
	return popItem, true
}

func (s *Stack) isEmpty() bool{
	return len(*s) == 0
}