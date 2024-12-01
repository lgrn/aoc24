package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Solve() (int, int) {

	file, err := os.Open("day1/input")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var list0 []int
	var list1 []int

	for scanner.Scan() {
		// split each line on space
		split := strings.SplitN(scanner.Text(), " ", 2)

		if len(split) != 2 {
			fmt.Println("Split error")
		}

		// convert strings to integers
		val0, err0 := strconv.Atoi(strings.TrimSpace(split[0]))
		val1, err1 := strconv.Atoi(strings.TrimSpace(split[1]))

		if err0 != nil || err1 != nil {
			fmt.Println("Error converting to int")
		}

		// append integers to list0/1
		list0 = append(list0, val0)
		list1 = append(list1, val1)
	}

	// for loop is done, sort the lists
	sort.Ints(list0)
	sort.Ints(list1)

	var solution1 int
	var solution2 int

	// 1. find the "distance" between each item in both lists
	// 2. find the "similarity score" being i * occurences to the right
	for i := 0; i < len(list0); i++ {
		num := list0[i] - list1[i]
		// if num is negative because list1 is actually a higher number,
		// invert it first to get "distance"
		if num < 0 {
			num = -num
		}
		// add distance to sum (solution 1)
		solution1 += num

		var occurences int

		// find how many times list0[i] occurs in list1
		for j := 0; j < len(list1); j++ {
			if list0[i] == list1[j] {
				occurences++
			}
		}

		solution2 += list0[i] * occurences
	}

	return solution1, solution2
}
