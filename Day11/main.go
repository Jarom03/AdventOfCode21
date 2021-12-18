package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type field struct {
	y int
	x int
}

func readTxtFile(input string) [][]int {
	file, _ := os.Open(input)
	lines, _ := ioutil.ReadAll(file)
	lineArray := strings.Split(string(lines), "\n")
	energyTable := make([][]int, 0)
	for _, l := range lineArray {
		if l == "" {
			continue
		}
		row := make([]int, 0)
		for _, v := range l {
			value, _ := strconv.Atoi(string(v))
			row = append(row, value)
		}
		energyTable = append(energyTable, row)
	}
	return energyTable
}

func incrementAdjacent(energyTable [][]int, pos field) {
	if pos.y > 0 {
		energyTable[pos.y-1][pos.x]++
		if energyTable[pos.y-1][pos.x] == 10 {
			incrementAdjacent(energyTable, field{pos.y - 1, pos.x})
		}
		if pos.x > 0 {
			energyTable[pos.y-1][pos.x-1]++
			if energyTable[pos.y-1][pos.x-1] == 10 {
				incrementAdjacent(energyTable, field{pos.y - 1, pos.x - 1})
			}
		}
		if pos.x < len(energyTable[pos.y])-1 {
			energyTable[pos.y-1][pos.x+1]++
			if energyTable[pos.y-1][pos.x+1] == 10 {
				incrementAdjacent(energyTable, field{pos.y - 1, pos.x + 1})
			}
		}
	}

	if pos.y < len(energyTable)-1 {
		energyTable[pos.y+1][pos.x]++
		if energyTable[pos.y+1][pos.x] == 10 {
			incrementAdjacent(energyTable, field{pos.y + 1, pos.x})
		}
		if pos.x > 0 {
			energyTable[pos.y+1][pos.x-1]++
			if energyTable[pos.y+1][pos.x-1] == 10 {
				incrementAdjacent(energyTable, field{pos.y + 1, pos.x - 1})
			}
		}
		if pos.x < len(energyTable[pos.y])-1 {
			energyTable[pos.y+1][pos.x+1]++
			if energyTable[pos.y+1][pos.x+1] == 10 {
				incrementAdjacent(energyTable, field{pos.y + 1, pos.x + 1})
			}
		}
	}

	if pos.x > 0 {
		energyTable[pos.y][pos.x-1]++
		if energyTable[pos.y][pos.x-1] == 10 {
			incrementAdjacent(energyTable, field{pos.y, pos.x - 1})
		}
	}

	if pos.x < len(energyTable[pos.y])-1 {
		energyTable[pos.y][pos.x+1]++
		if energyTable[pos.y][pos.x+1] == 10 {
			incrementAdjacent(energyTable, field{pos.y, pos.x + 1})
		}
	}

}

func part1(energyTable [][]int, intervals int) int {
	sumFlashes := 0
	for w := 0; w < intervals; w++ {
		for i := range energyTable {
			for n := range energyTable[i] {
				energyTable[i][n]++
				if energyTable[i][n] == 10 {
					incrementAdjacent(energyTable, field{i, n})
				}
			}
		}

		for i, _ := range energyTable {
			for n, v := range energyTable[i] {
				if v > 9 {
					sumFlashes++
					energyTable[i][n] = 0
				}
			}
		}
	}
	return sumFlashes
}

func part2(energyTable [][]int) int {
	count := 0
	intervalFlashes := 0
	totalOcto := len(energyTable) * len(energyTable[0])
	for intervalFlashes != totalOcto {
		for i := range energyTable {
			for n := range energyTable[i] {
				energyTable[i][n]++
				if energyTable[i][n] == 10 {
					incrementAdjacent(energyTable, field{i, n})
				}
			}
		}
		intervalFlashes = 0
		for i, _ := range energyTable {
			for n, v := range energyTable[i] {
				if v > 9 {
					intervalFlashes++
					energyTable[i][n] = 0
				}
			}
		}
		count++
	}
	return count

}

func main() {
	energyTable := readTxtFile("input.txt")
	sumFlashes := part1(energyTable, 100)
	fmt.Printf("The number of flashes after 100 intervales = %d\n", sumFlashes)

	energyTable = readTxtFile("input.txt")
	step := part2(energyTable)
	fmt.Printf("The first step when all flash together is step %d\n", step)
}
