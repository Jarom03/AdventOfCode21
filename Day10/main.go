package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func readTxtFile(input string) []string {
	file, _ := os.Open(input)
	lines, _ := ioutil.ReadAll(file)
	lineArray := strings.Split(string(lines), "\n")
	if lineArray[len(lineArray)-1] == "" {
		return lineArray[0 : len(lineArray)-1]
	}
	return lineArray
}

func part1(lineArray []string) int {
	/*
		): 3 points.
		]: 57 points.
		}: 1197 points.
		>: 25137 points.
	*/
	var scoreMap = map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}
	totalScore := 0
	stack := Stack{}

	for _, line := range lineArray {
		stop := false
		for _, r := range line {
			if stop {
				break
			}
			switch string(r) {
			case "(":
				stack.Push("(")
			case "[":
				stack.Push("[")
			case "{":
				stack.Push("{")
			case "<":
				stack.Push("<")
			case ")":
				value, _ := stack.Pop()
				if value != "(" {
					fmt.Printf("Expected %v, but had )\n", value)
					totalScore += scoreMap[")"]
					stop = true
				}
			case "]":
				value, _ := stack.Pop()
				if value != "[" {
					fmt.Printf("Expected %v, but had ]\n", value)
					totalScore += scoreMap["]"]
					stop = true
				}
			case "}":
				value, _ := stack.Pop()
				if value != "{" {
					fmt.Printf("Expected %v, but had }\n", value)
					totalScore += scoreMap["}"]
					stop = true
				}
			case ">":
				value, _ := stack.Pop()
				if value != "<" {
					fmt.Printf("Expected %v, but had >\n", value)
					totalScore += scoreMap[">"]
					stop = true
				}
			}
		}
	}
	return totalScore
}

func part2(lineArray []string) []int {
	allScores := make([]int, 0)
	// var scoreMap = map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}
	stack := Stack{}

	for _, line := range lineArray {
		stop := false
		for _, r := range line {
			if stop {
				break
			}
			switch string(r) {
			case "(":
				stack.Push("(")
			case "[":
				stack.Push("[")
			case "{":
				stack.Push("{")
			case "<":
				stack.Push("<")
			case ")":
				value, _ := stack.Pop()
				if value != "(" {
					stack = make(Stack, 0)
					stop = true
				}
			case "]":
				value, _ := stack.Pop()
				if value != "[" {
					stack = make(Stack, 0)
					stop = true
				}
			case "}":
				value, _ := stack.Pop()
				if value != "{" {
					stack = make(Stack, 0)
					stop = true
				}
			case ">":
				value, _ := stack.Pop()
				if value != "<" {
					stack = make(Stack, 0)
					stop = true
				}
			}
		}
		if !stack.IsEmpty() {
			fmt.Printf("incomplete")
			totalScore := 0
			for !stack.IsEmpty() {
				v, _ := stack.Pop()
				switch v {
				case "(":
					//1
					totalScore = (totalScore * 5) + 1
				case "{":
					//3
					totalScore = (totalScore * 5) + 3
				case "[":
					//2
					totalScore = (totalScore * 5) + 2
				case "<":
					//4
					totalScore = (totalScore * 5) + 4
				}
			}
			allScores = append(allScores, totalScore)
			fmt.Printf("%v -- \n", stack)
		}
	}
	return allScores
}

func main() {
	lineArray := readTxtFile("input.txt")
	fmt.Printf("The length of the lineArray is %d\n", len(lineArray))
	score := part1(lineArray)
	fmt.Printf("The total score is %d\n", score)
	// lineArray2 := readTxtFile("testInput.txt")
	scoreArray := part2(lineArray)
	sort.Slice(scoreArray, func(i, j int) bool {
		return scoreArray[i] < scoreArray[j]
	})
	middleIndex := len(scoreArray) / 2
	fmt.Printf("The length of the scoreArray is %d the middle index is %d and the middle score for part 2 is %d\n", len(scoreArray), middleIndex, scoreArray[len(scoreArray)/2])

}
