package main

import (
	"log"
	c "matrix"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	Clear()

	M := [][]uint8{
		{0, 1, 2},
		{4, 5, 6},
		{3, 4, 4},
	}

	preTranspose := c.PrintMatrixAsString(M)
	log.Print("Pre Transpose : \n", preTranspose)

	matrix := c.TransposMatrix(M)
	result := c.PrintMatrixAsString(matrix)
	log.Print("Transpose : \n", result)

}

func Clear() {
	cmd := "clear"
	if runtime.GOOS == "windows" {
		cmd = "cls"
	}
	c := exec.Command(cmd)
	c.Stdout = os.Stdout
	c.Run()
}
