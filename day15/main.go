package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/oyjoh/advent-of-code-2024/util"
)

func main() {
	path := "in_1.txt"
	matrix, ins := readInp(path)

	// util.PrintRuneMatrix(matrix)

	fmt.Printf("Part 1: %d\n", pt1(ins, matrix))
}

func pt1(ins string, matrix [][]rune) int {
	y, x := getStart(matrix)

	for _, dir := range ins {
		//fmt.Printf("DIR: %c\n", dir)

		switch dir {
		case '<':
			y, x = move(y, x, 0, -1, matrix)
		case '^':
			y, x = move(y, x, -1, 0, matrix)
		case '>':
			y, x = move(y, x, 0, 1, matrix)
		case 'v':
			y, x = move(y, x, 1, 0, matrix)
		}
		// util.PrintRuneMatrix(matrix)
		// fmt.Println()
	}
	return boxCoords(matrix)
}

func boxCoords(matrix [][]rune) int {
	sum := 0

	for y, row := range matrix {
		for x, val := range row {
			if val == 'O' {
				sum += 100*y + x
			}
		}
	}

	return sum
}

// func check move?? for up/down only

func move(y, x, dy, dx int, matrix [][]rune) (int, int) {
	token := matrix[y][x]
	if matrix[y+dy][x+dx] == '#' {
		return y, x
	}
	if matrix[y+dy][x+dx] == '.' {
		matrix[y][x] = '.'
		matrix[y+dy][x+dx] = token
		return y + dy, x + dx
	}

	cy, cx := move(y+dy, x+dx, dy, dx, matrix)

	if cy != y || cx != x {
		if matrix[y+dy][x+dx] == '.' {
			matrix[y][x] = '.'
			matrix[y+dy][x+dx] = token
			return y + dy, x + dx
		}
	}

	return y, x
}

func getStart(matrix [][]rune) (int, int) {
	for y, row := range matrix {
		for x, val := range row {
			if val == '@' {
				return y, x
			}
		}
	}
	return 0, 0
}

func readInp(path string) ([][]rune, string) {
	dat, _ := os.ReadFile(path)

	ab := strings.Split(string(dat), "\n\n")
	lines := strings.Split(ab[0], "\n")

	matrix := util.ReadRuneMatrixFromLines(lines)

	return matrix, ab[1]
}
