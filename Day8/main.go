package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type segmentGroup struct {
	zeroTen []string
	output  []string
}

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortStringArray(stringArray []string) []string {
	for i, s := range stringArray {
		r := StringToRuneSlice(s)
		sort.Slice(r, func(i, j int) bool {
			return r[i] < r[j]
		})
		stringArray[i] = string(r)
	}
	return stringArray
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func (sg segmentGroup) getValue() int {
	result := ""
	for _, s := range sg.output {
		for i, actual := range sg.zeroTen {
			if actual == s {
				result += strconv.Itoa(i)
			}
		}
	}
	value, _ := strconv.Atoi(result)
	return value
}

func (sg *segmentGroup) matchSegments() {
	one, two, three, four, five, six, seven, eight, nine, zero := "", "", "", "", "", "", "", "", "", ""
	fiveSegments := make([]string, 0) //2, 3, 5
	sixSegments := make([]string, 0)  //0, 6, 9

	for _, s := range sg.zeroTen {
		switch len(s) {
		case 2:
			one = s
		case 3:
			seven = s
		case 4:
			four = s
		case 5:
			fiveSegments = append(fiveSegments, s)
		case 6:
			sixSegments = append(sixSegments, s)
		case 7:
			eight = s
		}
	}

	//6 is the only 6 segment that doesn't have both "1 segments"
	removeIndex := -1
	for i, s := range sixSegments {
		if !strings.Contains(s, string(one[0])) || !strings.Contains(s, string(one[1])) {
			six = s
			removeIndex = i
			break
		}
	}
	remove(sixSegments, removeIndex)
	//0 is the remaining 6 segment that doesn't have all the "4 segments"
	for i, s := range sixSegments {
		if !strings.Contains(s, string(four[0])) || !strings.Contains(s, string(four[1])) || !strings.Contains(s, string(four[2])) || !strings.Contains(s, string(four[3])) {
			zero = s
			removeIndex = i
			break
		}
	}
	remove(sixSegments, removeIndex)
	//9 is the remaining 6 segment
	nine = sixSegments[0]
	//3 is the 5 segment that contains both "1 segments"
	for i, s := range fiveSegments {
		if strings.Contains(s, string(one[0])) && strings.Contains(s, string(one[1])) {
			three = s
			removeIndex = i
			break
		}
	}
	remove(fiveSegments, removeIndex)
	//the top segment in the 1 is the segment that is in one but not in 6
	topSegment := string(one[0])
	if strings.Contains(six, string(one[0])) {
		topSegment = string(one[1])
	}
	//5 is the 5 segment that doesn't contain the top segment from the "1 segments"
	for i, s := range fiveSegments {
		if !strings.Contains(s, topSegment) {
			five = s
			removeIndex = i
			break
		}
	}
	remove(fiveSegments, removeIndex)
	//2 is the remaining 5 segment
	two = fiveSegments[0]
	sg.zeroTen[0] = zero
	sg.zeroTen[1] = one
	sg.zeroTen[2] = two
	sg.zeroTen[3] = three
	sg.zeroTen[4] = four
	sg.zeroTen[5] = five
	sg.zeroTen[6] = six
	sg.zeroTen[7] = seven
	sg.zeroTen[8] = eight
	sg.zeroTen[9] = nine

}

func readTxtFile(input string) []segmentGroup {
	file, _ := os.Open(input)
	sgArray := make([]segmentGroup, 0)
	lines, _ := ioutil.ReadAll(file)
	lineArray := strings.Split(string(lines), "\n")
	for _, l := range lineArray {
		if l == "" {
			continue
		}
		twoSet := strings.Split(l, " | ")
		zeroTen := strings.Split(twoSet[0], " ")
		output := strings.Split(twoSet[1], " ")
		sgArray = append(sgArray, segmentGroup{SortStringArray(zeroTen), SortStringArray(output)})
	}
	return sgArray
}

func part1(sgArray []segmentGroup) int {
	uniqueCount := 0
	for _, sg := range sgArray {
		for _, seg := range sg.output {
			segLen := len(seg)
			if segLen == 2 || segLen == 4 || segLen == 3 || segLen == 7 {
				uniqueCount++
			}
		}
	}
	return uniqueCount
}

func part2(sgArray []segmentGroup) int {
	total := 0
	for _, sg := range sgArray {
		sg.matchSegments()
		total += sg.getValue()
	}
	return total
}

func main() {
	sgArray := readTxtFile("input.txt")
	result := part1(sgArray)
	fmt.Printf("There are %d unique (1, 4, 7, 8) segments\n", result)
	result = part2(sgArray)
	fmt.Printf("The total value is %d\n", result)
}
