package main

import (
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var data1, _ = os.ReadFile("resources/day1.txt")
var elves = strings.Split(strings.ReplaceAll(string(data1), "\r\n", "\n"), "\n\n")

func day1part1() int {

	biggest := 0

	for _, elf := range elves {
		cals := strings.Split(strings.ReplaceAll(elf, "\r\n", "\n"), "\n")
		total := 0
		for _, cal := range cals {
			c, _ := strconv.Atoi(cal)
			total += c
		}
		biggest = int(math.Max(float64(biggest), float64(total)))
	}

	return biggest
}

func day1part2() int {

	var amounts []int

	for _, elf := range elves {
		cals := strings.Split(strings.ReplaceAll(elf, "\r\n", "\n"), "\n")
		total := 0
		for _, cal := range cals {
			c, _ := strconv.Atoi(cal)
			total += c
			amounts = append(amounts, total)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(amounts)))
	return amounts[0] + amounts[1] + amounts[2]
}
