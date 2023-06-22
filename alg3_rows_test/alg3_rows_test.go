package alg3_rows_test

import (
	"fmt"
	"testing"
)

//[1, 0, 0, 0, 1], k = 1
//true // [1, 0, 1, 0, 1]
//
//[0, 0, 1, 0], k = 1
//true // [1, 0, 1, 0]
//
//[0, 1, 0, 0, 1, 0], k = 2
//false
func TestAlg3(t *testing.T) {

	firstRow := []int{1, 0, 0, 0, 1}
	secondRow := []int{0, 0, 1, 0}
	thirdRow := []int{0, 1, 0, 0, 1, 0}
	fourthRow := []int{0}
	fifthRow := []int{1}

	fmt.Println(Put(firstRow, 1))
	fmt.Println(Put(secondRow, 1))
	fmt.Println(Put(thirdRow, 2))
	fmt.Println(Put(fourthRow, 1))
	fmt.Println(Put(fifthRow, 1))

}

func Put(inpRow []int, k int) ([]int, bool) {
	res := make([]int, len(inpRow)+2)
	resBool := false
	copy(res[1:len(res)-1], inpRow)

	for i, _ := range inpRow {
		if k == 0 {
			resBool = true
			break
		}

		if (res[i] + res[i+1] + res[i+2]) == 0 {
			res[i+1] = 1
			k--
		}

	}
	if k == 0 {
		resBool = true
	}
	return res[1 : len(res)-1], resBool
}
