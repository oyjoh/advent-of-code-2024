package main

import (
	"fmt"

	"github.com/oyjoh/advent-of-code-2024/util"
)

func main() {

	path := "in_1.txt"
	matrix := util.ReadIntMatrix(path)

	fmt.Println(pt1(matrix))
	fmt.Println(pt2(matrix))
}

func pt1(matrix [][]int) int {
	totalScore := 0

	for y, row := range matrix {
		for x, val := range row {
			if val == 0 {
				visited := make(map[string]bool)
				totalScore += find(visited, matrix, y, x, 0)
			}
		}
	}
	return totalScore
}

func pt2(matrix [][]int) int {
	totalScore := 0

	for y, row := range matrix {
		for x, val := range row {
			if val == 0 {
				totalScore += findAll(matrix, y, x, 0)
			}
		}
	}
	return totalScore
}

func find(visited map[string]bool, matrix [][]int, y, x, num int) int {
	if !inside(matrix, y, x) {
		return 0
	}
	if num != matrix[y][x] {
		return 0
	}
	if matrix[y][x] == 9 {
		key := string(y) + "," + string(x)
		if visited[key] {
			return 0
		} else {
			visited[key] = true
			return 1
		}
	}

	return find(visited, matrix, y+1, x, num+1) + find(visited, matrix, y-1, x, num+1) + find(visited, matrix, y, x+1, num+1) + find(visited, matrix, y, x-1, num+1)

}

func inside(matrix [][]int, y, x int) bool {
	return y >= 0 && y < len(matrix) && x >= 0 && x < len(matrix[0])
}

func findAll(matrix [][]int, y, x, num int) int {
	if !inside(matrix, y, x) {
		return 0
	}
	if num != matrix[y][x] {
		return 0
	}
	if matrix[y][x] == 9 {
		return 1
	}

	return findAll(matrix, y+1, x, num+1) + findAll(matrix, y-1, x, num+1) + findAll(matrix, y, x+1, num+1) + findAll(matrix, y, x-1, num+1)

}
