package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	const epsilon = 1.0e-15
	var (
		z = 1.0
		old = x*x
		iters = 0
	)
	for delta := old - z; delta > epsilon; {
		old = z
		z -= (z*z - x) / (2*z)
		delta = math.Abs(old - z)
		iters++
	}
	fmt.Println("Iters: ", iters)
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
