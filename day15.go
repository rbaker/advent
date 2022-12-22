package main

import (
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type sensor struct {
	sensorLocation [2]int
	beaconLocation [2]int
	distance       int
}

var data15, _ = os.ReadFile("resources/day15.txt")
var sensors = strings.Split(strings.ReplaceAll(string(data15), "\r\n", "\n"), "\n")
var sensorList = prepareBeacons()

func day15part1() int {
	locations := make(map[[2]int]bool)
	y := 2000000
	for _, s := range sensorList {
		minSlice, maxSlice := s.sensorLocation[1]-s.distance, s.sensorLocation[1]+s.distance
		if y > minSlice && y < maxSlice {
			distanceFromY := int(math.Abs(float64(y - s.sensorLocation[1])))
			locations[[2]int{s.sensorLocation[0], y}] = true
			for i := 0; i <= s.distance-distanceFromY; i++ {
				locations[[2]int{s.sensorLocation[0] + i, y}] = true
				locations[[2]int{s.sensorLocation[0] - i, y}] = true
			}
		}
	}
	for _, s := range sensorList {
		delete(locations, s.beaconLocation)
	}
	return len(locations)
}

func day15part2() int {
	horizontal := [][2]int{}
	maximumRow := 4000000
	for y := 0; y < maximumRow; y++ {
		for _, s := range sensorList {
			minSlice, maxSlice := s.sensorLocation[1]-s.distance, s.sensorLocation[1]+s.distance
			if y > minSlice && y < maxSlice {
				distanceFromY := int(math.Abs(float64(y - s.sensorLocation[1])))
				min := int(math.Max(float64(s.sensorLocation[0]-(s.distance-distanceFromY)), 0))
				max := int(math.Min(float64(s.sensorLocation[0]+(s.distance-distanceFromY)), float64(maximumRow)))
				horizontal = append(horizontal, [2]int{min, max})
			}
		}
		sort.Slice(horizontal, func(i, j int) bool {
			return horizontal[i][0] < horizontal[j][0]
		})
		highestExtent := 0
		for h := 0; h < len(horizontal); h++ {
			if horizontal[h][0] > highestExtent+1 && horizontal[h][0] > 0 {
				return ((horizontal[h][0]-1)*4000000 + y)
			}
			if horizontal[h][1] >= highestExtent {
				highestExtent = horizontal[h][1]
			}
		}
		horizontal = nil
	}
	return 0
}

func prepareBeacons() []sensor {
	sensorList := []sensor{}
	var regex = regexp.MustCompile(`Sensor at x=(-?[0-9]*), y=(-?[0-9]*): closest beacon is at x=(-?[0-9]*), y=(-?[0-9]*)`)
	for _, s := range sensors {
		coords := regex.FindAllStringSubmatch(s, -1)
		sensorX, _ := strconv.Atoi(coords[0][1])
		sensorY, _ := strconv.Atoi(coords[0][2])
		beaconX, _ := strconv.Atoi(coords[0][3])
		beaconY, _ := strconv.Atoi(coords[0][4])

		distance := int(math.Abs(float64(sensorX-beaconX)) + math.Abs(float64(sensorY-beaconY)))

		sensorList = append(sensorList, sensor{[2]int{sensorX, sensorY}, [2]int{beaconX, beaconY}, distance})
	}
	return sensorList
}
