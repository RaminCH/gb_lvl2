// Part6: Task1
// Написать программу, которая использует мьютекс для безопасного доступа к данным
// из нескольких потоков. Выполните трассировку программы

package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

const count = 1000

func main() {

	trace.Start(os.Stderr)
	defer trace.Stop()

	var (
		counter int
		lock    sync.Mutex
		wg      sync.WaitGroup
	)
	wg.Add(count)
	for i := 0; i < count; i += 1 {
		go func() {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			counter += 1
		}()
	}

	wg.Add(2000)
	for i := 2000; i > 0; i -= 1 {
		go func() {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			counter -= 1
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

// ramin@ramin:~/go/src/GB/еуыеы$ go run -race main1.go 2>trace.out
// -1000
// ramin@ramin:~/go/src/GB/еуыеы$ go tool trace trace.out
