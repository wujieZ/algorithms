package insert

import "testing"

func TestInsertSort(t *testing.T)  {
	tests := []struct{
		arr []int
		result []int
	} {
		{
			[]int{2, 1, 3, 4},
			[]int{1, 2, 3, 4},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{},
			[]int{},
		},
		{
			[]int{23, 45, 131, 67, 76, 10, 23, 5, 7, 89},
			[]int{5, 7, 10, 23, 23, 45, 67, 76, 89, 131},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 10},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 10},
		},
		{
			[]int{10, 8, 7, 6, 5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 10},
		},
	}

	for _, data := range tests {
		Sort(data.arr)
		for i, item := range data.arr {
			if item != data.result[i] {
				t.Errorf("expectd(%v), actual(%v)", data.result, data.arr)
				break
			}
		}
	}
}
