package insert

func Sort(arr []int) {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}
	for i := 1; i < arrLen; i++ {
		currentValue := arr[i]
		insertIndex := i

		for j := i - 1; j >= 0; j-- {
			if currentValue < arr[j] {
				arr[j + 1] = arr[j]
				insertIndex = j
			} else {
				break
			}
		}
		arr[insertIndex] = currentValue
	}
}
