package main

import (
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var data7, _ = os.ReadFile("resources/day7.txt")
var output = strings.Split(strings.ReplaceAll(string(data7), "\r\n", "\n"), "\n")

func day7part1() int {
	dirMap := getDirectorySizes()

	total := 0
	for _, v := range dirMap {
		if v <= 100000 {
			total += v
		}
	}
	return total
}

func day7part2() int {
	dirMap := getDirectorySizes()
	required := dirMap[""] + 30000000 - 70000000

	smallest := dirMap[""]
	for _, v := range dirMap {
		if v >= required {
			smallest = int(math.Min(float64(v), float64(smallest)))
		}
	}
	return smallest
}

func getDirectorySizes() map[string]int {
	currentDirectory := ""
	cdRegex := `\$ cd ([a-zA-Z0-9\.]+)`
	fileRegex := `([0-9]+) (.*)`
	dirMap := make(map[string]int)

	for _, out := range output {
		if match, _ := regexp.MatchString(cdRegex, out); match {
			path := regexp.MustCompile(cdRegex).FindAllStringSubmatch(out, -1)
			changeDirectory(path[0][1], &currentDirectory)
		} else if match, _ := regexp.MatchString(fileRegex, out); match {
			file := regexp.MustCompile(fileRegex).FindAllStringSubmatch(out, -1)
			fileSize, _ := strconv.Atoi(file[0][1])

			path := currentDirectory
			for i := 0; strings.Contains(path, "/"); i++ {
				dirMap[path] = dirMap[path] + fileSize
				path = path[:strings.LastIndex(path, "/")]
			}
			dirMap[path] = dirMap[path] + fileSize

		}

	}
	return dirMap
}

func changeDirectory(path string, currentDirectory *string) {
	if path == ".." {
		*currentDirectory = (*currentDirectory)[:strings.LastIndex(*currentDirectory, "/")]
	} else {
		*currentDirectory = *currentDirectory + "/" + path
	}
}
