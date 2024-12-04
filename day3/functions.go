package day3

import (
	"fmt"
	"regexp"
	"strconv"
)

// calculateMul takes two strings extracted with regex, validates them as
// being integers, and if valid returns first * second. return 0 on error.
func calculateMul(first string, second string) int {

	int0, err0 := strconv.Atoi(first)
	int1, err1 := strconv.Atoi(second)

	if err0 != nil || err1 != nil {
		// at least one of the conversions failed
		fmt.Println(first, second, "failed")
		return 0
	}

	return int0 * int1

}

// findMul will find all mul() calls in a string and pass its integers
// to calculateMul(), returning the sum of them all.
func findMuls(input string) int {
	var mulSum int
	// find any mul() that contains two integers separated by comma
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	// add all matches (-1)
	matches := re.FindAllStringSubmatch(input, -1)
	// iterate over the matches we found and pass the two capture
	// groups at index 1,2 (strings) to calculateMul()
	for _, match := range matches {
		if len(match) == 3 {
			// match[0] is the full string, 1/2 are the capture groups,
			// pass them as strings and let the function handle it.
			mulSum += calculateMul(match[1], match[2])
		} else {
			fmt.Println("Refusing to pass: ", match)
		}
	}
	return mulSum
}
