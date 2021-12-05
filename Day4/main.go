package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type square struct {
	value   int
	isMatch bool
}

type bingoBoard struct {
	board         [][]square
	winningNumber int
	isWin         bool
}

func (bb bingoBoard) computeScore() int {
	score := 0
	for _, row := range bb.board {
		for _, square := range row {
			if !square.isMatch {
				score += square.value
			}
		}
	}
	return score * bb.winningNumber
}

func (bb *bingoBoard) checkForWin(lastNumber int) bool {
	columnMatch := make([]int, 5)
	for _, row := range bb.board {
		rowMatch := 0
		for i, v := range row {
			if v.isMatch {
				columnMatch[i]++
				rowMatch++
			}
		}
		if rowMatch == 5 {
			bb.winningNumber = lastNumber
			bb.isWin = true
			return true
		}
	}
	for _, m := range columnMatch {
		if m == 5 {
			bb.winningNumber = lastNumber
			bb.isWin = true
			return true
		}
	}
	return false
}

func (bb *bingoBoard) checkNumber(number int) bool {
	for i, row := range bb.board {
		for n, v := range row {
			if number == v.value {
				bb.board[i][n].isMatch = true
				return true
			}
		}
	}
	return false
}

func readTxtFile(input string) ([]bingoBoard, []int) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	lines, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	lineArray := strings.Split(string(lines), "\n")
	var bingoArray []bingoBoard
	var numbers []int
	inPuzzle := false

	var board [][]square
	for i, l := range lineArray {
		if i == 0 {
			stringArray := strings.Split(l, ",")
			for _, s := range stringArray {
				num, _ := strconv.Atoi(s)
				numbers = append(numbers, num)
			}
		} else {
			//a bingo board is 5 rows long and a new puzzle starts every 6 rows (empty row between)
			if l == "" {
				if inPuzzle {
					inPuzzle = false
					bingoArray = append(bingoArray, bingoBoard{board, 0, false})
					board = nil
				}
			} else {
				if !inPuzzle {
					inPuzzle = true
					board = make([][]square, 0)
				}
				row := make([]square, 0)
				rowString := strings.Split(l, " ")
				for _, s := range rowString {
					if s != "" {
						v, _ := strconv.Atoi(s)
						row = append(row, square{v, false})
					}
				}
				board = append(board, row)
			}
		}
	}
	return bingoArray, numbers
}

func part1(board []bingoBoard, numbers []int) bingoBoard {
	for _, num := range numbers {
		for _, g := range board {
			if g.checkNumber(num) {
				if g.checkForWin(num) {
					return g
				}
			}
		}
	}
	return bingoBoard{}
}

func part2(board []bingoBoard, numbers []int) bingoBoard {
	countWinningBoards := 0
	for _, num := range numbers {
		for i, _ := range board {
			if board[i].checkNumber(num) {
				if !board[i].isWin && board[i].checkForWin(num) {
					countWinningBoards++
					if countWinningBoards == len(board) {
						return board[i]
					}
				}
			}
		}
	}
	return bingoBoard{}
}

func main() {
	board, numbers := readTxtFile("input.txt")
	board2, numbers := readTxtFile("input.txt") //lazy...
	result := part1(board, numbers)
	fmt.Printf("The score of the winning board is %d\n", result.computeScore())
	result2 := part2(board2, numbers)
	fmt.Printf("The score of the last winning board is %v\n", result2.computeScore())
}
