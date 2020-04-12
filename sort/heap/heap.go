package heap

func adjustHeap(arr []int, i, arrLen int)  {
	temp := arr[i]
	k := i * 2 + 1
	for ; k < arrLen; k = k * 2 + 1 {
		if k + 1 < arrLen && arr[k] < arr[k + 1] {
			k++
		}
		if arr[k] > temp {
			arr[i] = arr[k]
			i = k
		} else {
			break
		}
	}
	arr[i] = temp
}

func Sort(arr []int)  {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}
	for i := arrLen / 2 - 1; i >= 0; i-- {
		adjustHeap(arr, i, arrLen)
	}
	for j := arrLen - 1; j > 0; j-- {
		arr[0], arr[j] = arr[j], arr[0]
		adjustHeap(arr, 0, j)
	}
}