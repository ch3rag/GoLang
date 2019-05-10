package main

import (
	"fmt"
	"math/rand"
	"math"
	"time"
)

func main() {
	fmt.Println("My Favourite Random Number Is:  ", rand.Intn(10))
	fmt.Println("Current time is: ", time.Now())
	fmt.Println(math.Sqrt(4))
	fmt.Println(math.Pi)
}