package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func Solve() (int, int) {

	file, err := os.Open("day3/input")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var mulSum int
	var cleanMulSum int
	var fullInput string

	// fullInput is the 'input' file with no linebreaks
	for scanner.Scan() {
		fullInput += scanner.Text()
	}

	// solution 1:
	mulSum = findMuls(fullInput)

	// solution 2:
	// mul() calls can be disabled when after don't() until next do()
	// appears, so these substrings should simply be removed.

	// regex to match anything between literal don't()->do()
	re := regexp.MustCompile(`don't\(\)(.*?)do\(\)`)
	// wipe out all regex matches, save into cleanString
	cleanString := re.ReplaceAllString(fullInput, "")

	// with a clean string, we can do solution #1 all over again
	cleanMulSum = findMuls(cleanString)

	return mulSum, cleanMulSum
}
