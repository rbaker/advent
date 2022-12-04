package main

import (
	"os"
	"strings"
)

var data3, _ = os.ReadFile("resources/day3.txt")
var itemList = strings.Split(strings.ReplaceAll(string(data3), "\r\n", "\n"), "\n")
var alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func day3part1() int {
	priority := 0

	for _, items := range itemList {
		comp1 := []byte(items[:len(items)/2])
		comp2 := []byte(items[len(items)/2:])

		m := make(map[byte]bool)
		for _, b := range comp1 {
			m[b] = true
		}
		for _, b := range comp2 {
			if m[b] {
				priority += strings.Index(alphabet, string(b)) + 1
				break
			}
		}
	}

	return priority
}

func day3part2() int {
	priority := 0

	m1 := make(map[byte]bool)
	m2 := make(map[byte]bool)

	for i, items := range itemList {
		item := []byte(items)

		for _, b := range item {

			if i%3 == 0 {
				m1[b] = true
			}

			if i%3 == 1 {
				m2[b] = true
			}

			if i%3 == 2 && m1[b] && m2[b] {
				priority += strings.Index(alphabet, string(b)) + 1
				m1 = make(map[byte]bool)
				m2 = make(map[byte]bool)
				break
			}

		}

	}

	return priority
}
