package main

import (
	"fmt"
	// "time"
	// "math/rand"
	// "math"
	"math/cmplx"
)

// var c, python, java bool // Define boolean
// var i, j int = 1, 2 // Variables with initializers

// func add(x int, y int) int {
// 	return x + y
// }

// func add1(x,y int) int {
// 	return x + y
// }

// func swap(x, y string) (string, string) {
// 	return y, x
// }

// func split(sum int) (x, y int) {
// 	x = sum * 4 / 9
// 	y = sum - x
// 	return
// }

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// func main() {
// 	var i, j int = 1, 2
// 	k := 3
	
// 	fmt.Println("Hello, World")
// 	fmt.Println("This time is", time.Now())
// 	fmt.Println("My favorite number is", rand.Intn(10))

// 	fmt.Printf("Now you have %g problem.\n", math.Sqrt(7)) // Use Printf

// 	fmt.Println(math.Pi)

// 	fmt.Println(add(42, 13))
// 	fmt.Println(add1(42, 13))

// 	a, b := swap("hello", "world")
// 	fmt.Println(a, b)
// 	fmt.Println(split(17))
// 	var c, python, java = true, false, "no!"
// 	fmt.Println(i, j, k, c, python, java)

// }

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}