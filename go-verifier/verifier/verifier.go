package verifier

import (
	"math"
)

type BoardSize uint

type RowNumber uint

type ColumnNumber uint

type Queen struct {
	Row    RowNumber
	Column ColumnNumber
}

type Solution struct {
	N      BoardSize
	Queens []Queen
}

func VerifySolution(solution Solution) bool {
	// validate that there are the correct number of queens
	if len(solution.Queens) < int(solution.N) {
		return false
	}

	// loop through all pairs of queens and check if any threaten each other
	for i, queen1 := range solution.Queens {
		for _, queen2 := range solution.Queens[(i + 1):] {
			if queensThreaten(queen1, queen2) {
				return false
			}
		}
	}

	return true
}

func queensThreaten(queen1 Queen, queen2 Queen) bool {
	// horizontally
	if queen1.Row == queen2.Row {
		return true
	}

	// vertically
	if queen1.Column == queen2.Column {
		return true
	}

	// diagonally
	rowDiff := int(queen1.Row) - int(queen2.Row)
	columnDiff := int(queen1.Column) - int(queen2.Column)

	// take absolute value to account for both diagonal directions
	return math.Abs(float64(rowDiff)) == math.Abs(float64(columnDiff))
}
