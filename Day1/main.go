package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func readTxtFile(input string) []string {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	lines, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	lineArray := strings.Split(string(lines), "\n")
	return lineArray
}

func part1(lineArray []string) {
	increased := 0
	prevValue := -1

	for i, l := range lineArray {
		if i != 0 {
			currentValue, _ := strconv.Atoi(l)
			if currentValue > prevValue {
				increased += 1
			}
		}
		prevValue, _ = strconv.Atoi(l)
	}
	fmt.Printf("the number of increases were %d\n", increased)
}

func part2(lineArray []string) {
	increased := 0
	prevSum := 0

	for i, l := range lineArray {
		if i > 1 {
			prevPrevValue, _ := strconv.Atoi(lineArray[i-2])
			prevValue, _ := strconv.Atoi(lineArray[i-1])
			currentValue, _ := strconv.Atoi(l)
			curSum := currentValue + prevValue + prevPrevValue

			if i > 2 && curSum > prevSum {
				increased += 1
			}
			prevSum = curSum
		}

	}
	fmt.Printf("the number of increases were %d\n", increased)
}

func main() {
	lineArray := readTxtFile("input.txt")
	part1(lineArray)
	part2(lineArray)
}
