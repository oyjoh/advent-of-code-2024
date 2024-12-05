package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	matrix := filetomatrix("in_ex.txt")

	fmt.Printf("example part 1: %d\n", solvePart1(matrix))
	fmt.Printf("example part 2: %d\n", solvePart2(matrix))

	fmt.Printf("part1: %d\n", solvePart1(filetomatrix("in_1.txt")))
	fmt.Printf("part2: %d\n", solvePart2(filetomatrix("in_1.txt")))
}

func solvePart1(matrix [][]int) int {
	safeReports := 0

	for _, report := range matrix {
		safeReports += checkReport(report)
	}

	return safeReports
}

func solvePart2(matrix [][]int) int {
	safeReports := 0

	for _, report := range matrix {
		safeReports += checkReportDamp(report)
	}

	return safeReports
}

func checkReportDamp(report []int) int {
	for i := -1; i < len(report); i++ {
		res := checkReport(remove(report, i))
		if res == 1 {
			return 1
		}
	}
	return 0
}

func remove(slice []int, s int) []int {
	if s == -1 {
		return slice
	}
	cop := make([]int, len(slice))
	copy(cop, slice)

	newSlice := append(cop[:s], cop[s+1:]...)

	return newSlice
}

func checkReport(report []int) int {
	inc := true
	if report[1] < report[0] {
		inc = false
	}

	prev := report[0]
	for i := 1; i < len(report); i++ {

		if report[i] == prev {
			return 0
		}

		diff := report[i] - prev

		switch diff {
		case 1, 2, 3:
			if !inc {
				return 0
			}
		case -1, -2, -3:
			if inc {
				return 0
			}
		default:
			return 0
		}

		prev = report[i]
	}

	return 1

}

func filetomatrix(filename string) [][]int {
	dat, err := os.ReadFile(filename)
	check(err)

	lines := strings.Split(string(dat), "\n")

	var matrix [][]int

	for _, line := range lines {
		fields := strings.Fields(line)

		var row []int
		for _, field := range fields {
			num, _ := strconv.Atoi(field)
			row = append(row, num)
		}

		matrix = append(matrix, row)
	}

	return matrix
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
