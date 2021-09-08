package main

import (
	"sort"
)

func test_fknum() {
	array := []int{1, 5, 2, 6, 3, 7, 4}
	commands := [][]int{{2, 5, 3}, {4, 4, 1}, {1, 7, 3}}

	println(finKnum(array, commands))
}

func finKnum(array []int, commands [][]int) []int {
	answer := []int{}
	for _, s := range commands {
		arrK := make([]int, s[1]-(s[0]-1))
		copy(arrK, array[s[0]-1:s[1]])
		sort.Ints(arrK)
		answer = append(answer, arrK[s[2]-1])
	}
	return answer
}
