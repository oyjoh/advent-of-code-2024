package util

import (
	"reflect"
	"testing"
)

func TestPadRuneMatrixBasic(t *testing.T) {

	input := [][]rune{{'a', 'b'}, {'b', 'a'}}
	padSize := 1

	expectedResult := [][]rune{
		{'\x00', '\x00', '\x00', '\x00'},
		{'\x00', 'a', 'b', '\x00'},
		{'\x00', 'b', 'a', '\x00'},
		{'\x00', '\x00', '\x00', '\x00'},
	}

	result := PadRuneMatrix(input, padSize)

	if !reflect.DeepEqual(expectedResult, result) {
		t.Errorf("expected result was %v, but got %v instead", expectedResult, result)
	}
}

func TestPadRuneMatrixDouble(t *testing.T) {

	input := [][]rune{{'a', 'b'}, {'b', 'a'}}
	padSize := 2

	expectedResult := [][]rune{
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', 'a', 'b', '\x00', '\x00'},
		{'\x00', '\x00', 'b', 'a', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
		{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
	}

	result := PadRuneMatrix(input, padSize)

	if !reflect.DeepEqual(expectedResult, result) {
		t.Errorf("expected result was %v, but got %v instead", expectedResult, result)
	}
}
