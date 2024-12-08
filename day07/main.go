package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Print("Part 1: ")
	fmt.Println(pt1(readLines("in_1.txt")))

	start := time.Now()
	fmt.Print("Part 2: ")
	fmt.Println(pt2(readLines("in_1.txt")))
	elapsed := time.Since(start)
	fmt.Printf("part 2 took %s\n", elapsed)

}

func pt2(matrix [][]int) int {
	sum := 0

	for _, eq := range matrix {
		target := eq[0]
		if compute2(eq[2:], 0, target) >= 1 {
			sum += target
		}
	}

	return sum
}

func compute2(eq []int, sum int, target int) int {
	if sum > target {
		return 0
	}
	if sum == 0 {
		return compute2(eq[1:], sum+eq[0], target)
	}

	if sum == target && len(eq) == 0 {
		return 1
	} else if len(eq) == 0 {
		return 0
	}

	return 0 + compute2(eq[1:], sum*eq[0], target) + compute2(eq[1:], sum+eq[0], target) + compute2(eq[1:], concat(sum, eq[0]), target)
}

func concat(a int, b int) int {
	c := strconv.Itoa(a) + strconv.Itoa(b)
	cnum, _ := strconv.Atoi(c)
	return cnum
}

func pt1(matrix [][]int) int {
	sum := 0

	for _, eq := range matrix {
		target := eq[0]
		if compute(eq[1:], 0, target) >= 1 {
			sum += target
		}
	}

	return sum
}

func compute(eq []int, sum int, target int) int {

	if sum == target && len(eq) == 0 {
		return 1
	} else if len(eq) == 0 {
		return 0
	}

	return 0 + compute(eq[1:], sum*eq[0], target) + compute(eq[1:], sum+eq[0], target)
}

func readLines(path string) [][]int {
	dat, _ := os.ReadFile(path)

	var matrix [][]int

	lines := strings.Split(string(dat), "\n")

	for _, line := range lines {
		var row []int
		parts := strings.Split(line, ":")
		goal, _ := strconv.Atoi(parts[0])
		row = append(row, goal)

		rest := strings.Split(parts[1], " ")

		for _, val := range rest {
			num, _ := strconv.Atoi(val)
			row = append(row, num)
		}

		matrix = append(matrix, row)
	}

	return matrix
}
