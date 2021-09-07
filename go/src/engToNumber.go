package main

import (
	"strconv"
	"strings"
)

func testEngToNum() {
	println(engToNumAlgo1("one4seveneight"))
	println(engToNumAlgo2("three10zerotwothree"))
}

func engToNumAlgo1(s string) int {
	//println(s)
	strNum := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "zero"}
	intNum := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	a := s

	for i := 0; i < len(strNum); i++ {
		a = strings.ReplaceAll(a, strNum[i], intNum[i])
	}
	b, _ := strconv.Atoi(a)

	return b
}

func engToNumAlgo2(s string) int {
	f := strings.NewReplacer(
		"zero", "0",
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
	)
	i, _ := strconv.Atoi(f.Replace(s))
	return i
}
