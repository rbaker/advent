package main

import (
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var data13, _ = os.ReadFile("resources/day13.txt")

type numList struct {
	number int
	list   []numList
}

func day13part1() int {
	packagePairs := strings.Split(strings.ReplaceAll(string(data13), "\r\n", "\n"), "\n\n")
	total := 0

	for i, pair := range packagePairs {
		left, right := strings.Split(pair, "\n")[0], strings.Split(pair, "\n")[1]

		l := parseList(left)
		r := parseList(right)

		if compare(l[0].list, r[0].list) == 1 {
			total += i + 1
		}
	}
	return total
}

func day13part2() int {
	packages := strings.Split(strings.ReplaceAll(strings.ReplaceAll(string(data13), "\r\n", "\n"), "\n\n", "\n"), "\n")
	packages = append(packages, "[[2]]")
	packages = append(packages, "[[6]]")
	allPackages := [][]numList{}
	for _, p := range packages {
		allPackages = append(allPackages, parseList(p))
	}
	sort.Slice(allPackages, func(i, j int) bool {
		return compare(allPackages[i], allPackages[j]) > 0
	})
	total := 1
	for i, p := range allPackages {
		if len(p[0].list) == 1 && len(p[0].list[0].list) == 1 && (p[0].list[0].list[0].number == 2 || p[0].list[0].list[0].number == 6) {
			total *= i + 1
		}
	}

	return total
}

func compare(left []numList, right []numList) int {

	largest := math.Max(float64(len(left)), float64(len(right)))

	for i := 0; i < int(largest); i++ {
		if i > len(left)-1 {
			return 1
		}
		if i > len(right)-1 {
			return -1
		}
		if left[i].number >= 0 && right[i].number >= 0 {
			if left[i].number < right[i].number {
				return 1
			} else if left[i].number > right[i].number {
				return -1
			}
		}
		if left[i].number == -1 && right[i].number == -1 {
			c := compare(left[i].list, right[i].list)
			if c != 0 {
				return c
			}
		}
		if left[i].number == -1 && right[i].number >= 0 {
			c := compare(left[i].list, []numList{{number: right[i].number}})
			if c != 0 {
				return c
			}
		}
		if left[i].number >= 0 && right[i].number == -1 {
			c := compare([]numList{{number: left[i].number}}, right[i].list)
			if c != 0 {
				return c
			}
		}
	}
	return 0
}

func parseList(list string) []numList {
	nums := make(map[int][]numList)
	current := 0
	currentNum := ""
	for _, c := range list {
		if c == ',' && currentNum != "" {
			n, _ := strconv.Atoi(currentNum)
			nums[current] = append(nums[current], numList{number: n, list: nil})
			currentNum = ""
		} else if c == '[' {
			current++
		} else if c == ']' {
			if currentNum != "" {
				n, _ := strconv.Atoi(currentNum)
				nums[current] = append(nums[current], numList{number: n, list: nil})
				currentNum = ""
			}
			nums[current-1] = append(nums[current-1], numList{number: -1, list: nums[current]})
			nums[current] = nil
			current--
		} else {
			_, err := strconv.Atoi(string(c))
			if err == nil {
				currentNum += string(c)
			}
		}
	}
	return nums[0]
}
