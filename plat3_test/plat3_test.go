package plat3_test

import (
	"fmt"
	"testing"
)

type Counter struct {
	value int
}

func TestPlat3(t *testing.T) {
	var res = make([]*Counter, 3)
	for i, a := range []Counter{{1}, {2}, {3}} {
		res[i] = &a
	}
	fmt.Println("res:", res[0].value, res[1].value, res[2].value)
}
