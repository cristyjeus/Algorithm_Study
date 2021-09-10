package main

import (
	"fmt"
)

func testLottos() {
	lottos, win_nums := []int{44, 1, 0, 0, 31, 25}, []int{31, 10, 45, 1, 6, 19}

	fmt.Println(chkLottoMaxMin(lottos, win_nums))
}

func chkLottoMaxMin(lottos []int, win_nums []int) []int {
	rs := []int{}
	mapRank := []int{6, 6, 5, 4, 3, 2, 1}
	cnt := 0
	zeroCnt := 0
	for _, lotto := range lottos {
		for _, num := range win_nums {
			if num == lotto {
				cnt++
			}
		}
		if lotto == 0 {
			zeroCnt++
		}
	}
	rs = append(rs, mapRank[cnt+zeroCnt])
	rs = append(rs, mapRank[cnt])
	return rs
}
