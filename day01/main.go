package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("in_1.txt")
	check(err)

	lines := strings.Split(string(dat), "\n")
	//fmt.Print(string(dat))

	//fmt.Println(solvea(lines))
	fmt.Println(solveb(lines))

}

func solveb(lines []string) int {
	var alst, blst []int

	for i := 0; i < len(lines); i++ {
		fields := strings.Fields(lines[i])

		numa, _ := strconv.Atoi(fields[0])
		numb, _ := strconv.Atoi(fields[1])

		alst = append(alst, numa)
		blst = append(blst, numb)
	}

	sort.Ints(alst)
	sort.Ints(blst)

	sum := 0

	for i := 0; i < len(alst); i++ {
		sum += countOcc(blst, alst[i]) * alst[i]
	}

	return sum
}

func solvea(lines []string) int {
	var alst, blst []int

	for i := 0; i < len(lines); i++ {
		fields := strings.Fields(lines[i])

		numa, _ := strconv.Atoi(fields[0])
		numb, _ := strconv.Atoi(fields[1])

		alst = append(alst, numa)
		blst = append(blst, numb)
	}

	sort.Ints(alst)
	sort.Ints(blst)

	sum := 0

	for i := 0; i < len(alst); i++ {
		sum += absInt(alst[i] - blst[i])
	}

	return sum
}

func countOcc(slice []int, num int) int {
	count := 0
	for _, v := range slice {
		if v == num {
			count++
		}
	}
	return count
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
