package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve() int {

	file, err := os.Open("day2/input")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var safeReports int
	for scanner.Scan() {
		// split each line on space
		split := strings.Split(scanner.Text(), " ")

		// iterate over split line, convert to int, add to intLine and
		// pass intLine to isSafe()

		var intLine []int
		for i := 0; i < len(split); i++ {
			// int conversion
			val, err := strconv.Atoi(strings.TrimSpace(split[i]))

			if err != nil {
				fmt.Println("Error converting to int")
				os.Exit(1)
			}
			intLine = append(intLine, val)
		}

		// finally, increment if this []int is "safe"
		if isSafe(intLine) {
			safeReports++
		}
	}

	return safeReports
}
