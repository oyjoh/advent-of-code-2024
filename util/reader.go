package util

import (
	"os"
	"strconv"
	"strings"
)

func RuneMatrix(path string) [][]rune {
	dat, _ := os.ReadFile(path)

	lines := strings.Split(string(dat), "\n")

	var matrix [][]rune

	for _, line := range lines {

		var row []rune
		for r := 0; r < len(line); r++ {
			row = append(row, rune(line[r]))
		}
		matrix = append(matrix, row)

	}
	return matrix
}

func ReadRuneMatrixFromLines(lines []string) [][]rune {
	var matrix [][]rune

	for _, line := range lines {

		var row []rune
		for r := 0; r < len(line); r++ {
			row = append(row, rune(line[r]))
		}
		matrix = append(matrix, row)

	}
	return matrix
}

func ReadIntMatrix(path string) [][]int {
	dat, _ := os.ReadFile(path)

	lines := strings.Split(string(dat), "\n")

	var matrix [][]int

	for _, line := range lines {

		var row []int
		for r := 0; r < len(line); r++ {
			num, _ := strconv.Atoi(string(line[r]))
			row = append(row, num)
		}
		matrix = append(matrix, row)

	}
	return matrix
}
