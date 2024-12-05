package util

import (
	"fmt"
	"testing"
)

func TestReadRuneMatrixBasic(t *testing.T) {
	f := "rune-matrix.txt"

	result := RuneMatrix(f)
	fmt.Println(result)
}
