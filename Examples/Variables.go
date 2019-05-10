package main

import (
	"fmt"
	"math/cmplx"
)

var c, python, java bool
var x, y int = 2, 3

// Variable Declaration List
var (
	test 	bool 		= false
	maxInt	uint64 		= 1 << 64 - 1
	z		complex128	= cmplx.Sqrt(-5 + 12i)
)
func main() {
	var i int
	var name string = "Chirag"
	fmt.Println(c, python, java, i)
	fmt.Println(x, y)
	fmt.Println(test, maxInt, z)
	fmt.Println(name)
}