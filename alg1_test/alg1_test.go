package alg1_test

import (
	"fmt"
	"testing"
)

func TestAlg1(t *testing.T) {
	// даны 2 отсортированных массива по возрастанию (слайсы)
	// вернуть третий слайс с мержем 2х списков. Порядок по возрастанию

	// 1 слайс len = a
	// 2 слайс len = b
	// a+b=n
	// временная сложность: O(n);
	// пространственная сложность: a+b = n; 2n * 3; O(n)

	a := []int{1, 1, 2, 5, 7, 8, 11, 27}
	b := []int{1, 3, 4, 5, 7, 8, 11, 27}

	fmt.Println(mergeSlice(a, b))

}

func mergeSlice(a, b []int) []int {
	resLen := len(a) + len(b)
	res := make([]int, resLen)
	iteratorA := 0
	iteratorB := 0
	for i, _ := range res {
		if iteratorA >= len(a) {
			//res[i] = b[iteratorB]
			//iteratorB++
			//
			copy(res[i:], b[iteratorB:])
			break
		}
		if iteratorB >= len(b) {
			res[i] = a[iteratorA]
			iteratorA++
			continue
		}

		if a[iteratorA] > b[iteratorB] {
			res[i] = b[iteratorB]
			iteratorB++
		} else {
			res[i] = a[iteratorA]
			iteratorA++
		}

	}
	return res
}
