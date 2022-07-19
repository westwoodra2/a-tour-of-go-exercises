package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number:%v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	prev := 0.0
	curr := 1.0
	i := 0

	for curr != prev {
		prev = curr
		curr -= (curr*curr - x) / (2 * curr)
		i += 1
		//fmt.Println(i, prev, curr)
	}

	return curr, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
