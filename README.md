# N-Queens solver and verifier

A solver for the [n-queens problem](https://en.wikipedia.org/wiki/Eight_queens_puzzle) written in [SWI Prolog](https://www.swi-prolog.org/), as well as a solution verifier written in Go.

## Finding a solution

First, make sure to [install SWI Prolog](https://www.swi-prolog.org/Download.html). Once it's installed:

1. Run `swipl` at the command line in the `prolog-solver` directory to open the Prolog console.
1. Run `[solver].` (note the `.`) to load the solver program.
1. Find a solution by running the query `solution(N, Queens).` with N replaced by an actual integer. For example, `solution(4, Queens).` will generate a solution with 4 queens for a 4x4 board.

## Verifying a solution

1. Hardcode your solution in `go-verifier/main.go` as the value for `possibleSolution`, following the examples there.
1. In the `go-verifier` directory, run `go run main.go`.
