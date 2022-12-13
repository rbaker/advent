package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items       []string
	operation   string
	divisor     int
	trueThrow   int
	falseThrow  int
	inspections int
}

var data11, _ = os.ReadFile("resources/day11.txt")
var monkeyList = strings.Split(strings.ReplaceAll(string(data11), "\r\n", "\n"), "\n\n")

func day11part1() int {
	monkeys := prepareMonkeys()

	for r := 0; r < 20; r++ {
		for i, m := range monkeys {
			for _, item := range m.items {
				e := strings.ReplaceAll(m.operation, "old", item)
				monkeys[i].inspections++
				result := eval(e)
				result /= 3

				if result%m.divisor == 0 {
					monkeys[m.trueThrow].items = append(monkeys[m.trueThrow].items, strconv.Itoa(result))
				} else {
					monkeys[m.falseThrow].items = append(monkeys[m.falseThrow].items, strconv.Itoa(result))
				}
			}
			monkeys[i].items = nil
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})

	return monkeys[0].inspections * monkeys[1].inspections
}

func day11part2() int {
	monkeys := prepareMonkeys()
	mod := 1
	for _, m := range monkeys {
		mod *= m.divisor
	}

	for r := 1; r <= 10000; r++ {
		for i, m := range monkeys {
			for _, item := range m.items {
				e := strings.ReplaceAll(m.operation, "old", item)
				monkeys[i].inspections++
				result := eval(e)
				result = result % mod

				if result%m.divisor == 0 {
					monkeys[m.trueThrow].items = append(monkeys[m.trueThrow].items, strconv.Itoa(result))
				} else {
					monkeys[m.falseThrow].items = append(monkeys[m.falseThrow].items, strconv.Itoa(result))
				}
			}
			monkeys[i].items = nil
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})

	return monkeys[0].inspections * monkeys[1].inspections
}

func prepareMonkeys() []monkey {
	monkeyMap := []monkey{}
	for _, m := range monkeyList {
		lines := strings.Split(m, "\n")
		mo := monkey{}
		mo.items = strings.Split(lines[1][18:], ", ")
		mo.operation = lines[2][19:]
		mo.divisor, _ = strconv.Atoi(lines[3][21:])
		mo.trueThrow, _ = strconv.Atoi(lines[4][29:])
		mo.falseThrow, _ = strconv.Atoi(lines[5][30:])

		monkeyMap = append(monkeyMap, mo)
	}
	return monkeyMap
}

func eval(src string) int {
	parts := strings.Split(src, " ")
	first, _ := strconv.Atoi(parts[0])
	second, _ := strconv.Atoi(parts[2])
	if parts[1] == "+" {
		return first + second
	} else {
		return first * second
	}
}
