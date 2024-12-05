package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

func main() {
	ruleBook, updates := read("in_ex.txt")
	fmt.Println(pt1(ruleBook, updates))
	fmt.Println(pt2(ruleBook, updates))

	ruleBook, updates = read("in_1.txt")
	fmt.Println(pt1(ruleBook, updates))
	fmt.Println(pt2(ruleBook, updates))
}

func pt2(ruleBook map[int][]int, updates [][]int) int {
	res := 0
	for _, update := range updates {

		if validate(ruleBook, update) == -1 {
			continue
		} else {
			for problemIdx := validate(ruleBook, update); problemIdx != -1; problemIdx = validate(ruleBook, update) {
				update[problemIdx-1], update[problemIdx] = update[problemIdx], update[problemIdx-1]
			}
			res += update[len(update)/2]

		}

	}

	return res
}

func validate(ruleBook map[int][]int, arr []int) int {
	problemIdx := -1
	prev := mapset.NewSet[int]()

	for i, num := range arr {
		if prev.Contains(num) {
			problemIdx = i
			break
		}

		prev.Add(num)
		for _, n := range ruleBook[num] {
			prev.Add(n)
		}

	}
	return problemIdx
}

func pt1(ruleBook map[int][]int, updates [][]int) int {
	res := 0
	for _, update := range updates {
		valid := true
		prev := mapset.NewSet[int]()

		for _, num := range update {
			if prev.Contains(num) {
				valid = false
				break
			}

			prev.Add(num)
			for _, n := range ruleBook[num] {
				prev.Add(n)
			}

		}
		if valid {
			res += update[len(update)/2]

		}

	}

	return res
}

func read(path string) (map[int][]int, [][]int) {
	ruleBook := make(map[int][]int)
	var updates [][]int

	dat, _ := os.ReadFile(path)

	lines := strings.Split(string(dat), "\n")

	rules := true
	for _, line := range lines {
		if line == "" {
			rules = false
			continue
		}

		if rules {
			parts := strings.Split(line, "|")
			a, b := parts[0], parts[1]
			val, _ := strconv.Atoi(a)
			key, _ := strconv.Atoi(b)
			ruleBook[key] = append(ruleBook[key], val)
		} else {
			var ins []int
			parts := strings.Split(line, ",")
			for _, part := range parts {
				val, _ := strconv.Atoi(part)
				ins = append(ins, val)
			}
			updates = append(updates, ins)

		}

	}

	return ruleBook, updates
}
