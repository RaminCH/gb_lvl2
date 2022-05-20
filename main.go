package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	values := 1000
	count := 0
	// нужно вспомнить про горутины
	var wg sync.WaitGroup
	for i := 1; i <= values; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
            // time.Sleep(1 * time.Second)
			count += 1
            time.Sleep(1 * time.Second)
			fmt.Println(i, count)
			wg.Done()
			// и про копирование переменной v, иначе трижды "c" отпечатает
		}(i, &wg)
	}
	wg.Wait()
}
