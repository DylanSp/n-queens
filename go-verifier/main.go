package main

import (
	"fmt"

	"github.com/DylanSp/n-queens/go-verifier/verifier"
)

func makeSolution(queens [][]uint) verifier.Solution {
	solution := verifier.Solution{
		N:      verifier.BoardSize(len(queens)),
		Queens: []verifier.Queen{},
	}

	for _, queen := range queens {
		if len(queen) != 2 {
			panic(fmt.Sprint("queen is not exactly 2 coordinates; has ", len(queen), " elements"))
		}

		solution.Queens = append(solution.Queens, verifier.Queen{
			Row:    verifier.RowNumber(queen[0]),
			Column: verifier.ColumnNumber(queen[1]),
		})
	}

	return solution
}

func main() {
	// possibleSolution := verifier.Solution{
	// 	N: 2,
	// 	Queens: []verifier.Queen{
	// 		{
	// 			Row:    0,
	// 			Column: 0,
	// 		},
	// 		{
	// 			Row:    1,
	// 			Column: 1,
	// 		},
	// 	},
	// }

	possibleSolution := verifier.Solution{
		N: 5,
		Queens: []verifier.Queen{
			{
				Row:    0,
				Column: 0,
			},
			{
				Row:    1,
				Column: 2,
			},
			{
				Row:    2,
				Column: 4,
			},
			{
				Row:    3,
				Column: 1,
			},
			{
				Row:    4,
				Column: 3,
			},
		},
	}

	isSolutionCorrect := verifier.VerifySolution(possibleSolution)

	if isSolutionCorrect {
		fmt.Println("Correct solution :D")
	} else {
		fmt.Println("Incorrect solution :(")
	}
}
