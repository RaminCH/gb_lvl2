// Написать функцию, которая принимает на вход структуру in (struct или кастомную struct) и
// values map[string]interface{} (key - название поля структуры, которому нужно присвоить value
// этой мапы). Необходимо по значениям из мапы изменить входящую структуру in с помощью
// пакета reflect. Функция может возвращать только ошибку error. Написать к данной функции
// тесты (чем больше, тем лучше - зачтется в плюс).

package main

import (
	"encoding/json"
	"fmt"
)

type MyData struct {
	One   int
	Two   string
	Three int
}

func main() {
	in := &MyData{One: 1, Two: "second"}

	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(in)
	json.Unmarshal(inrec, &inInterface)

	// iterate through inrecs
	for field, val := range inInterface {
		fmt.Println("KV Pair: ", field, val)
	}
}
