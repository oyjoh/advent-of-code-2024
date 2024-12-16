package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/oyjoh/advent-of-code-2024/util"
)

func main() {
	// path := "in_1.txt"
	// matrix, ins := readInp(path)

	// // util.PrintRuneMatrix(matrix)

	// fmt.Printf("Part 1: %d\n", pt1(ins, matrix))

	path := "in_ex.txt"
	matrix, ins := readInp(path)

	v2Matrix := v2(matrix)
	v2Matrix[2][7] = '['
	v2Matrix[2][8] = ']'
	v2Matrix[2][9] = '.'

	util.PrintRuneMatrix(v2Matrix)

	fmt.Printf("Part 2: %d\n", pt2(ins, v2Matrix))
}

func pt2(ins string, matrix [][]rune) int {
	y, x := getStart(matrix)

	for _, dir := range ins {
		//fmt.Printf("DIR: %c\n", dir)

		switch dir {
		case '<':
			y, x = move(y, x, 0, -1, matrix)
		case '^':
			y, x = move2(y, x, -1, 0, matrix)
		case '>':
			y, x = move(y, x, 0, 1, matrix)
		case 'v':
			y, x = move2(y, x, 1, 0, matrix)
		}
		util.PrintRuneMatrix(matrix)
		// fmt.Println()
	}
	return boxCoords(matrix)
}

// func check move?? for up/down only

func check(matrix [][]rune, y, x, dy, dx int) bool {
	fmt.Printf("inspecting y: %d, x: %d\n", y, x)
	//token := matrix[y][x]

	var x2 int
	if matrix[y+dy][x] == '[' {
		x2 = x + 1
	} else {
		x2 = x - 1
	}

	if matrix[y+dy][x] == '#' || matrix[y+dy][x2] == '#' {
		return false
	}
	if matrix[y+dy][x] == '.' || matrix[y+dy][x2] == '.' {
		return true
	}

	return check(matrix, y+dy, x, dy, dx) && check(matrix, y+dy, x2, dy, dx)
}

func move2(y, x, dy, dx int, matrix [][]rune) (int, int) {
	token := matrix[y][x]
	fmt.Printf("moving y: %d, x: %d\n", y, x)
	if matrix[y+dy][x+dx] == '#' {
		return y, x
	}
	if matrix[y+dy][x+dx] == '.' {
		matrix[y][x] = '.'
		matrix[y+dy][x+dx] = token
		return y + dy, x + dx
	}

	var x2 int
	if matrix[y+dy][x] == '[' {
		x2 = x + 1
	} else {
		x2 = x - 1
	}

	if check(matrix, y+dy, x, dy, dx) && check(matrix, y+dy, x2, dy, dx) {
		cy, cx := move(y+dy, x+dx, dy, dx, matrix)
		cy2, cx2 := move(y+dy, x2+dx, dy, dx, matrix)

		if cy != y || cx != x {
			if matrix[y+dy][x+dx] == '.' {
				matrix[y][x] = '.'
				matrix[y+dy][x+dx] = token
				return y + dy, x + dx
			}
		}

		if cy2 != y || cx2 != x2 {
			if matrix[y+dy][x2+dx] == '.' {
				matrix[y][x2] = '.'
				matrix[y+dy][x2+dx] = token
				return y + dy, x2 + dx
			}
		}
	}

	return y, x
}

func v2(matrix [][]rune) [][]rune {
	var v2 [][]rune

	for _, row := range matrix {
		var ins []rune
		for _, val := range row {
			switch val {
			case '#':
				ins = append(ins, '#', '#')
			case 'O':
				ins = append(ins, '[', ']')
			case '.':
				ins = append(ins, '.', '.')
			case '@':
				ins = append(ins, '@', '.')

			}
		}
		v2 = append(v2, ins)
	}

	return v2
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
