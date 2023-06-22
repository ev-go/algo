package plat4_test

import (
	"fmt"
	"runtime"
	"testing"
)

func TestPlat4(t *testing.T) {
	var x int
	threads := runtime.GOMAXPROCS(1)
	for i := 0; i < threads; i++ {
		go func() {
			for {
				x++
			}
		}()
	}
	//time.Sleep(time.Second) //точка переключения кгонтекста
	runtime.Gosched()
	fmt.Print("x = ", x)
}
