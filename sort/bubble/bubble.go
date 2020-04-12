package bubble

func Sort(arr []int)  {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}
	for i := 0; i < arrLen - 1; i++ {
		isSorted := true
		for j := 0; j < arrLen - i - 1; j++ {
			if arr[j] > arr[j + 1] {
				isSorted = false
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
			}
		}
		if isSorted {
			break
		}
	}
}