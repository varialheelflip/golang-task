package main

import "fmt"

func main() {
	var shape1 Shape = Rectangle{Width: 1, Length: 2}
	fmt.Println(shape1.Perimeter())
	fmt.Println(shape1.Area())
	var shape2 Shape = &Circle{Diameter: 2}
	fmt.Println(shape2.Perimeter())
	fmt.Println(shape2.Area())
}
