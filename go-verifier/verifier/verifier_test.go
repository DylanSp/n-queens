package verifier

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type VerifierTestSuite struct {
	suite.Suite
}

func TestVerifierTestSuite(t *testing.T) {
	suite.Run(t, new(VerifierTestSuite))
}

func (suite *VerifierTestSuite) TestNotEnoughQueensShouldBeFalse() {
	solution := Solution{
		N:      2,
		Queens: []Queen{},
	}

	suite.False(VerifySolution(solution))
}

func (suite *VerifierTestSuite) TestSameRowShouldBeFalse() {
	solution := Solution{
		N: 2,
		Queens: []Queen{
			{
				Row:    0,
				Column: 0,
			},
			{
				Row:    0,
				Column: 1,
			},
		},
	}

	suite.False(VerifySolution(solution))
}

func (suite *VerifierTestSuite) TestSameColumnShouldBeFalse() {
	solution := Solution{
		N: 2,
		Queens: []Queen{
			{
				Row:    0,
				Column: 0,
			},
			{
				Row:    1,
				Column: 0,
			},
		},
	}

	suite.False(VerifySolution(solution))
}

func (suite *VerifierTestSuite) TestDiagonalsShouldBeFalse() {
	// upper-left to lower-right
	solution1 := Solution{
		N: 2,
		Queens: []Queen{
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

	suite.False(VerifySolution(solution1))

	// upper-right to lower-left
	solution2 := Solution{
		N: 2,
		Queens: []Queen{
			{
				Row:    0,
				Column: 1,
			},
			{
				Row:    1,
				Column: 0,
			},
		},
	}

	suite.False(VerifySolution(solution2))
}

func (suite *VerifierTestSuite) TestCorrectSolutionShouldReturnTrue() {
	// solution for n=4 taken from https://developers.google.com/optimization/cp/queens
	solution := Solution{
		N: 4,
		Queens: []Queen{
			{
				Row:    0,
				Column: 2,
			},
			{
				Row:    1,
				Column: 0,
			},
			{
				Row:    2,
				Column: 3,
			},
			{
				Row:    3,
				Column: 1,
			},
		},
	}

	suite.True(VerifySolution(solution))
}
