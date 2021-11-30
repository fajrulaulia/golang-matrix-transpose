package strr

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestTransposeMatrix3x3(t *testing.T) {

	testCase := [][]uint8{
		{0, 1, 2},
		{4, 5, 6},
		{3, 4, 4},
	}

	testReality := "| 0  1  2 |\n| 4  5  6 |\n| 3  4  4 |\n"
	preTranspose := PrintMatrixAsString(testCase)
	assert.Equal(t, preTranspose, testReality)

	testReality = "| 0  4  3 |\n| 1  5  4 |\n| 2  6  4 |\n"
	matrix := TransposMatrix(testCase)
	result := PrintMatrixAsString(matrix)
	assert.Equal(t, result, testReality)
}

func TestTransposeMatrix3x2(t *testing.T) {

	testCase := [][]uint8{
		{0, 1},
		{4, 5},
		{3, 4},
	}

	testReality := "| 0  1 |\n| 4  5 |\n| 3  4 |\n"
	preTranspose := PrintMatrixAsString(testCase)
	assert.Equal(t, preTranspose, testReality)

	testReality = "| 0  4  3 |\n| 1  5  4 |\n"
	matrix := TransposMatrix(testCase)
	result := PrintMatrixAsString(matrix)
	assert.Equal(t, result, testReality)
}
