package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	path := "in_1.txt"

	inp, _ := os.ReadFile(path)
	sl := toSlice(inp)

	fmt.Println(pt1(sl))

	// lazy
	inp, _ = os.ReadFile(path)
	sl = toSlice(inp)

	fmt.Println(pt2(sl))
}

func pt2(inp []int) int {
	lastFileIdx := len(inp) - 1
	lastFileId := ((len(inp) - 1) / 2)
	lastFileSize := inp[lastFileIdx]

	spaces := make(map[int][]int)
	sizes := make(map[int]int)

	for i := lastFileIdx; i > 1; i = i - 2 {

		for j := 1; j < i; j = j + 2 {

			space := inp[j]

			if space >= lastFileSize {
				inp[j] = inp[j] - lastFileSize
				inp[lastFileIdx-1] = inp[lastFileIdx-1] + lastFileSize
				inp[i] = 0
				spaces[j] = append(spaces[j], lastFileId)
				sizes[lastFileId] = lastFileSize
				break
			}

		}

		lastFileIdx = lastFileIdx - 2
		lastFileId--
		lastFileSize = inp[lastFileIdx]
	}

	endaEnTeller := 0
	res := 0
	for idx, val := range inp {

		if idx%2 == 0 {
			for i := 0; i < val; i++ {
				id := idx / 2
				//fmt.Print(id)
				res += endaEnTeller * id
				endaEnTeller++
			}
		} else {
			for i := 0; i < len(spaces[idx]); i++ {
				for k := 0; k < sizes[spaces[idx][i]]; k++ {
					//fmt.Print(spaces[idx][i])
					res += endaEnTeller * spaces[idx][i]
					endaEnTeller++
				}
			}
			if val > 0 {
				for p := 0; p < val; p++ {
					//fmt.Print(".")
					endaEnTeller++
				}
			}
		}
	}
	return res
}

func pt1(inp []int) int {
	grandSum := 0

	grandIdx := 0
	curFileId := 0

	lastFileIdx := len(inp) - 1
	lastFileId := ((len(inp) - 1) / 2)
	lastFileVal := inp[lastFileIdx]

	for idx, num := range inp {
		if idx > lastFileIdx {
			break
		}

		space := idx%2 != 0

		for k := 0; k < num; k++ {

			if space {
				grandSum += grandIdx * lastFileId

				lastFileVal--
				inp[lastFileIdx] = inp[lastFileIdx] - 1

				if lastFileVal == 0 {
					lastFileIdx -= 2
					lastFileId--
					lastFileVal = inp[lastFileIdx]
				}
			} else {
				grandSum += grandIdx * curFileId
			}
			grandIdx++

		}

		if !space {
			curFileId++
		}
	}

	return grandSum
}

func toSlice(inp []byte) []int {
	sl := make([]int, len(inp))

	for idx, val := range inp {
		num, _ := strconv.Atoi(string(val))
		sl[idx] = num
	}

	return sl
}
