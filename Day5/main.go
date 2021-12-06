package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type lineFragment struct {
	point1 point
	point2 point
}

func readTxtFile(input string) []lineFragment {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	lines, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	lineArray := strings.Split(string(lines), "\n")
	lineFragmentArray := make([]lineFragment, 0)

	for _, l := range lineArray {
		if l == "" {
			continue
		}
		lineSplit := strings.Split(l, " -> ")
		point1String := strings.Split(lineSplit[0], ",")
		point2String := strings.Split(lineSplit[1], ",")
		x1, _ := strconv.Atoi(point1String[0])
		y1, _ := strconv.Atoi(point1String[1])
		x2, _ := strconv.Atoi(point2String[0])
		y2, _ := strconv.Atoi(point2String[1])
		lineFragmentArray = append(lineFragmentArray, lineFragment{point{x1, y1}, point{x2, y2}})
	}
	return lineFragmentArray
}

func part1(lineFrags []lineFragment, diag bool) int {
	pointCrossed := make(map[string]int)
	for _, frag := range lineFrags {
		//only care if it's horizontal or vertical
		if frag.point1.x == frag.point2.x {
			if frag.point1.y < frag.point2.y {
				for i := frag.point1.y; i <= frag.point2.y; i++ {
					key := strconv.Itoa(frag.point1.x) + "," + strconv.Itoa(i)
					pointCrossed[key] += 1
				}
			} else {
				for i := frag.point2.y; i <= frag.point1.y; i++ {
					key := strconv.Itoa(frag.point1.x) + "," + strconv.Itoa(i)
					pointCrossed[key] += 1
				}
			}
		} else if frag.point1.y == frag.point2.y {
			if frag.point1.x < frag.point2.x {
				for i := frag.point1.x; i <= frag.point2.x; i++ {
					key := strconv.Itoa(i) + "," + strconv.Itoa(frag.point1.y)
					pointCrossed[key] += 1
				}
			} else {
				for i := frag.point2.x; i <= frag.point1.x; i++ {
					key := strconv.Itoa(i) + "," + strconv.Itoa(frag.point1.y)
					pointCrossed[key] += 1
				}
			}
		} else if diag {
			if frag.point1.x < frag.point2.x && frag.point1.y > frag.point2.y {
				for i, n := frag.point1.x, frag.point1.y; i <= frag.point2.x; i, n = i+1, n-1 {
					key := strconv.Itoa(i) + "," + strconv.Itoa(n)
					pointCrossed[key] += 1
				}
			} else if frag.point1.x > frag.point2.x && frag.point1.y < frag.point2.y {
				for i, n := frag.point2.x, frag.point2.y; i <= frag.point1.x; i, n = i+1, n-1 {
					key := strconv.Itoa(i) + "," + strconv.Itoa(n)
					pointCrossed[key] += 1
				}
			} else if frag.point1.x < frag.point2.x && frag.point1.y < frag.point2.y {
				for i, n := frag.point1.x, frag.point1.y; i <= frag.point2.x; i, n = i+1, n+1 {
					key := strconv.Itoa(i) + "," + strconv.Itoa(n)
					pointCrossed[key] += 1
				}
			} else {
				for i, n := frag.point2.x, frag.point2.y; i <= frag.point1.x; i, n = i+1, n+1 {
					key := strconv.Itoa(i) + "," + strconv.Itoa(n)
					pointCrossed[key] += 1
				}
			}
		}
	}
	num2s := 0
	for _, v := range pointCrossed {
		if v >= 2 {
			num2s++
		}
	}
	return num2s
}

func main() {
	lineFrags := readTxtFile("input.txt")
	result := part1(lineFrags, false)
	fmt.Printf("%d points are crossed by two horizontal or vertical lines\n", result)
	result2 := part1(lineFrags, true)
	fmt.Printf("%d points are crossed by two horizontal, vertical, or diagnoal lines\n", result2)
}
