package main

import "os"

var data6, _ = os.ReadFile("resources/day6.txt")

func day6part1() int {
	return getUniqueLocation(4)
}

func day6part2() int {
	return getUniqueLocation(14)
}

func getUniqueLocation(lookBack int) int {
	for i := lookBack - 1; i <= len(data6); i++ {
		letterMap := make(map[byte]bool)
		for j := 0; j < lookBack; j++ {
			letterMap[data6[i-j]] = true
		}
		if len(letterMap) == lookBack {
			return i + 1
		}
	}
	return 0
}
