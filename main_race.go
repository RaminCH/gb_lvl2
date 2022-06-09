// Part6: task3
// Написать многопоточную программу, в которой будет использоваться явный вызов
// планировщика. Выполните трассировку программы

package main

import (
	"fmt"
	"sync"
)

var num = 0

func SuperIncrementer(wg *sync.WaitGroup) {
	//mut.Lock()
	num = num + 2
	//mut.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go SuperIncrementer(&wg)
	}
	wg.Wait()
	fmt.Println("Result is:", num)
}



//Solution (eliminate race condition)

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var num = 0

// func SuperIncrementer(wg *sync.WaitGroup, mut *sync.Mutex) {
// 	mut.Lock()						//pervaya dobravshayasa gorutina zakroyet za soboy shkaf poka ne reshit logiku koda 'num = num + 2' i zatem
// 	num = num + 2						//unlochitsa dav vozmojnost drugim
// 	mut.Unlock()
// 	wg.Done()
// }

// func main() {
// 	var mut sync.Mutex
// 	var wg sync.WaitGroup
// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go SuperIncrementer(&wg, &mut)
// 	}
// 	wg.Wait()
// 	fmt.Println("Result is:", num)
// }
