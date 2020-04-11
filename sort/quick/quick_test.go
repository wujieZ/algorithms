package quick

import (
	"math/rand"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T)  {
	tests := []struct{
		original []int
		result []int
	} {
		{
			[]int{2, 1},
			[]int{1, 2},
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
			[]int{1, 2, 3, 5, 4},
			[]int{1, 2, 3, 4, 5},
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
		{
			[]int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, test := range tests {
		Sort(test.original)
		for i, item := range test.original {
			if item != test.result[i] {
				t.Errorf("expectd(%v), actual(%v)", test.result, test.original)
				break
			}
		}
	}
}

func BenchmarkQuickSort(b *testing.B) {
	// 数据准备
	var s []int
	for i := 0; i < 10000000; i++ {
		s = append(s, rand.Intn(10000))
	}
	result := make([]int, 10000000)
	copy(result, s)
	sort.Ints(result)
	// 重置测试时间
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Sort(s)
		for index, item := range s{
			if item != result[index] {
				b.Errorf("expectd(%v), actual(%v)", result, s)
				break
			}
		}
	}
}