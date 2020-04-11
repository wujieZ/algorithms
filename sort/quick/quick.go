package quick

func Sort(arr []int)  {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}
	quickSort(arr, 0, arrLen - 1)
}


func quickSort(arr []int, start, end int)  {
	if start >= end {
		return
	}
	midIndex := findMidIndex(arr, start, end)
	quickSort(arr, start, midIndex - 1)
	quickSort(arr, midIndex + 1, end)
}

func findMidIndex(arr []int, start, end int) int {
	pivot := arr[end]
	i, j := start, end - 1
	for i < j {
		for arr[i] > pivot && i < j {
			arr[i], arr[j] = arr[j], arr[i]
			j--
		}
		if i < j {
			i++
		} else {
				break
		}
	}
	var midIndex int
	if arr[i] > pivot {
		midIndex = i
	} else {
		midIndex = i + 1
	}
	arr[midIndex], arr[end] = arr[end], arr[midIndex]
	return midIndex
}