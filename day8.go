package main

import (
	"math"
	"os"
	"strconv"
	"strings"
)

var data8, _ = os.ReadFile("resources/day8.txt")
var treelines = strings.Split(strings.ReplaceAll(string(data8), "\r\n", "\n"), "\n")

func day8part1() int {

	counted := make(map[[2]int]bool)

	for rowNum, treeline := range treelines {
		leftHighest, rightHighest := -1, -1

		for l, r := 0, len(treeline)-1; l < len(treeline); l, r = l+1, r-1 {
			left, _ := strconv.Atoi(string(treeline[l]))
			right, _ := strconv.Atoi(string(treeline[r]))

			if left > leftHighest {
				counted[[2]int{rowNum, l}] = true
				leftHighest = left
			}
			if right > rightHighest {
				counted[[2]int{rowNum, r}] = true
				rightHighest = right
			}
		}
	}

	for l := 0; l < len(treelines[0]); l++ {
		topHighest, bottomHighest := -1, -1

		for t, b := 0, len(treelines)-1; t < len(treelines); t, b = t+1, b-1 {
			top, _ := strconv.Atoi(string(treelines[t][l]))
			bottom, _ := strconv.Atoi(string(treelines[b][l]))

			if top > topHighest {
				topHighest = top
				counted[[2]int{t, l}] = true
			}
			if bottom > bottomHighest {
				bottomHighest = bottom
				counted[[2]int{b, l}] = true
			}
		}
	}

	return len(counted)
}

func day8part2() int {

	biggest := 0

	for rowNum := 1; rowNum < len(treelines)-1; rowNum++ {
		treeline := treelines[rowNum]

		for i := 1; i < len(treeline)-1; i++ {
			n, _ := strconv.Atoi(string(treeline[i]))
			search := [4]int{}
			countL, countR, countT, countB := 1, 1, 1, 1
			for l, r, t, b := i, i, rowNum, rowNum; search[0] == 0 || search[1] == 0 || search[2] == 0 || search[3] == 0; l, r, t, b = l-1, r+1, t-1, b+1 {

				if l > 0 && search[0] == 0 {
					left, _ := strconv.Atoi(string(treeline[l-1]))

					if left >= n {
						search[0] = countL
					} else {
						countL++
					}

				} else if search[0] == 0 {
					search[0] = i
				}

				if r < len(treeline)-1 && search[1] == 0 {
					right, _ := strconv.Atoi(string(treeline[r+1]))

					if right >= n {
						search[1] = countR
					} else {
						countR++
					}

				} else if search[1] == 0 {
					search[1] = len(treeline) - i
				}

				if t > 0 && search[2] == 0 {
					top, _ := strconv.Atoi(string(treelines[t-1][i]))

					if top >= n {
						search[2] = countT
					} else {
						countT++
					}

				} else if search[2] == 0 {
					search[2] = rowNum
				}

				if b < len(treelines)-1 && search[3] == 0 {
					bottom, _ := strconv.Atoi(string(treelines[b+1][i]))

					if bottom >= n {
						search[3] = countB
					} else {
						countB++
					}

				} else if search[3] == 0 {
					search[3] = len(treelines) - 1 - rowNum
				}

			}
			biggest = int(math.Max(float64(biggest), float64(search[0]*search[1]*search[2]*search[3])))
		}

	}

	return biggest
}
