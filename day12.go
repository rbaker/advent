package main

import (
	"math"
	"os"
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
	node     node
	distance int
}

var data12, _ = os.ReadFile("resources/day12.txt")
var mapLines = strings.Split(strings.ReplaceAll(string(data12), "\r\n", "\n"), "\n")
var g = createGraph()

func day12part1() int {
	return getDistance(g.start)
}
func day12part2() int {
	lowest := getDistance(g.start)
	for i, line := range mapLines {
		for j, letter := range line {
			if letter == 'a' {
				distance := getDistance(node{i, j})
				if distance > 0 {
					lowest = int(math.Min(float64(distance), float64(lowest)))
				}
			}
		}

	}
	return lowest
}

func getDistance(start node) int {
	g.nodes[start] = nodeInformation{start, 0}
	settled := []nodeInformation{}
	unsettled := make(map[node]nodeInformation)
	unsettled[start] = nodeInformation{start, 0}
	for len(unsettled) > 0 {
		currentNode := getLowest(unsettled)
		delete(unsettled, currentNode.node)

		for _, adjacent := range g.vertices[currentNode.node] {
			if !contains(settled, adjacent) {
				if currentNode.distance+1 < g.nodes[adjacent].distance {
					g.nodes[adjacent] = nodeInformation{adjacent, currentNode.distance + 1}
				}
				unsettled[adjacent] = g.nodes[adjacent]
			}
		}
		settled = append(settled, g.nodes[currentNode.node])
	}

	return g.nodes[g.end].distance
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
			g.nodes[node{i, j}] = nodeInformation{node{i, j}, int(9999999999999)}

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
func getLowest(nodes map[node]nodeInformation) nodeInformation {
	var n nodeInformation
	for _, v := range nodes {
		if v.distance <= n.distance || n.distance == 0 {
			n = v
		}
	}
	return n
}
