package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func main() {
	matrix := readInp("in_1.txt")
	//fmt.Println(matrix)
	fmt.Printf("Part 1: %d\n", pt1(matrix))
	fmt.Printf("Part 2: %d\n", pt2(matrix))
}

func pt2(matrix [][]int) int64 {
	var cost int64 = 0

	for _, machine := range matrix {

		ay, ax, by, bx, py, px := machine[0], machine[1], machine[2], machine[3], machine[4], machine[5]

		// part 2
		var fpy int64 = int64(py) + 10000000000000
		var fpx int64 = int64(px) + 10000000000000

		A := mat.NewDense(2, 2, []float64{float64(ay), float64(by), float64(ax), float64(bx)})
		B := mat.NewVecDense(2, []float64{float64(fpy), float64(fpx)})

		var x mat.VecDense
		if err := x.SolveVec(A, B); err != nil {
			fmt.Println(err)
		}
		a := x.RawVector().Data[0]
		b := x.RawVector().Data[1]

		// check if we can convert to int
		const epsilon = 1e-5 // Margin of error
		if _, frac := math.Modf(math.Abs(a)); frac < epsilon || frac > 1.0-epsilon {
			if _, frac := math.Modf(math.Abs(b)); frac < epsilon || frac > 1.0-epsilon {
				aI := int64(math.Round(a))
				bI := int64(math.Round(b))

				if aI >= 0 && bI >= 0 {
					//fmt.Println(machine)
					// fmt.Printf("a: %d, b: %d\n", aI, bI)
					// fmt.Println(x.RawVector().Data)
					cost += aI * 3
					cost += bI
				}
			}
		}

	}

	return cost
}

func pt1(matrix [][]int) int {
	cost := 0

	for _, machine := range matrix {

		ay, ax, by, bx, py, px := machine[0], machine[1], machine[2], machine[3], machine[4], machine[5]

		A := mat.NewDense(2, 2, []float64{float64(ay), float64(by), float64(ax), float64(bx)})
		B := mat.NewVecDense(2, []float64{float64(py), float64(px)})

		var x mat.VecDense
		if err := x.SolveVec(A, B); err != nil {
			fmt.Println(err)
		}
		a := x.RawVector().Data[0]
		b := x.RawVector().Data[1]

		// check if we can convert to int
		const epsilon = 1e-9 // Margin of error
		if _, frac := math.Modf(math.Abs(a)); frac < epsilon || frac > 1.0-epsilon {
			if _, frac := math.Modf(math.Abs(b)); frac < epsilon || frac > 1.0-epsilon {
				aI := int(math.Round(a))
				bI := int(math.Round(b))

				if aI <= 100 && bI <= 100 && aI >= 0 && bI >= 0 {
					// fmt.Println(machine)
					// fmt.Printf("a: %d, b: %d\n", aI, bI)
					// fmt.Println(x.RawVector().Data)
					cost += aI * 3
					cost += bI
				}
			}
		}

	}

	return cost
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
