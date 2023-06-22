package plat5_test

import (
	"fmt"
	"reflect"
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
	fmt.Println(count.Load())
	v := reflect.ValueOf(count)
	field := v.FieldByName("v")
	if field.IsValid() {
		countValue := field.Int()
		fmt.Println(countValue)
	} else {
		fmt.Println("Field not found")
	}

	var w interface{} = "Hello, World!"

	switch g := w.(type) {
	case string:
		fmt.Println("Значение является типом string:", g)
	case int:
		fmt.Println("Значение является типом int:", g)
	default:
		fmt.Println("Неизвестный тип значения")
	}
}
