package main

import (
	"fmt"

	"github.com/DylanSp/n-queens/go-verifier/verifier"
)

func main() {
	possibleSolution := verifier.Solution{
		N: 2,
		Queens: []verifier.Queen{
			{
				Row:    0,
				Column: 0,
			},
			{
				Row:    1,
				Column: 1,
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
