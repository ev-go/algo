package plat6_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"testing"
	"time"
)

// Есть n кол-во URL передаем в функцию.
// Сделать n запросов на каждый Url (по одному на каждый). Выполнение асинхронное
// Если кеакойто из запросов не успел вернуть значение за 1 сек
// то не дожидаемся и возвращаем / записываем результат запроса ошибку err

//map ключ - url, value - что получил по url или timeout exit если не уложился в таймаут

func TestUrls(t *testing.T) {

}
func urlCheck(urls []string) map[string]string {
	resultMap := make(map[string]string, len(urls))
	wg := &sync.WaitGroup{}
	wg.Add(len(urls))
	mux := &sync.Mutex{}
	for _, val := range urls {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		resultChan := make(chan string, 1)
		go func(resultChan chan string, val string) {

			var respByte []byte
			resp, err := http.Get(val)
			if err != nil {
				resultChan <- fmt.Sprintf("smth wrong: %v ", err)

			} else {
				respByte, err = io.ReadAll(resp.Body)
				if err != nil {
					resultChan <- fmt.Sprintf("smth wrong: %v ", err)
				}
				resultChan <- string(respByte)
			}
		}(resultChan, val)
		go func(val string) {
			defer wg.Done()
			select {
			case res := <-resultChan:
				mux.Lock()
				resultMap[val] = res
				mux.Unlock()
			case <-ctx.Done(): // 2 условия. cancel - закрывает. timeout - возвращается
				mux.Lock()
				resultMap[val] = "timeout exit"
				mux.Unlock()
			}
		}(val)

	}
	wg.Wait()
	return resultMap
}

type Mem string

func main() {
	var w interface{} = Mem("Hello, World!")

	if mem, ok := w.(Mem); ok {
		fmt.Println(string(mem))
	} else {
		fmt.Println("Не удалось привести к типу Mem")
	}
}
var w interface{} = "Hello, World!"
switch v := w.(type) {
case string:
fmt.Println("Значение является типом string:", v)
case int:
fmt.Println("Значение является типом int:", v)
default:
fmt.Println("Неизвестный тип значения")
}