package main

import (
	"fmt"

	"github.com/oyjoh/advent-of-code-2024/util"
)

func main() {
	inp := util.RuneMatrix("in_1.txt")
	//inp := util.RuneMatrix("in_ex.txt")
	matrix := util.PadRuneMatrix(inp, 3)

	fmt.Println(p1(matrix))
	fmt.Println(p2(matrix))

}

func p2(matrix [][]rune) int {
	cnt := 0

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			if matrix[y][x] == 'A' {

				if string(matrix[y+1][x+1])+string(matrix[y-1][x-1]) == "MS" || string(matrix[y+1][x+1])+string(matrix[y-1][x-1]) == "SM" {
					if string(matrix[y+1][x-1])+string(matrix[y-1][x+1]) == "MS" || string(matrix[y+1][x-1])+string(matrix[y-1][x+1]) == "SM" {
						cnt++
					}
				}

			}
		}
	}

	return cnt
}

func p1(matrix [][]rune) int {
	cnt := 0

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			if matrix[y][x] == 'X' {
				if string(matrix[y][x+1])+string(matrix[y][x+2])+string(matrix[y][x+3]) == "MAS" {
					cnt++
				}
				if string(matrix[y][x-1])+string(matrix[y][x-2])+string(matrix[y][x-3]) == "MAS" {
					cnt++
				}
				if string(matrix[y+1][x])+string(matrix[y+2][x])+string(matrix[y+3][x]) == "MAS" {
					cnt++
				}
				if string(matrix[y-1][x])+string(matrix[y-2][x])+string(matrix[y-3][x]) == "MAS" {
					cnt++
				}
				if string(matrix[y+1][x+1])+string(matrix[y+2][x+2])+string(matrix[y+3][x+3]) == "MAS" {
					cnt++
				}
				if string(matrix[y-1][x-1])+string(matrix[y-2][x-2])+string(matrix[y-3][x-3]) == "MAS" {
					cnt++
				}
				if string(matrix[y-1][x+1])+string(matrix[y-2][x+2])+string(matrix[y-3][x+3]) == "MAS" {
					cnt++
				}
				if string(matrix[y+1][x-1])+string(matrix[y+2][x-2])+string(matrix[y+3][x-3]) == "MAS" {
					cnt++
				}

			}
		}

	}
	return cnt

}
