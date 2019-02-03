# Command-line implementation of Newton's method for calculating square-root

This uses Newton's approximation of iteratively improving over a guess to find the roots of a real-valued function. 

The method starts with an initial guess of 1, then iteratively improves over this by:

![eq](https://wikimedia.org/api/rest_v1/media/math/render/svg/710c11b9ec4568d1cfff49b7c7d41e0a7829a736)

until a sufficiently accurate value is obtained by determing that the change is less than *epsilon*.

## How to install
1. Run `git clone` on this repository
2. Run `go install NewtonSqrt.go` to install this into your `$GOBIN` path

## How to use
1. General structure: `NewtonSqrt <number> <epsilon>`. Note that if no `epsilon` is given, then a default value of `0.00000001` would be used.

