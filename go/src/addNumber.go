package main

import "fmt"

func testAddNumber() {
	fmt.Println(addNumber([]int{5, 8, 4, 0, 6, 7, 9}))
}

func addNumber(numbers []int) int {
	chkNum := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
		6: 6,
		7: 7,
		8: 8,
		9: 9,
		0: 0,
	}

	for _, num := range numbers {
		delete(chkNum, num)
	}
	rs := 0
	for _, num := range chkNum {
		rs += num
	}
	return rs
}
