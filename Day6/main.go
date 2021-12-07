package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func readTxtFile(input string) []int {
	file, _ := os.Open(input)
	lines, _ := ioutil.ReadAll(file)
	lineArray := strings.Split(string(lines), "\n")
	numString := strings.Split(lineArray[0], ",")
	values := make([]int, len(numString))
	for i, s := range numString {
		values[i], _ = strconv.Atoi(s)
	}
	return values
}

func part1(timerArray []int, numDays int) []int {
	for i := 0; i < numDays; i++ {
		arraySize := len(timerArray)
		for n := 0; n < arraySize; n++ {
			if timerArray[n] == 0 {
				timerArray[n] = 6
				timerArray = append(timerArray, 8)
			} else {
				timerArray[n] -= 1
			}
		}
		// fmt.Printf("%v\n", timerArray)
	}
	return timerArray
}

func part2(timerArray []int, numDays int) int {
	timerMap := make(map[int]int)
	newTimerMap := make(map[int]int)
	for i, _ := range timerArray {
		timerMap[timerArray[i]] += 1
	}

	for d := 0; d < numDays; d++ {
		newTimer := timerMap[d%7]
		timerMap[d%7] += newTimerMap[d%9]
		newTimerMap[d%9] += newTimer
	}

	total := 0
	for _, v := range timerMap {
		total += v
	}
	for _, v := range newTimerMap {
		total += v
	}

	return total
}

func main() {
	timerArray := readTxtFile("input.txt")
	timerArray2 := readTxtFile("input.txt") //lazy
	result := part1(timerArray, 80)
	fmt.Printf("After 80 days there will be %d fish\n", len(result))
	result2 := part2(timerArray2, 256)
	fmt.Printf("After 256 days there will be %d fish\n", result2)
}
