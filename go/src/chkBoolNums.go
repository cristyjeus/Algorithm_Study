package main

func testBoolNums() {
	absolutes := []int{1, 2, 3}
	signs := []bool{false, false, true}
	println(chkBoolNums(absolutes, signs))
}

func chkBoolNums(absolutes []int, signs []bool) int {
	res := 0
	for i, ab := range absolutes {
		if signs[i] == true {
			res += ab
		} else {
			res -= ab
		}
	}

	return res
}
