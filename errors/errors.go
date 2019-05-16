package main

import (
	"fmt"
	"math"
)

// ErrNegativeSqrt implements the error
// interface to throw and error for a
// negative square root
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("%s%v", "cannot Sqrt negative number:", float64(e))
}

// Sqrt estimates the square root of a given
// number
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	const epsilon = 1.0e-15
	z := 1.0
	old := x * x
	for delta := old - z; delta > epsilon; {
		old = z
		z -= (z*z - x) / (2 * z)
		delta = math.Abs(old - z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
