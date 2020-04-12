package queue

type Queue []int

func (q *Queue) Push(items... int) int{
	*q = append(*q, items...)
	return len(*q)
}

func (q *Queue) Shift() (int, bool){
	queueLen := len(*q)
	if queueLen <= 0 {
		return 0, false
	}
	shiftItem := (*q)[0]
	*q = (*q)[1:queueLen]
	return shiftItem, true
}

func (q *Queue) getLen() int{
	return len(*q)
}

