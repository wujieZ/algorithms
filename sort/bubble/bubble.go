package bubble

func Sort(arr []int)  {
	len := len(arr)
	if len <= 1 {
		return
	}
	for i := 0; i < len - 1; i++ {
		isSort := false
		for j := 0; j < len - i - 1; j++ {
			if arr[j] > arr[j + 1] {
				isSort = true
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
			}
		}
		if !isSort {
			break
		}
	}
}