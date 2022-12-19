package main

import (
	"math"
	"os"
	"strconv"
	"strings"
)

var data14, _ = os.ReadFile("resources/day14.txt")
var rockLines = strings.Split(strings.ReplaceAll(string(data14), "\r\n", "\n"), "\n")

func day14part1() int {
	total := 0
	o, max := prepareMap()
	edge := false
	for !edge {
		atRest := false
		pos := [2]int{500, 0}
		for !atRest {
			if pos[1] > max {
				edge = true
				break
			}

			if !o[[2]int{pos[0], pos[1] + 1}] {
				pos = [2]int{pos[0], pos[1] + 1}
			} else if !o[[2]int{pos[0] - 1, pos[1] + 1}] {
				pos = [2]int{pos[0] - 1, pos[1] + 1}
			} else if !o[[2]int{pos[0] + 1, pos[1] + 1}] {
				pos = [2]int{pos[0] + 1, pos[1] + 1}
			} else {
				o[[2]int{pos[0], pos[1]}] = true
				atRest = true
				total++
			}
		}
	}
	return total
}

func day14part2() int {
	total := 0
	o, max := prepareMap()
	top := false
	for !top {
		atRest := false
		pos := [2]int{500, 0}
		for !atRest {
			if o[[2]int{500, 0}] {
				top = true
				break
			}

			if pos[1]+1 == max+2 {
				o[[2]int{pos[0], pos[1]}] = true
				atRest = true
				total++
			} else if !o[[2]int{pos[0], pos[1] + 1}] {
				pos = [2]int{pos[0], pos[1] + 1}
			} else if !o[[2]int{pos[0] - 1, pos[1] + 1}] {
				pos = [2]int{pos[0] - 1, pos[1] + 1}
			} else if !o[[2]int{pos[0] + 1, pos[1] + 1}] {
				pos = [2]int{pos[0] + 1, pos[1] + 1}
			} else {
				o[[2]int{pos[0], pos[1]}] = true
				atRest = true
				total++
			}
		}
	}
	return total
}

func prepareMap() (map[[2]int]bool, int) {
	occupiedMap := make(map[[2]int]bool)
	max := 0

	for _, rocks := range rockLines {
		coords := strings.Split(rocks, " -> ")
		for i := 0; i < len(coords)-1; i++ {
			start := strings.Split(coords[i], ",")
			end := strings.Split(coords[i+1], ",")

			startX, _ := strconv.Atoi(start[0])
			startY, _ := strconv.Atoi(start[1])
			endX, _ := strconv.Atoi(end[0])
			endY, _ := strconv.Atoi(end[1])

			if startY > max || endY > max {
				max = int(math.Max(float64(startY), float64(endY)))
			}

			if startX == endX {
				s, l := math.Min(float64(startY), float64(endY)), math.Max(float64(startY), float64(endY))
				for j := s; j <= l; j++ {
					occupiedMap[[2]int{startX, int(j)}] = true
				}
			}

			if startY == endY {
				s, l := math.Min(float64(startX), float64(endX)), math.Max(float64(startX), float64(endX))
				for j := s; j <= l; j++ {
					occupiedMap[[2]int{int(j), startY}] = true
				}
			}

		}
	}
	return occupiedMap, max
}
