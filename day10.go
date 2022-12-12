package main

import (
	"os"
	"strconv"
	"strings"
)

var data10, _ = os.ReadFile("resources/day10.txt")
var instructions = strings.Split(strings.ReplaceAll(string(data10), "\r\n", "\n"), "\n")

func day10part1() int {
	x := 1
	history := []int{}
	registers := []int{20, 60, 100, 140, 180, 220}
	for _, instruction := range instructions {
		if strings.HasPrefix(instruction, "addx") {
			history = append(history, x)
			history = append(history, x)
			value, _ := strconv.Atoi(strings.Split(instruction, " ")[1])
			x += value
		}
		if strings.HasPrefix(instruction, "noop") {
			history = append(history, x)
		}
	}

	total := 0
	for _, register := range registers {
		total += history[register-1] * register
	}

	return total
}

func day10part2() string {
	x := 1
	count := 0
	lines := [6]string{}

	for _, instruction := range instructions {
		if strings.HasPrefix(instruction, "addx") {
			lines[count/40] += drawPixel(count%40, x)
			count++
			lines[count/40] += drawPixel(count%40, x)
			count++
			value, _ := strconv.Atoi(strings.Split(instruction, " ")[1])
			x += value
		}
		if strings.HasPrefix(instruction, "noop") {
			lines[count/40] += drawPixel(count%40, x)
			count++
		}
	}

	image := ""
	for _, l := range lines {
		image = image + l + "\n"
	}
	return image
}

func drawPixel(pos int, x int) string {
	if pos == x-1 || pos == x || pos == x+1 {
		return "#"
	} else {
		return "."
	}
}
