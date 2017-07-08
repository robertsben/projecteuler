# Project Euler Solutions

## Usage

Run `go test` to run tests.

Run `go build .` to create executable; run `./projecteuler -problem {x} {args}` for problem solver.


## Architecture

There is an entry point `start.go`, with a main function, that handles input from the command line.
This dynamically maps to a solution function based on input.
This is possible because each solution is in a file named after the problem number in the `main` package, with a main function `Problem{X}` that maps onto `type Solution`.
