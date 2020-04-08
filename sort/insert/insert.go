package insert

func Sort(arr []int) {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}
	for i := 1; i < arrLen; i++ {
		currentValue := arr[i]
		j := i

		for ; j > 0; j-- {
			if currentValue < arr[j - 1] {
				arr[j] = arr[j - 1]
			} else {
				break
			}
		}
		arr[j] = currentValue
	}
}
