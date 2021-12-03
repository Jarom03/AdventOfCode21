package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var indexCount int

func readTxtFile(input string) [][]int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	lines, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	lineArray := strings.Split(string(lines), "\n")
	binaryArray := make([][]int, len(lineArray))

	for i, l := range lineArray {
		if i == 0 {
			indexCount = len(l)
		}
		innerArray := make([]int, len(l))
		for n, c := range l {
			innerArray[n], err = strconv.Atoi(string(c))
			if err != nil {
				log.Fatal(err)
			}
		}
		binaryArray[i] = innerArray
	}
	return binaryArray
}

func part1(binaryArray [][]int) (int, int) {
	finalCount := make([]int, indexCount)
	for n, innerArray := range binaryArray {
		for i, value := range innerArray {
			//initialize
			if n == 0 {
				finalCount[i] = 0
			}
			if value >= 1 {
				finalCount[i] += 1
			} else {
				finalCount[i] -= 1
			}
		}
	}
	var gamma string
	var epsilon string
	for _, v := range finalCount {
		if v > 0 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"

		}
	}

	gammaValue, err := strconv.ParseInt(gamma, 2, 32) //convert from binary to decimal
	if err != nil {
		log.Fatal(err)
	}
	epsilonValue, err := strconv.ParseInt(epsilon, 2, 32) //convert from binary to decimal
	if err != nil {
		log.Fatal(err)
	}
	return int(gammaValue), int(epsilonValue)
}

func part2(binaryArray [][]int) (int, int) {
	oxyArray := make([][]int, len(binaryArray))
	copy(oxyArray, binaryArray) //copy since pass by ref
	for i, j := 0, len(oxyArray)-1; i < j; i, j = i+1, j-1 {
		oxyArray[i], oxyArray[j] = oxyArray[j], oxyArray[i]
	}
	c02Array := make([][]int, len(binaryArray))
	copy(c02Array, binaryArray) //copy since pass by ref
	for i, j := 0, len(c02Array)-1; i < j; i, j = i+1, j-1 {
		c02Array[i], c02Array[j] = c02Array[j], c02Array[i]
	}
	oxyOutput := findCommonValues(oxyArray, true)
	c02Output := findCommonValues(c02Array, false)

	oxyString := convertIntArrayToString(oxyOutput)
	c02String := convertIntArrayToString(c02Output)

	oxyValue, err := strconv.ParseInt(oxyString, 2, 32)
	if err != nil {
		log.Fatal(err)
	}
	c02Value, err := strconv.ParseInt(c02String, 2, 32)
	if err != nil {
		log.Fatal(err)
	}
	return int(oxyValue), int(c02Value)
}

//convert all the values of int array to a binary string
func convertIntArrayToString(intValue []int) string {
	strValue := fmt.Sprint(intValue)
	strValue = strings.ReplaceAll(strValue, " ", "")
	strValue = strings.ReplaceAll(strValue, "[", "")
	strValue = strings.ReplaceAll(strValue, "]", "")
	return strValue
}

func findCommonValues(inputArray [][]int, isMostCommon bool) []int {
	for currentIndex := 0; currentIndex < indexCount; currentIndex++ {
		bitCrit := findBitCriteria(inputArray, currentIndex, isMostCommon)
		for i := len(inputArray) - 1; i >= 0; i-- {

			if inputArray[i][currentIndex] != bitCrit {
				inputArray = removeIndex(inputArray, i)

				if len(inputArray) == 1 {
					return inputArray[0]
				}
			}
		}
	}
	return nil
}

func findBitCriteria(inputArray [][]int, currentIndex int, isMostCommon bool) int {
	value := 0
	for _, v := range inputArray {
		if v[currentIndex] == 1 {
			value += 1
		} else {
			value -= 1
		}
	}

	if !isMostCommon {
		value *= -1
		if value == 0 { //if it's most common and equal number return 0
			return 0
		}
	}

	if value >= 0 {
		return 1
	}

	return 0
}

//convenience function to remove values from array
func removeIndex(s [][]int, index int) [][]int {
	return append(s[:index], s[index+1:]...)
}

func main() {
	binaryArray := readTxtFile("input.txt")
	gamma, epsilon := part1(binaryArray)
	fmt.Printf("Gamma value = %d and epsilon value = %d for a product of %d\n", gamma, epsilon, gamma*epsilon)

	oxy, c02 := part2(binaryArray)
	fmt.Printf("Oxygen value = %d and C02 value = %d for a product of %d\n", oxy, c02, oxy*c02)
}
