package main

import ("fmt"
		"math/rand"
	"math"
)

func foo() {
	fmt.Println("the square root of 4 is", math.Sqrt(4))
}
func main() {
	fmt.Println("A number from 1-100 is", rand.Intn(100))
	foo()
}
