package day2

// isSafe() takes a slice of int, and returns whether safe (bool)
func isSafe(levels []int) bool {
	var increases int
	var decreases int

	// loop per value-pair, beginning with 1,0
	for i := 1; i < len(levels); i++ {
		// begin by storing diff between values, if negative: flip
		diff := levels[i-1] - levels[i]
		if diff < 0 {
			diff = -diff
		}

		// if diff is zero OR more than three, return immediately
		if diff < 1 || diff > 3 {
			return false
		}

		// store and react to increases/decreases
		if levels[i] < levels[i-1] {
			decreases++
			// this is a decrease. if we have seen increases in the
			// past, break immediately
			if increases > 0 {
				return false
			}
		}
		if levels[i] > levels[i-1] {
			increases++
			// this is an increase. if we have seen decreases in the
			// past, break immediately
			if decreases > 0 {
				return false
			}
		}

	}

	return true
}
