package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve() (int, int) {

	file, err := os.Open("day2/input")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var safeReports int
	var fixedReports int
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

		if isSafe(intLine) {
			// solution 1: increment if this []int is SAFE
			safeReports++
		} else if !isSafe(intLine) {
			// solution 2: if this []int is UNSAFE, try deleting one value
			// at a time to see if that makes it safe
			for i := 0; i < len(intLine); i++ {
				var intLineShort []int
				// intLineShort is intLine without int at index i
				// append _up to_ i
				intLineShort = append(intLineShort, intLine[:i]...)
				// append _following_ i
				intLineShort = append(intLineShort, intLine[i+1:]...)

				if isSafe(intLineShort) {
					fixedReports++
					break
				}
			}
		}
	}

	return safeReports, safeReports + fixedReports
}
