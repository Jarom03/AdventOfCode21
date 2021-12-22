package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type foldInstructions struct {
	x int
	y int
}

type transparentOragami struct {
	grid         [][]rune
	instructions []foldInstructions
}

func addField(grid [][]rune, x int, y int) [][]rune {
	if y > len(grid) {
		grid = append(grid, make([][]rune, y-len(grid)+1)...)
	}
	if len(grid[y]) == 0 || x > len(grid[y]) {
		grid[y] = append(grid[y], make([]rune, x-len(grid[y])+1)...)
	}
	grid[y][x] = '#'
	return grid
}

func readTxtFile(input string) transparentOragami {
	file, _ := os.Open(input)
	lines, _ := ioutil.ReadAll(file)
	lineArray := strings.Split(string(lines), "\n")
	grid := make([][]rune, 0)
	instructions := make([]foldInstructions, 0)
	for _, l := range lineArray {
		if l == "" {
			continue
		}
		if strings.Contains(l, "fold") {
			inst := foldInstructions{}
			if strings.Contains(l, "x=") {
				inst.x, _ = strconv.Atoi(strings.Split(l, "x=")[1])
			} else {
				inst.y, _ = strconv.Atoi(strings.Split(l, "y=")[1])
			}
			instructions = append(instructions, inst)
			continue
		}
		coords := strings.Split(l, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		grid = addField(grid, x, y)

	}
	origami := transparentOragami{grid, instructions}
	return origami
}

func foldPaper(grid [][]rune, instruction foldInstructions) [][]rune {
	for y := range grid {
		for x := range grid[y] {
			if instruction.x != 0 {
				//every point that is > than x needs to move
				if x > instruction.x {
					if grid[y][x] == '#' {
						moveDistance := x - instruction.x
						grid = addField(grid, instruction.x-moveDistance, y)
						grid[y][x] = '.'
					}
				}
			}
			if instruction.y != 0 {
				if y > instruction.y {
					if grid[y][x] == '#' {
						moveDistance := y - instruction.y
						grid = addField(grid, x, instruction.y-moveDistance)
						grid[y][x] = '.'
					}
				}
			}
		}
	}
	return grid
}

func part1(origami transparentOragami) int {
	//complete only the first fold
	grid := foldPaper(origami.grid, origami.instructions[0])
	dotCount := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == '#' {
				dotCount++
			}
		}
	}
	return dotCount
}

func part2(origami transparentOragami) {
	grid := origami.grid
	for _, fold := range origami.instructions {
		grid = foldPaper(grid, fold)
	}
	for _, row := range grid {
		for _, value := range row {
			if value == '#' {
				fmt.Printf("# ")
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	origami := readTxtFile("input.txt")
	dotCount := part1(origami)
	fmt.Printf("The number of visible dots after first instruction is %d\n", dotCount)

	part2(origami)
}
