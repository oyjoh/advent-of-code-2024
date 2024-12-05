package util

func PadRuneMatrix(matrix [][]rune, padSize int) [][]rune {

	paddedMatrix := make([][]rune, len(matrix)+padSize*2)
	for i := range paddedMatrix {
		paddedMatrix[i] = make([]rune, len(matrix[0])+padSize*2)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			paddedMatrix[i+padSize][j+padSize] = matrix[i][j]
		}
	}

	return paddedMatrix
}
