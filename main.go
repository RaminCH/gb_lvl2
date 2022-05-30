// Task1 _ simple
package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan int)
	go func() {
		time.Sleep(time.Second * 3)
		messages <- 1
	}()
	go func() {
		time.Sleep(time.Second * 2)
		messages <- 2
	}()
	go func() {
		time.Sleep(time.Second * 1)
		messages <- 3
	}()
	for i := 0; i < 3; i++ {
		fmt.Println(<-messages)
	}
}
