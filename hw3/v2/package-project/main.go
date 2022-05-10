package main

import (
	"fmt"
	_ v2/package-project/rectangle
)

func main() {
	rectangle := Rectangle{
		breadth: 10,
		len:     8,
	}
	fmt.Println("Area of the rectangle:", rectangle, " is: ",
		rectangle.Area())
}
