package alg2_test

import (
	"fmt"
	"sort"
	"testing"
)

func TestAlg2(t *testing.T) {
	// []int - найти k наиболее частых значений. порядок не имеет зхначения
	// неупорядочен. Могут отрицательные

	slice := []int{1, 2, 1, 1, 3, 7, -5, 1, -7, 22, 22, 22, 6, 3}

	fmt.Println(mostFriquent(slice, 40))

}

func mostFriquent(inpSlice []int, k int) []int {
	// len of inpSlice = a
	keyStorage := make(map[int]int)
	//currentValue := 0
	for _, value := range inpSlice { // o(a) // mem O(b) uniq
		currentValue, _ := keyStorage[value]
		keyStorage[value] = currentValue + 1
	}

	lenOfRes := 0
	if len(keyStorage) > k {
		lenOfRes = k
	} else {
		lenOfRes = len(keyStorage)
	}
	res := make([]int, lenOfRes)

	type keyValue struct {
		key   int
		value int
	}
	sliceOfValues := make([]keyValue, len(keyStorage)) // mem O(b)
	iteratorSlice := 0
	// len of unic B; B<=a
	for key, value := range keyStorage { // O(b)
		sliceOfValues[iteratorSlice] = keyValue{key, value}
		iteratorSlice++
	}
	//quick sort O(nlogn) // mem const
	sort.Slice(sliceOfValues, func(i, j int) bool { // O(bLogb)
		return sliceOfValues[i].value > sliceOfValues[j].value
	})

	for i, value := range sliceOfValues { // O(k), // mem O(min(k, b))
		res[i] = value.key
		if i == k-1 {
			break
		}
	}

	// O(a + b + blogb + k) // худший/лучший/средний
	// mem// O(2b + min(b,K))

	return res

}

func mostFriquent2(inpSlice []int, k int) []int {
	// len of inpSlice = a
	keyStorage := make(map[int]int)
	//currentValue := 0
	for _, value := range inpSlice { // o(a) // mem O(b) uniq
		currentValue, _ := keyStorage[value]
		keyStorage[value] = currentValue + 1
	}

	lenOfRes := 0
	if len(keyStorage) > k {
		lenOfRes = k
	} else {
		lenOfRes = len(keyStorage)
	}
	res := make([]int, 0, lenOfRes) // mem min(k, b)

	type keyValue struct {
		key   int
		value int
	}
	sliceOfValues := make([][]int, len(inpSlice)) // mem O(a)
	//iteratorSlice := 0
	// len of unic B; B<=a
	for key, value := range keyStorage { // O(b)
		sliceOfValues[value] = append(sliceOfValues[value], key)
	}

	currentIndex := 0
	for i := len(sliceOfValues) - 1; i >= 0; i-- { //  O(a)
		if currentIndex == lenOfRes {
			break
		}
		if sliceOfValues[i] == nil {
			continue
		}

		for _, value := range sliceOfValues[i] {

			res[currentIndex] = value
			currentIndex++
		}

	}

	// O(3a) // худший/лучший/средний
	// mem// O(2a + min(k, b)

	return res

}
