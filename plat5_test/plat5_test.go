package plat5_test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const numRequests = 10000

var count atomic.Int64

func networkRequest() {
	time.Sleep(time.Millisecond)
	count.Add(1)
}

func TestPlat5(t *testing.T) { // сколько работает по времени
	wg := &sync.WaitGroup{}
	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			networkRequest()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
