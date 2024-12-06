package day4

import (
	"bufio"
	"fmt"
	"os"
)

func Solve() (int, int) {

	file, err := os.Open("day4/input")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var xmasMatrix [][]string

	// build xmasMatrix so that we can navigate any line and character
	// as xmasMatrix[1][2]
	for scanner.Scan() {
		line := scanner.Text()
		var xmasRow []string
		for i := 0; i < len(line); i++ {
			xmasRow = append(xmasRow, string(line[i]))
		}
		xmasMatrix = append(xmasMatrix, xmasRow)
	}

	return findXmas(xmasMatrix), findMas(xmasMatrix)
}
