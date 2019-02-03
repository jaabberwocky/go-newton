# Command-line implementation of Newton's method for calculating square-root

This uses Newton's approximation of iteratively improving over a guess to find the roots of a real-valued function. 

## How to install
1. Run `git clone` on this repository
2. Run `go install NewtonSqrt.go` to install this into your `$GOBIN` path

## How to use
1. General structure: `NewtonSqrt <number> <epsilon>`. Note that if no `epsilon` is given, then a default value of `0.00000001` would be used.

