package verifier

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
	panic("Not yet implemented")
}
