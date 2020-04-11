package merge

func Sort(arr []int)  {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}
	mergeSort(arr, 0, arrLen - 1)
}

func mergeSort(arr []int, start, end int)  {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(arr ,start, mid)
	mergeSort(arr, mid + 1, end)
	if arr[mid] > arr[mid + 1] { // 优化->判断是否需要合并
		merge(arr, start, mid, end)
	}
}

func merge(arr []int, start, mid, end int)  {
	temp := make([]int, end - start + 1)
	i, j, k := start, mid + 1, 0
	for i <= mid && j <= end {
		if arr[i] < arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}
	for ; i <= mid; i++ {
		temp[k] = arr[i]
		k++
	}
	for ; j <= end; j++ {
		temp[k] = arr[j]
		k++
	}
	for _, item := range temp {
		arr[start] = item
		start++
	}
}
