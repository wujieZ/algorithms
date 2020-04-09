package selection

func Sort(arr []int)  {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}
	for i := 0; i < arrLen; i++ {
		minIndex, minValue := i, arr[i]
		for j := i; j < arrLen - 1; j++ {
			if minValue > arr[j + 1] {
				minIndex = j + 1
				minValue = arr[j + 1]
			}
		}
		if i != minIndex {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}
