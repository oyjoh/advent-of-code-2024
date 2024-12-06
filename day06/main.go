package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/oyjoh/advent-of-code-2024/util"
)

func main() {
	example := "in_1.txt"
	matrix := util.RuneMatrix(example)
	matrix = util.PadRuneMatrix(matrix, 1)

	//fmt.Println(pt1(matrix))
	start := time.Now()
	fmt.Println(pt2Concurrent(matrix))
	elapsed := time.Since(start)
	fmt.Printf("concurrent version took %s\n", elapsed)

	matrix = util.RuneMatrix(example)
	matrix = util.PadRuneMatrix(matrix, 1)

	//fmt.Println(pt1(matrix))
	start = time.Now()
	fmt.Println(pt2(matrix))
	elapsed = time.Since(start)
	fmt.Printf("dumb version took %s\n", elapsed)
}

func pt2Concurrent(matrix [][]rune) int {
	cntObstacles := 0

	var wg sync.WaitGroup
	results := make(chan int, len(matrix)*len(matrix[0]))

	for y := 1; y < len(matrix)-1; y++ {
		for x := 1; x < len(matrix[0])-1; x++ {
			if matrix[y][x] == '.' {
				wg.Add(1)
				go func(y, x int) {
					defer wg.Done()
					matrixCopy := make([][]rune, len(matrix))
					for i := range matrix {
						matrixCopy[i] = make([]rune, len(matrix[i]))
						copy(matrixCopy[i], matrix[i])
					}
					matrixCopy[y][x] = '#'
					if pt1(matrixCopy) == -1 {
						results <- 1
					} else {
						results <- 0
					}
				}(y, x)
			}
		}
	}

	wg.Wait()
	close(results)

	for result := range results {
		cntObstacles += result
	}

	return cntObstacles
}

func pt2(matrix [][]rune) int {
	cntObstacles := 0

	for y := 1; y < len(matrix)-1; y++ {
		for x := 1; x < len(matrix[0])-1; x++ {
			if matrix[y][x] == '.' {
				matrixCopy := make([][]rune, len(matrix))
				for i := range matrix {
					matrixCopy[i] = make([]rune, len(matrix[i]))
					copy(matrixCopy[i], matrix[i])
				}
				matrixCopy[y][x] = '#'
				if pt1(matrixCopy) == -1 {
					cntObstacles++
				}
			}
		}
	}
	return cntObstacles
}

func pt1(matrix [][]rune) int {
	yPos, xPos := getStart(matrix)
	dy, dx := -1, 0

	cntVisited := 1
	dirMap := make(map[string]string) // part 2

	for insideMatrix(matrix, yPos, xPos) {

		if matrix[yPos+dy][xPos+dx] == '#' {
			dy, dx = r90(dy, dx)
			continue
		}
		yPos += dy
		xPos += dx

		if matrix[yPos][xPos] == '.' {
			matrix[yPos][xPos] = 'X'
			dirMap[string(yPos)+string(xPos)] = string(dy) + string(dx) //part 2
			cntVisited++
		} else if matrix[yPos][xPos] == 'X' {
			if dirMap[string(yPos)+string(xPos)] == string(dy)+string(dx) { // part 2
				return -1
			}
		}
	}
	return cntVisited
}

func insideMatrix(matrix [][]rune, y, x int) bool {
	if y < 1 || y >= len(matrix)-1 {
		return false
	}
	if x < 1 || x >= len(matrix[0])-1 {
		return false
	}

	return true
}

func r90(dy, dx int) (int, int) {
	return dx, -dy
}

func getStart(matrix [][]rune) (int, int) {
	for y, row := range matrix {
		for x, cell := range row {
			if cell == '^' {
				return y, x
			}
		}
	}
	return -1, -1
}
