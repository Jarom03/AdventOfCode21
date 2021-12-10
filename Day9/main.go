package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type position struct {
	y int
	x int
}

func readTxtFile(input string) [][]int {
	file, _ := os.Open(input)
	lines, _ := ioutil.ReadAll(file)
	lineArray := strings.Split(string(lines), "\n")
	heightMap := make([][]int, len(lineArray)-1)
	for i, l := range lineArray {
		if l == "" {
			continue
		}
		row := make([]int, 0)
		for _, r := range l {
			v, _ := strconv.Atoi(string(r))
			row = append(row, v)
		}
		heightMap[i] = row
	}
	return heightMap
}

func part1(heightMap [][]int) (int, []position) {
	sum := 0
	posArray := make([]position, 0)
	for y, _ := range heightMap {
		for x, _ := range heightMap[y] {
			if y > 0 {
				if heightMap[y][x] >= heightMap[y-1][x] {
					continue
				}
			}
			if y < len(heightMap)-1 {
				if heightMap[y][x] >= heightMap[y+1][x] {
					continue
				}
			}
			if x > 0 {
				if heightMap[y][x] >= heightMap[y][x-1] {
					continue
				}
			}
			if x < len(heightMap[y])-1 {
				if heightMap[y][x] >= heightMap[y][x+1] {
					continue
				}
			}
			sum += heightMap[y][x] + 1
			posArray = append(posArray, position{y, x})
		}
	}
	return sum, posArray
}

func calcBasin(heightMap [][]int, pos position, visited *[]position, size *int) {
	*visited = append(*visited, pos)
	*size++
	if pos.y > 0 && !hasVisitedPosition(position{pos.y - 1, pos.x}, *visited) {
		if heightMap[pos.y-1][pos.x] != 9 {
			calcBasin(heightMap, position{pos.y - 1, pos.x}, visited, size)
		}
	}
	if pos.y < len(heightMap)-1 && !hasVisitedPosition(position{pos.y + 1, pos.x}, *visited) {
		if heightMap[pos.y+1][pos.x] != 9 {
			calcBasin(heightMap, position{pos.y + 1, pos.x}, visited, size)
		}
	}
	if pos.x > 0 && !hasVisitedPosition(position{pos.y, pos.x - 1}, *visited) {
		if heightMap[pos.y][pos.x-1] != 9 {
			calcBasin(heightMap, position{pos.y, pos.x - 1}, visited, size)
		}
	}
	if pos.x < len(heightMap[pos.y])-1 && !hasVisitedPosition(position{pos.y, pos.x + 1}, *visited) {
		if heightMap[pos.y][pos.x+1] != 9 {
			calcBasin(heightMap, position{pos.y, pos.x + 1}, visited, size)
		}
	}
}

func hasVisitedPosition(pos position, visited []position) bool {
	for _, p := range visited {
		if p.x == pos.x && p.y == pos.y {
			return true
		}
	}
	return false
}

func part2(heightMap [][]int, lowestPositions []position) []int {
	largestBasins := make([]int, 0)
	for _, pos := range lowestPositions {
		currSize := 0
		visited := make([]position, 0)
		calcBasin(heightMap, pos, &visited, &currSize)
		// if len(largestBasins) < 3 {
		largestBasins = append(largestBasins, currSize)
	}
	sort.Ints(largestBasins)
	return largestBasins[len(largestBasins)-3:]
}

func main() {
	heightMap := readTxtFile("input.txt")
	result, posArray := part1(heightMap)
	fmt.Printf("The sum of risk levels is %d\n", result)
	largestBasins := part2(heightMap, posArray)
	fmt.Printf("the largest basins are %v and have a product of %d\n", largestBasins, largestBasins[0]*largestBasins[1]*largestBasins[2])
}
