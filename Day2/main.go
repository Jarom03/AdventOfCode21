package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type movement struct {
	direction string
	distance  int
}

func readTxtFile(input string) []movement {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	lines, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	lineArray := strings.Split(string(lines), "\n")
	moveArray := make([]movement, len(lineArray))

	for i, l := range lineArray {
		lineSplit := strings.Split(l, " ")
		direction := lineSplit[0]
		distance, err := strconv.Atoi(lineSplit[1])
		if err != nil {
			log.Fatal(err)
		}
		moveArray[i] = movement{direction, distance}
	}
	return moveArray
}

func part1(moveArray []movement) (position int, depth int) {
	position = 0
	depth = 0
	for _, move := range moveArray {
		switch mv := move.direction; mv {
		case "up":
			depth -= move.distance
		case "down":
			depth += move.distance
		case "forward":
			position += move.distance
		}
	}
	return position, depth
}

func part2(moveArray []movement) (aim int, position int, depth int) {
	position = 0
	depth = 0
	aim = 0
	for _, move := range moveArray {
		switch mv := move.direction; mv {
		case "up":
			aim -= move.distance
		case "down":
			aim += move.distance
		case "forward":
			position += move.distance
			depth += aim * move.distance
		}
	}
	return aim, position, depth
}

func main() {
	moveArray := readTxtFile("input.txt")
	position, depth := part1(moveArray)
	fmt.Printf("The horizontal position is %d and the depth is %d with a product of %d\n", position, depth, position*depth)

	aim, position, depth := part2(moveArray)
	fmt.Printf("The horizontal position is %d and the depth is %d with a product of %d and an end aim of %d\n", position, depth, position*depth, aim)
}
