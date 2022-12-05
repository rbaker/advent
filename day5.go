package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

var data5, _ = os.ReadFile("resources/day5.txt")
var lines = strings.Split(strings.ReplaceAll(string(data5), "\r\n", "\n"), "\n")
var rexExp = regexp.MustCompile(`(?m)move (?P<amount>[0-9]*) from (?P<col1>[0-9]*) to (?P<col2>[0-9]*)`)

func day5part1() string {

	piles := generatePiles()
	for i := 10; i < len(lines); i++ {
		result := generateRegexResultMap(lines[i])
		for j := 0; j < result["amount"]; j++ {
			last := piles[result["col1"]-1][len(piles[result["col1"]-1])-1]
			piles[result["col1"]-1] = piles[result["col1"]-1][:len(piles[result["col1"]-1])-1]
			piles[result["col2"]-1] = append(piles[result["col2"]-1], last)
		}
	}
	return getResult(piles)
}

func day5part2() string {

	piles := generatePiles()
	for i := 10; i < len(lines); i++ {
		result := generateRegexResultMap(lines[i])
		end := piles[result["col1"]-1][len(piles[result["col1"]-1])-result["amount"]:]
		piles[result["col1"]-1] = piles[result["col1"]-1][:len(piles[result["col1"]-1])-result["amount"]]
		piles[result["col2"]-1] = append(piles[result["col2"]-1], end...)
	}

	return getResult(piles)

}

func generatePiles() [9][]string {
	piles := [9][]string{}

	for i := 0; i < 8; i++ {
		line := lines[i]
		for j := 0; j < 9; j++ {
			position := j*4 + 1
			char := line[position]
			if char != 32 {
				piles[j] = append([]string{string(char)}, piles[j]...)
			}
		}
	}
	return piles
}

func generateRegexResultMap(line string) map[string]int {
	match := rexExp.FindStringSubmatch(line)
	result := make(map[string]int)
	for i, name := range rexExp.SubexpNames() {
		if i != 0 && name != "" {
			result[name], _ = strconv.Atoi(match[i])
		}
	}
	return result
}

func getResult(piles [9][]string) string {
	result := ""

	for _, pile := range piles {
		result = result + pile[len(pile)-1]
	}

	return result
}
