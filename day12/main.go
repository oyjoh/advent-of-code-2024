package main

import (
	"fmt"
	"slices"

	"github.com/oyjoh/advent-of-code-2024/util"
)

func main() {
	path := "in_1.txt"
	matrix := util.RuneMatrix(path)

	fmt.Printf("Part 1: %d\n", pt1(matrix))
	fmt.Printf("Part 2: %d\n", pt2(matrix))
}

func pt2(matrix [][]rune) int {
	visited := make(map[string]bool)
	price := 0

	for y, row := range matrix {
		for x := range row {
			key := fmt.Sprintf("%d,%d", y, x)

			if !visited[key] {
				sides := make(map[string][]int)
				a := discoverE(sides, visited, matrix, y, x, 1)
				price += a * countSides(sides)
			}

		}
	}

	return price
}

func countSides(sides map[string][]int) int {
	sideCnt := 0
	for key := range sides {
		ar := sides[key]
		slices.Sort(ar)

		sideCnt++

		for i := 1; i < len(ar); i++ {
			if ar[i]-ar[i-1] != 1 {
				sideCnt++
			}
		}
	}

	return sideCnt
}

func discoverE(sides map[string][]int, visited map[string]bool, matrix [][]rune, y, x int, area int) int {
	key := fmt.Sprintf("%d,%d", y, x)
	visited[key] = true
	val := matrix[y][x]

	switch check(visited, matrix, y+1, x, val) {
	case -1:
		key := fmt.Sprintf("u%d", y)
		sides[key] = append(sides[key], x)
	case 0:
		key := fmt.Sprintf("u%d", y)
		sides[key] = append(sides[key], x)
	case 1:
		area = discoverE(sides, visited, matrix, y+1, x, area+1)
	}

	switch check(visited, matrix, y-1, x, val) {
	case -1:
		key := fmt.Sprintf("l%d", y)
		sides[key] = append(sides[key], x)
	case 0:
		key := fmt.Sprintf("l%d", y)
		sides[key] = append(sides[key], x)
	case 1:
		area = discoverE(sides, visited, matrix, y-1, x, area+1)
	}

	switch check(visited, matrix, y, x+1, val) {
	case -1:
		key := fmt.Sprintf("e%d", x)
		sides[key] = append(sides[key], y)
	case 0:
		key := fmt.Sprintf("e%d", x)
		sides[key] = append(sides[key], y)
	case 1:
		area = discoverE(sides, visited, matrix, y, x+1, area+1)
	}

	switch check(visited, matrix, y, x-1, val) {
	case -1:
		key := fmt.Sprintf("w%d", x)
		sides[key] = append(sides[key], y)
	case 0:
		key := fmt.Sprintf("w%d", x)
		sides[key] = append(sides[key], y)
	case 1:
		area = discoverE(sides, visited, matrix, y, x-1, area+1)
	}

	return area
}

func pt1(matrix [][]rune) int {
	visited := make(map[string]bool)
	price := 0

	for y, row := range matrix {
		for x := range row {
			key := fmt.Sprintf("%d,%d", y, x)

			if !visited[key] {
				a, b := discover(visited, matrix, y, x, 1, 4)
				price += a * b
			}

		}
	}

	return price
}

func discover(visited map[string]bool, matrix [][]rune, y, x int, area, fence int) (int, int) {
	key := fmt.Sprintf("%d,%d", y, x)
	visited[key] = true
	val := matrix[y][x]

	if c := check(visited, matrix, y+1, x, val); c != -1 {
		if c == 2 {
			fence--
		}
		if c == 1 {
			fence--
			area, fence = discover(visited, matrix, y+1, x, area+1, fence+4)
		}
	}
	if c := check(visited, matrix, y-1, x, val); c != -1 {
		if c == 2 {
			fence--
		}
		if c == 1 {
			fence--
			area, fence = discover(visited, matrix, y-1, x, area+1, fence+4)
		}
	}
	if c := check(visited, matrix, y, x+1, val); c != -1 {
		if c == 2 {
			fence--
		}
		if c == 1 {
			fence--
			area, fence = discover(visited, matrix, y, x+1, area+1, fence+4)
		}
	}
	if c := check(visited, matrix, y, x-1, val); c != -1 {
		if c == 2 {
			fence--
		}
		if c == 1 {
			fence--
			area, fence = discover(visited, matrix, y, x-1, area+1, fence+4)
		}
	}

	return area, fence
}

func check(visited map[string]bool, matrix [][]rune, y, x int, t rune) int {
	if y < 0 || y >= len(matrix) || x < 0 || x >= len(matrix[0]) {
		return -1
	}

	key := fmt.Sprintf("%d,%d", y, x)
	if visited[key] {
		if matrix[y][x] == t {
			return 2
		} else {
			return 0
		}
	}

	if matrix[y][x] == t {
		return 1
	} else {
		return 0
	}
}
