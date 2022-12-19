package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type graph struct {
	nodes    map[node]nodeInformation
	vertices map[node][]node
	start    node
	end      node
}
type node [2]int
type nodeInformation struct {
	node         node
	shortestPath []node
	distance     int
}

var data12, _ = os.ReadFile("resources/day12.txt")
var mapLines = strings.Split(strings.ReplaceAll(string(data12), "\r\n", "\n"), "\n")
var g = createGraph()

func day12part1() int {
	g.nodes[g.start] = nodeInformation{g.start, []node{}, 0}
	settled := []nodeInformation{}
	unsettled := []nodeInformation{{g.start, []node{}, 0}}
	fmt.Println(g)
	for len(unsettled) > 0 {

		sort.Slice(unsettled, func(i, j int) bool {
			return unsettled[i].distance < unsettled[j].distance
		})
		currentNode := unsettled[0]

		for _, adjacent := range g.vertices[currentNode.node] {
			if !contains(settled, adjacent) {
				if currentNode.distance+1 < g.nodes[adjacent].distance {
					currentNode.shortestPath = append(currentNode.shortestPath, currentNode.node)
					g.nodes[adjacent] = nodeInformation{adjacent, currentNode.shortestPath, currentNode.distance + 1}
				}
				unsettled = append(unsettled, g.nodes[adjacent])
			}
		}
		settled = append(settled, g.nodes[currentNode.node])
		copy(unsettled, unsettled[1:])

	}
	fmt.Println(g.nodes[g.end])

	return 0
}

func createGraph() graph {
	g := graph{}

	for i, line := range mapLines {
		if strings.Contains(line, "S") {
			g.start = node{i, strings.Index(line, "S")}
			mapLines[i] = strings.ReplaceAll(mapLines[i], "S", "a")
		}
		if strings.Contains(line, "E") {
			g.end = node{i, strings.Index(line, "E")}
			mapLines[i] = strings.ReplaceAll(mapLines[i], "E", "z")
		}
	}

	g.nodes = make(map[node]nodeInformation)
	g.vertices = make(map[node][]node)
	for i, line := range mapLines {
		for j, letter := range line {
			g.nodes[node{i, j}] = nodeInformation{node{i, j}, []node{}, 9223372036854775807}

			if i > 0 && mapLines[i-1][j] <= byte(letter)+1 {
				g.vertices[node{i, j}] = append(g.vertices[node{i, j}], node{i - 1, j})
			}
			if i < len(mapLines)-1 && mapLines[i+1][j] <= byte(letter)+1 {
				g.vertices[node{i, j}] = append(g.vertices[node{i, j}], node{i + 1, j})
			}
			if j > 0 && mapLines[i][j-1] <= byte(letter)+1 {
				g.vertices[node{i, j}] = append(g.vertices[node{i, j}], node{i, j - 1})
			}
			if j < len(line)-1 && mapLines[i][j+1] <= byte(letter)+1 {
				g.vertices[node{i, j}] = append(g.vertices[node{i, j}], node{i, j + 1})
			}
		}
	}
	return g
}
func contains(s []nodeInformation, e node) bool {
	for _, a := range s {
		if a.node == e {
			return true
		}
	}
	return false
}
