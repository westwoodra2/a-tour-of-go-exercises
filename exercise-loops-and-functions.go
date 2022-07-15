package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	prev := 0.0
	curr := 1.0
	i := 0

	for curr != prev {
		prev = curr
		curr -= (curr*curr - x) / (2 * curr)
		i += 1
	}

	return curr
}

func main() {
	fmt.Println(Sqrt(2))
}
