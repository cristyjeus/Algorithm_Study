package main

import (
	"fmt"
	"math"
)

func testSQRT() {
	input := []int64{121, 3}
	expect := []int64{144, -1}
	for i := range input {
		result := sqrtAlgo(input[i])
		if result != expect[i] {
			fmt.Printf("Test%d: Wrong result", i+1)
			fmt.Printf("expect:", expect[i])
			fmt.Printf("result:", result)
		}
	}
}

func sqrtAlgo(n int64) int64 {
	sqrt := int64(math.Sqrt(float64(n)))
	if sqrt*sqrt == n {
		sqrt++
		return sqrt * sqrt
	}

	return -1
}
