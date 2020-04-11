package shell

import "golang-algorithm/sort/insertion"

func Sort(arr []int)  {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}
	inc := 1
	for inc < arrLen / 3 { //寻找合适的增量
		inc = 3 * inc + 1
	}
	for ; inc > 0; inc = inc / 3 {
		insertion.Sort(arr, inc)
	}
}
