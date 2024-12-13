package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	matrix := readInp("in_ex.txt")
	fmt.Println(matrix)
	fmt.Printf("Part 1: %d", pt1(matrix))
}

func pt1(matrix [][]int) int {
	cost := 0

	for _, machine := range matrix {

		cost += press(0, 0, 0, 0, 0, machine)

		break
	}

	return cost
}

func press(cost, y, x, pa, pb int, dat []int) int {
	fmt.Printf("y: %d, x: %d\n", y, x)
	fmt.Printf("pa: %d, pb: %d\n", pa, pb)

	if y == dat[4] && x == dat[5] {
		return cost
	}
	if y > dat[4] || x > dat[5] {
		return 0
	}
	if pa >= 100 || pb >= 100 {
		return 0
	}

	a, b := press(cost+3, y+dat[0], x+dat[1], pa+1, pb, dat), press(cost+1, y+dat[2], x+dat[3], pa, pb+1, dat)

	if a > b && b > 0 {
		return b
	} else {
		return a
	}

}

func readInp(path string) [][]int {
	dat, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	lines := strings.Split(string(dat), "\n\n")
	var result [][]int

	for _, line := range lines {
		parts := strings.Split(line, "\n")

		a := strings.Split(parts[0], " ")
		b := strings.Split(parts[1], " ")
		c := strings.Split(parts[2], " ")

		ax, _ := strconv.Atoi(a[2][2 : len(a[2])-1])
		ay, _ := strconv.Atoi(a[3][2:])
		bx, _ := strconv.Atoi(b[2][2 : len(b[2])-1])
		by, _ := strconv.Atoi(b[3][2:])
		px, _ := strconv.Atoi(c[1][2 : len(c[1])-1])
		py, _ := strconv.Atoi(c[2][2:])

		vals := []int{ay, ax, by, bx, py, px}

		result = append(result, vals)
	}

	return result
}
