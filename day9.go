package main

import (
	"os"
	"strconv"
	"strings"
)

var data9, _ = os.ReadFile("resources/day9.txt")
var ropeMoves = strings.Split(strings.ReplaceAll(string(data9), "\r\n", "\n"), "\n")

func day9part1() int {
	return getTailMap(2)
}

func day9part2() int {
	return getTailMap(10)
}

func getTailMap(size int) int {
	pos := make(map[[2]int]int)
	coords := make([][2]int, size)

	for _, move := range ropeMoves {
		parts := strings.Split(move, " ")
		c, _ := strconv.Atoi(parts[1])
		for i := 0; i < c; i++ {
			moveHead(parts[0], &coords[0])
			for j := 0; j < size-1; j++ {
				if !isTouching(coords[j], coords[j+1]) {
					moveTail(&coords[j], &coords[j+1])
				}
			}
			pos[coords[size-1]]++
		}

	}
	return len(pos)
}

func isTouching(target [2]int, current [2]int) bool {
	verticalDifference := target[0] - current[0]
	horizontalDifference := target[1] - current[1]
	return verticalDifference >= -1 && verticalDifference <= 1 && horizontalDifference >= -1 && horizontalDifference <= 1
}

func moveHead(direction string, head *[2]int) {
	switch direction {
	case "R":
		head[0]++
	case "L":
		head[0]--
	case "U":
		head[1]++
	case "D":
		head[1]--
	}
}
func moveTail(target *[2]int, current *[2]int) {
	if target[0] > current[0] {
		current[0]++
	}
	if target[0] < current[0] {
		current[0]--
	}
	if target[1] > current[1] {
		current[1]++
	}
	if target[1] < current[1] {
		current[1]--
	}
}
