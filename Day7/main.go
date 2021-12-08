package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func readTxtFile(input string) ([]int, int, int) {
	file, _ := os.Open(input)
	min, max := 9999999, 0
	lines, _ := ioutil.ReadAll(file)
	lineArray := strings.Split(string(lines), "\n")
	numString := strings.Split(lineArray[0], ",")
	values := make([]int, len(numString))
	for i, s := range numString {
		values[i], _ = strconv.Atoi(s)
		if values[i] > max {
			max = values[i]
		}
		if values[i] < min {
			min = values[i]
		}
	}
	return values, min, max
}

func part1(values []int, min int, max int) (int, int) {
	bestPosition, lowestFuel := -1, 99999999999999
	for pos := min; pos < max; pos++ {
		fuelBurned := 0
		for _, v := range values {
			fuelBurned += int(math.Abs(float64(v - pos)))
		}
		if fuelBurned < lowestFuel {
			lowestFuel = fuelBurned
			bestPosition = pos
		}
	}
	return bestPosition, lowestFuel
}

func part2(values []int, min int, max int) (int, int) {
	bestPosition, lowestFuel := -1, 99999999999999999
	for pos := min; pos < max; pos++ {
		fuelBurned := 0
		for _, v := range values {
			fuelBurned += calculateSumOfSeries(int(math.Abs(float64(v - pos))))
		}
		if fuelBurned < lowestFuel {
			lowestFuel = fuelBurned
			bestPosition = pos
		}
	}
	return bestPosition, lowestFuel
}

func calculateSumOfSeries(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum
}

func main() {
	values, min, max := readTxtFile("input.txt")
	fmt.Printf("There are %d submarines with the lowest position being %d and highest position is %d\n", len(values), min, max)
	bestPos, fuel := part1(values, min, max)
	fmt.Printf("%d is the best horizontal position using only %d fuel\n", bestPos, fuel)
	bestPos, fuel = part2(values, min, max)
	fmt.Printf("For part 2: %d is the best horizontal position using only %d fuel\n", bestPos, fuel)
}
