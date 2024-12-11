package main

import (
	"fmt"
	"strconv"
)

func main() {
	//data := []int{125, 17}
	data := []int{9759, 0, 256219, 60, 1175776, 113, 6, 92833}

	fmt.Printf("Part 1: %d\n", pt1(data, 25))
	fmt.Printf("Part 2: %d\n", pt1(data, 75))
}

func pt1(input []int, blinks int) int {
	stones := 0

	memo := make(map[string]int)

	for _, num := range input {
		stones += blink(num, blinks, memo)
	}

	return stones
}

func blink(num int, blinks int, memo map[string]int) int {
	if blinks == 0 {
		return 1
	}

	key := strconv.Itoa(num) + "," + strconv.Itoa(blinks)
	if _, ok := memo[key]; ok {
		return memo[key]
	} else {
		if num == 0 {
			memo[key] = blink(1, blinks-1, memo)
		} else if len(strconv.Itoa(num))%2 == 0 {
			s := strconv.Itoa(num)
			a, _ := strconv.Atoi(s[:len(s)/2])
			b, _ := strconv.Atoi(s[len(s)/2:])
			memo[key] = blink(a, blinks-1, memo) + blink(b, blinks-1, memo)
		} else {
			memo[key] = blink(num*2024, blinks-1, memo)
		}
	}

	return memo[key]
}
