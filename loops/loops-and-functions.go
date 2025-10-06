package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	z := 1.0

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("%v: %v\n", i+1, z)
	}

	return fmt.Sprint(z)
}

func countLoop(x float64) int {
	z := 1.0
	i := 0

	for {
		i++
		next := z - (z*z-x)/(2*z)

		if math.Abs(next-z) < 1e-10 {
			break
		}

		z = next
	}

	return i
}

func main() {
	x := 16.0

	fmt.Println("Loops: ", countLoop(x))
	fmt.Println(sqrt(x))
	fmt.Println(math.Sqrt(x))
}
