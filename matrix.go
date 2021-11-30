package strr

import (
	"fmt"
)

func TransposMatrix(M [][]uint8) [][]uint8 {
	var result [][]uint8

	for i := 0; i < len(M[0]); i++ {
		var result2s []uint8
		for j := 0; j < len(M); j++ {
			result2s = append(result2s, M[j][i])
		}
		result = append(result, result2s)
	}

	return result
}

func PrintMatrixAsString(M [][]uint8) string {
	var data string
	for _, i := range M {
		data += "|"
		for _, j := range i {
			data += fmt.Sprintf(" %d ", j)
		}
		data += "|"
		data += "\n"
	}
	return data
}
