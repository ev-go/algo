package plat7_test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSome(t *testing.T) {
	var w interface{} = "Hello, World!"

	switch v := w.(type) {
	case string:
		fmt.Println("Значение является типом string:", v)
	case int:
		fmt.Println("Значение является типом int:", v)
	default:
		fmt.Println("Неизвестный тип значения")
	}
}

// package main

// import "fmt"

// type Person struct {
//     Name string
//     Age  int
// }

// func set10(val *int) {
//     *val = 10
// }

// func changePerson(person *Person) {
//     *person = Person{
//         Name: "Alice",
//         Age: 10,
//     }
// }

// func main() {
//     person := &Person{
//         Name: "Bob",
//         Age: 43,
//     }
//     fmt.Printf("%+v\n", person)
//     changePerson(person)
//     fmt.Printf("%+v\n", person)
// }

// ========================================


// package main

// import (
//     "fmt"
//     "sync"
//     "time"
// )

// func filter(i int) bool {
//     time.Sleep(time.Millisecond)
//     return i%2 == 0
// }

// func main() {
//     var max int
//     wg := &sync.WaitGroup{}

//     for i := 1000; i > 0; i-- {
//         wg.Add(1)
//         i := i
//         go func() {
//             defer wg.Done()
//             if i > max && filter(i) {
//                 max = i
//             }
//         }()
//     }

//     wg.Wait()
//     fmt.Printf("Maximum is %d", max)
// }

// =======================================

// package main

// import (
//     "fmt"
//     "time"
//     "sync/atomic"
// )

// const numRequests = 100000

// var count int64

// func networkRequest() {
//     time.Sleep(time.Millisecond) // Эмуляция сетевого запроса. //10 sec
//     atomic.AddInt64(&count, 1)
// }

// func main() {

//     wg:= &sync.WaitGroup{}
//     for i := 0; i < numRequests; i++ {
//         wg.Add(1)
//         go func(){
//             defer wg.Done()
//             networkRequest()
//         }()
//     }
//     wg.Wait()
//     fmt.Println(count)
// }


// ======================================

// func getWeather(locationID int) int {
//     time.Sleep(1 * time.Second)
//     return rand.Intn(70) - 30
// }

// func main() {
//     weather := getWeather()

//     go func() {
//         for {
//             time.Sleep(time.Minute)
//             weather = getWeather()
//         }
//     }()


//     http.HandleFunc("/weather/highload", func(resp http.ResponseWriter, req *http.Request) {

//     })
// }

// ===============================

package main

import (
"fmt"
"math/rand"
"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Есть функция, работающая неопределённо долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос).
func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)

	return rnd
}

// Нужно изменить функцию обёртку, которая будет работать с заданным таймаутом (например, 1 секунду).
// Если "длинная" функция отработала за это время - отлично, возвращаем результат.
// Если нет - возвращаем ошибку. Результат работы в этом случае нам не важен.
//
// Дополнительно нужно измерить, сколько выполнялась эта функция (просто вывести в лог).
// Сигнатуру функцию обёртки менять можно.
func predictableFunc(ctx context.Context) (int64, error) { // ctx уже с таймаутом
	resultCh := make(chan int64, 1)
	go func(){
		resultCh <-unpredictableFunc()
	}()

	select{
	case res := <-resultCh:
		return res, nil
	case <- ctx.Done():
		return 0, ctx.Err()
	}
}

func main() {
	fmt.Println("started")

	fmt.Println(predictableFunc())
}

// +=================


var a interface{}
var b = interface{}(nil)

a == b
