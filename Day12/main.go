package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func readTxtFile(input string) graph {
	file, _ := os.Open(input)
	lines, _ := ioutil.ReadAll(file)
	lineArray := strings.Split(string(lines), "\n")
	g := graph{make(map[string]node), make([]string, 0)}
	for _, l := range lineArray {
		if l == "" {
			continue
		}
		nodeArray := strings.Split(l, "-")
		g.addEdge(nodeArray[0], nodeArray[1])
		g.addEdge(nodeArray[1], nodeArray[0])
	}
	return g
}

type node struct {
	matches []string
}

func (n *node) addEdgeValue(v string) {
	n.matches = append(n.matches, v)
}

type graph struct {
	adjList  map[string]node
	allPaths []string
}

func (g *graph) addEdge(u string, v string) {
	//add v to u's list
	n := g.adjList[u]
	n.addEdgeValue(v)
	g.adjList[u] = n
}

func hasMoreThanOne(arr []string) bool {
	//Create a   dictionary of values for each element
	dict := make(map[string]int)
	for _, num := range arr {
		if string(num[0]) == strings.ToLower(string(num[0])) {
			dict[num] = dict[num] + 1
		}
	}
	count := 0
	for _, v := range dict {
		if v > 1 {
			count++
		}
		if count > 1 {
			return true
		}
	}

	return false
}

func (g *graph) part1(u string, d string, isVisited map[string]int, localPathList []string) {
	if u == d {
		g.allPaths = append(g.allPaths, strings.Join(localPathList, ","))
		return
	}

	if string(u[0]) != strings.ToUpper(string(u[0])) {
		isVisited[u] += 1
	}

	for _, i := range g.adjList[u].matches {
		if isVisited[i] < 1 {
			localPathList = append(localPathList, i)
			g.part1(i, d, isVisited, localPathList)

			localPathList = removeIndex(localPathList, i)
		}
	}
	isVisited[u] = 0
}

func (g *graph) part2(u string, d string, isVisited map[string]int, localPathList []string) {
	//I could find all paths that allow all nodes to be visited twice, then remove the extra paths
	if u == d {
		if !hasMoreThanOne(localPathList) {
			g.allPaths = append(g.allPaths, strings.Join(localPathList, ","))
		}
		return
	}

	if string(u[0]) != strings.ToUpper(string(u[0])) {
		isVisited[u] += 1
	}

	for _, i := range g.adjList[u].matches {
		if isVisited[i] < 1 || (i != "start" && i != "end" && isVisited[i] < 2) {
			localPathList = append(localPathList, i)
			g.part2(i, d, isVisited, localPathList)

			localPathList = removeIndex(localPathList, i)
		}
	}
	isVisited[u] -= 1

}

//convenience function to remove values from array
func removeIndex(s []string, v string) []string {
	for i, value := range s {
		if v == value {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s

}

func main() {
	g := readTxtFile("input.txt")
	s := "start"
	d := "end"

	isVisited := make(map[string]int)
	pathList := make([]string, 0)

	//add source to path[]
	pathList = append(pathList, s)

	//call recursive utility
	g.part1(s, d, isVisited, pathList)
	fmt.Printf("There are %d different paths from %v to %v\n", len(g.allPaths), s, d)

	g.allPaths = make([]string, 0)
	isVisited = make(map[string]int)
	pathList = make([]string, 0)

	//add source to path[]
	pathList = append(pathList, s)
	g.part2(s, d, isVisited, pathList)
	fmt.Printf("There are %d different paths from %v to %v for part 2\n", len(g.allPaths), s, d)
}
