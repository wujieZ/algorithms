package insertion

/**
 params: inc增量, 方便希尔排序调用
 */
func Sort(arr []int, inc int) {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}
	for i := inc; i < arrLen; i++ {
		currentValue := arr[i]
		j := i

		for ; j >= inc; j = j - inc {
			if currentValue < arr[j - inc] {
				arr[j] = arr[j - inc]
			} else {
				break
			}
		}
		arr[j] = currentValue
	}
}
