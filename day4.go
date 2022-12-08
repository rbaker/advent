package main

import (
	"os"
	"strconv"
	"strings"
)

var data4, _ = os.ReadFile("resources/day4.txt")
var pairings = strings.Split(strings.ReplaceAll(string(data4), "\r\n", "\n"), "\n")

func day4part1() int {
	contained := 0

	for _, pairing := range pairings {

		elf1min, elf1max, elf2min, elf2max := getElves(pairing)

		if (elf1min <= elf2min && elf1max >= elf2max) || (elf1min >= elf2min && elf1max <= elf2max) {
			contained++
		}
	}

	return contained
}

func day4part2() int {
	contained := 0

	for _, pairing := range pairings {

		elf1min, elf1max, elf2min, elf2max := getElves(pairing)

		if (elf1min <= elf2min && elf1max >= elf2min) || (elf1min >= elf2min && elf1min <= elf2max) {
			contained++
		}
	}

	return contained
}

func getElves(pairing string) (int, int, int, int) {
	elves := strings.Split(pairing, ",")
	elf1 := strings.Split(elves[0], "-")
	elf2 := strings.Split(elves[1], "-")
	elf1min, _ := strconv.Atoi(elf1[0])
	elf1max, _ := strconv.Atoi(elf1[1])
	elf2min, _ := strconv.Atoi(elf2[0])
	elf2max, _ := strconv.Atoi(elf2[1])
	return elf1min, elf1max, elf2min, elf2max
}
