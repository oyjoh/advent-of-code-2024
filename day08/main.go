package main

import (
	"fmt"

	"github.com/oyjoh/advent-of-code-2024/util"
)

func main() {

	matrix := util.RuneMatrix("in_1.txt")

	fmt.Println(pt1(matrix))
	fmt.Println(pt2(matrix))

}

func pt2(matrix [][]rune) int {
	visited := make(map[string]bool)
	antinodes := 0

	for y, row := range matrix {
		for x, val := range row {

			if val != '.' && val != '#' {
				antinodes += findNodes2(&visited, matrix, y, x)

			}

		}
	}
	return antinodes
}

func findNodes2(visited *map[string]bool, matrix [][]rune, yC, xC int) int {
	tp := matrix[yC][xC]
	cnt := 0

	for y, row := range matrix {
		for x, val := range row {

			if yC == y && xC == x {
				continue
			}

			if val == tp {

				dy, dx := dist(yC, xC, y, x)

				newY := yC + dy
				newX := xC + dx

				for inside(matrix, newY, newX) {
					key := fmt.Sprintf("%d,%d", newY, newX)
					if _, ok := (*visited)[key]; ok {

					} else {
						(*visited)[key] = true
						cnt++
					}

					newY += dy
					newX += dx
				}

			}

		}
	}

	return cnt
}

func pt1(matrix [][]rune) int {
	visited := make(map[string]bool)

	antinodes := 0

	for y, row := range matrix {
		for x, val := range row {

			if val != '.' && val != '#' {

				antinodes += findNodes(&visited, matrix, y, x)

			}

		}
	}
	return antinodes
}

func findNodes(visited *map[string]bool, matrix [][]rune, yC, xC int) int {
	tp := matrix[yC][xC]
	cnt := 0

	for y, row := range matrix {
		for x, val := range row {

			if yC == y && xC == x {
				continue
			}

			if val == tp {
				dy, dx := dist(yC, xC, y, x)

				newY := yC + dy*2
				newX := xC + dx*2
				if inside(matrix, newY, newX) {

					if _, ok := (*visited)[string(newY)+","+string(newX)]; ok {
						continue
					} else {
						(*visited)[string(newY)+","+string(newX)] = true
						cnt++
					}

				}
			}

		}
	}

	return cnt
}

func inside(matrix [][]rune, y, x int) bool {
	//fmt.Printf("y: %d, x: %d", y, x)
	if y < 0 || y >= len(matrix) {
		return false
	}
	if x < 0 || x >= len(matrix[0]) {
		return false
	}

	return true
}

func dist(ay, ax, by, bx int) (int, int) {
	dy := by - ay
	dx := bx - ax

	return dy, dx
}
