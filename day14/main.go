package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	matrix := readInp("in_1.txt")

	fmt.Printf("Part 1: %d\n", pt1(matrix, 100, 101, 103))
	fmt.Printf("Part 2: %d\n", pt2(matrix, 101, 103))

}

func pt1(matrix [][]int, seconds, width, height int) int {
	ul, ur, ll, lr := 0, 0, 0, 0

	for _, robot := range matrix {
		py, px, vy, vx := robot[0], robot[1], robot[2], robot[3]

		newPy := (py + (vy*seconds)%height + height) % height
		newPx := (px + (vx*seconds)%width + width) % width

		ymid := height / 2
		xmid := width / 2

		switch {
		case newPy < ymid && newPx < xmid:
			ul++
		case newPy < ymid && newPx > xmid:
			ur++
		case newPy > ymid && newPx < xmid:
			ll++
		case newPy > ymid && newPx > xmid:
			lr++
		}
	}

	return ul * ur * ll * lr
}

func pt2(matrix [][]int, width, height int) int {

	var a [][]int = make([][]int, height)
	for i := range a {
		a[i] = make([]int, width)
	}
	for _, robot := range matrix {
		py, px := robot[0], robot[1]
		a[py][px] = a[py][px] + 1

	}

	cnt := 0

	for {
		tree := false

		for _, robot := range matrix {
			py, px, vy, vx := robot[0], robot[1], robot[2], robot[3]

			newPy := (py + (vy)%height + height) % height
			newPx := (px + (vx)%width + width) % width

			robot[0] = newPy
			robot[1] = newPx

			a[py][px] = a[py][px] - 1

			a[newPy][newPx] = a[newPy][newPx] + 1

			if newPx > 2 && newPx < width-2 && newPy < height-2 {
				if a[newPy+1][newPx-1] >= 1 && a[newPy+1][newPx+1] >= 1 {
					if a[newPy+2][newPx-2] >= 1 && a[newPy+2][newPx+2] >= 1 {
						if a[newPy+3][newPx-3] >= 1 && a[newPy+3][newPx+3] >= 1 {
							tree = true
						}
					}
				}
			}

		}
		cnt++

		if tree {
			break
		}

	}
	//util.PrintIntMatrix(a)

	return cnt
}

func readInp(path string) [][]int {
	dat, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	lines := strings.Split(string(dat), "\n")
	var result [][]int

	for _, line := range lines {
		parts := strings.Split(line, " ")

		a := strings.Split(parts[0][2:], ",")
		px, _ := strconv.Atoi(a[0])
		py, _ := strconv.Atoi(a[1])

		b := strings.Split(parts[1][2:], ",")
		vx, _ := strconv.Atoi(b[0])
		vy, _ := strconv.Atoi(b[1])

		rob := []int{py, px, vy, vx}
		result = append(result, rob)
	}

	return result
}
