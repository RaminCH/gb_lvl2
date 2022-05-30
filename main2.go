//Unlocking Mutex using defer

package main

import (
	"errors"
	"fmt"
	"sync"
)

var mu sync.Mutex
var sum = 0

func Add(i int) int {
	mu.Lock()
	defer mu.Unlock()

	return sum + i
}

func Divide(val int, i int) (int, error) {
	mu.Lock()
	defer mu.Unlock()

	if i == 0 {
		errors.New("Can't divide by zero!")
	}

	res := val / i
	return res, nil
}

func main() {
	a := Add(5)
	fmt.Println(a)
	b, _ := Divide(6, 2)
	fmt.Println(b)
}
