package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func swap(x, y int) (int, int) {
	return y, x
}

func squareAndCube(n int) (square int, cube int) {
	square = n * n
	cube = square * n
	return
}

func main() {
	fmt.Println(add(3, 4))
	fmt.Println(multiply(3, 4))

	//var x, y int = swap(4, 3)
	// OR
	 x, y := swap(4, 5)


	fmt.Println(x, y)

	a, b := squareAndCube(3)

	fmt.Println(a, b)
}