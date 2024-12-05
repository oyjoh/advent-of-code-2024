package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(solvePart1(read("in_ex.txt")))
	fmt.Println(solvePart2(read("in_ex.txt")))

	fmt.Println(solvePart1(read("in_1.txt")))
	fmt.Println(solvePart2(read("in_1.txt")))
}

func solvePart1(exp string) int {
	sum := 0
	for i := 0; i < len(exp)-4; i++ {
		sum += check(exp[i:])
	}
	return sum
}

func solvePart2(exp string) int {
	do := true
	sum := 0
	for i := 0; i < len(exp)-7; i++ {
		up := checkIns(exp[i:])
		if up == 1 {
			do = false
		} else if up == 2 {
			do = true
		}
		if do {
			sum += check(exp[i:])
		}
	}
	return sum
}

func checkIns(sub string) int {
	if sub[:7] == "don't()" {
		return 1
	}

	if sub[:4] == "do()" {
		return 2
	}

	return 0
}

func check(sub string) int {
	if sub[:4] == "mul(" {
		for i := 4; i < len(sub); i++ {
			if sub[i] == ')' {
				meat := sub[4:i]
				whitespace := regexp.MustCompile(`\s`).MatchString(meat)
				if whitespace {
					return 0
				}

				parts := strings.Split(meat, ",")

				if len(parts) > 2 || len(parts) < 2 {
					return 0
				}

				a, err1 := strconv.Atoi(parts[0])
				b, err2 := strconv.Atoi(parts[1])

				if err1 != nil || err2 != nil {
					return 0
				}

				if a > 999 || b > 999 {
					return 0
				}

				return a * b
			}
			//fmt.Println(unicode.IsDigit(rune(sub[i])))
		}
	}

	return 0
}

func read(filename string) string {
	fileContent, _ := os.ReadFile(filename)

	// Convert []byte to string
	text := string(fileContent)
	return text
}
