package main

import (
	"os"
	"strings"
)

var colA = "ABC"
var colB = "XYZ"

var data2, _ = os.ReadFile("resources/day2.txt")
var rounds = strings.Split(strings.ReplaceAll(string(data2), "\r\n", "\n"), "\n")

func day2part1() int {

	totalScore := 0
	pointMatrix := [][]int{
		{4, 8, 3},
		{1, 5, 9},
		{7, 2, 6},
	}

	for _, round := range rounds {
		players := strings.Split(round, " ")
		totalScore += pointMatrix[strings.Index(colA, players[0])][strings.Index(colB, players[1])]
	}
	return totalScore

}

func day2part2() int {

	totalScore := 0
	pointMatrix := [][]int{
		{3, 4, 8},
		{1, 5, 9},
		{2, 6, 7},
	}

	for _, round := range rounds {
		players := strings.Split(round, " ")
		totalScore += pointMatrix[strings.Index(colA, players[0])][strings.Index(colB, players[1])]
	}
	return totalScore

}
