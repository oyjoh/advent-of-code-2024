package util

import (
	"os"
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
