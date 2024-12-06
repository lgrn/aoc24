package day4

// findXmas takes a matrix of lines and characters, looks for the
// letter X, and if found calls scanX().
func findXmas(matrix [][]string) int {
	var xmasFound int
	// iterate over each line
	for l := 0; l < len(matrix); l++ {
		// iterate over each character
		for c := 0; c < len(matrix[l]); c++ {
			if string(matrix[l][c]) == "X" {
				// possible start of XMAS, pass matrix and position
				xmasFound += scanX(l, c, matrix)
			}
		}
	}
	return xmasFound
}

// findMas takes a matrix of lines and characters, looks for the
// letter M, and if found calls scanM().
func findMas(matrix [][]string) int {
	var masFound int
	// iterate over each line
	for l := 0; l < len(matrix); l++ {
		// iterate over each character
		for c := 0; c < len(matrix[l]); c++ {
			if string(matrix[l][c]) == "M" {
				// possible start of MAS-X, pass matrix and position
				masFound += scanM(l, c, matrix)
			}
		}
	}
	return masFound
}

// scanX() takes line and character position of a verified X in the
// matrix, then performs the following, if space allows it:
// - horizontal search (two axes)
// - vertical search (two axes)
// - diagonal search (four axes)
func scanX(l int, c int, matrix [][]string) int {

	var xmasCount int

	// horizontal search: forwards
	// only run if four letters fit
	if len(matrix[l])-c >= 4 {
		if matrix[l][c+1]+matrix[l][c+2]+matrix[l][c+3] == "MAS" {
			xmasCount++
		}
	}
	// horizontal search: backwards
	// only run if absolute position of X (c) is at least 3 (0,1,2)
	if c >= 3 {
		if matrix[l][c-1]+matrix[l][c-2]+matrix[l][c-3] == "MAS" {
			xmasCount++
		}
	}
	// vertical search: upwards
	// only run if line is at least 3 (0,1,2)
	if l >= 3 {
		if matrix[l-1][c]+matrix[l-2][c]+matrix[l-3][c] == "MAS" {
			xmasCount++
		}
	}
	// vertical search: downwards
	// only run if remaining space is at least 4
	if len(matrix)-l >= 4 {
		if matrix[l+1][c]+matrix[l+2][c]+matrix[l+3][c] == "MAS" {
			xmasCount++
		}
	}

	// horizontal search: upwards left
	// only run if line is at least 3 *and* c is at least 3
	if l >= 3 && c >= 3 {
		if matrix[l-1][c-1]+matrix[l-2][c-2]+matrix[l-3][c-3] == "MAS" {
			xmasCount++
		}
	}

	// horizontal search: upwards right
	// run if line is at least 3 *and* four letters fit to EOL
	if l >= 3 && len(matrix[l])-c >= 4 {
		if matrix[l-1][c+1]+matrix[l-2][c+2]+matrix[l-3][c+3] == "MAS" {
			xmasCount++
		}
	}

	// horizontal search: downwards left
	// run if 4 lines fit *and* c at least 3
	if len(matrix)-l >= 4 && c >= 3 {
		if matrix[l+1][c-1]+matrix[l+2][c-2]+matrix[l+3][c-3] == "MAS" {
			xmasCount++
		}
	}

	// horizontal search: downwards right
	// run if 4 lines fit *and* four letters fit to EOL
	if len(matrix)-l >= 4 && len(matrix[l])-c >= 4 {
		if matrix[l+1][c+1]+matrix[l+2][c+2]+matrix[l+3][c+3] == "MAS" {
			xmasCount++
		}
	}

	return xmasCount
}

// scanM() looks for two MAS in an X form beginning from a confirmed M.
// to avoid a cross being counted twice, the center "A" is replaced with
// a dot on match.
func scanM(l int, c int, matrix [][]string) int {

	var masCount int

	// S ?
	//  A
	// ? M
	// horizontal search: upwards left
	if l >= 2 && c >= 2 {
		if matrix[l-1][c-1]+matrix[l-2][c-2] == "AS" {
			// now check if center "A" has crossing M/S in either direction
			if matrix[l][c-2] == "M" && matrix[l-2][c] == "S" {
				masCount++
				matrix[l-1][c-1] = "."
			}
			if matrix[l][c-2] == "S" && matrix[l-2][c] == "M" {
				masCount++
				matrix[l-1][c-1] = "."
			}
		}
	}

	// ? S
	//  A
	// M ?
	// horizontal search: upwards right
	if l >= 2 && len(matrix[l])-c >= 3 {
		if matrix[l-1][c+1]+matrix[l-2][c+2] == "AS" {
			// now check if center "A" has crossing M/S in either direction
			if matrix[l][c+2] == "M" && matrix[l-2][c] == "S" {
				masCount++
				matrix[l-1][c+1] = "."
			}
			if matrix[l][c+2] == "S" && matrix[l-2][c] == "M" {
				masCount++
				matrix[l-1][c+1] = "."
			}
		}
	}

	// ? M
	//  A
	// S ?
	// horizontal search: downwards left
	if len(matrix)-l >= 3 && c >= 2 {
		if matrix[l+1][c-1]+matrix[l+2][c-2] == "AS" {
			// now check if center "A" has crossing M/S in either direction
			if matrix[l][c-2] == "M" && matrix[l+2][c] == "S" {
				masCount++
				matrix[l+1][c-1] = "."
			}
			if matrix[l][c-2] == "S" && matrix[l+2][c] == "M" {
				masCount++
				matrix[l+1][c-1] = "."
			}
		}
	}

	// M ?
	//  A
	// ? S
	// horizontal search: downwards right
	if len(matrix)-l >= 3 && len(matrix[l])-c >= 3 {
		if matrix[l+1][c+1]+matrix[l+2][c+2] == "AS" {
			// now check if center "A" has crossing M/S in either direction
			if matrix[l][c+2] == "M" && matrix[l+2][c] == "S" {
				masCount++
				matrix[l+1][c+1] = "."
			}
			if matrix[l][c+2] == "S" && matrix[l+2][c] == "M" {
				masCount++
				matrix[l+1][c+1] = "."
			}
		}
	}

	return masCount
}
