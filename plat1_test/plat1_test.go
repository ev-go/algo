package plat1_test

import (
	"fmt"
	"testing"
)

func TestPlat1(t *testing.T) {
	i := 1

	defer fmt.Printf("First defer: %d\n", i) // 1
	i++

	defer func() {
		fmt.Printf("Second defer: %d\n", i) // //third 4
	}()

	i++

	defer func(i int) {
		fmt.Printf("Third defer: %d\n", i) //second 3
	}(i)

	i++

	defer fmt.Printf("First defer: %d\n", i) //first 4
}
