package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// Sqrt : based on Newton's method
func Sqrt(x float64, ep float64) float64 {

	if ep == 0.0 {
		ep = 0.00000001
	}

	if x < 0.0 {
		return 0.0
	}

	// initial guess
	z := 1.0
	var change float64

	// endless loop until change is less than epsilon
	for {
		change = (z*z - x) / (2 * z)
		if math.Abs(change) < ep {
			break
		} else {
			z -= change
		}
	}
	return z
}

func main() {
	args := os.Args
	num := 0.0
	ep := 0.0
	var err1 error
	var err2 error

	if len(args) < 3 {
		// no epsilon provided
		ep = 0.0
	} else {
		ep, err2 = strconv.ParseFloat(args[2], 64)
	}
	num, err1 = strconv.ParseFloat(args[1], 64)
	if err1 != nil || err2 != nil {
		panic("Uh oh! Conversion failed")
	}

	if num < 0.0 {
		panic("Number cannot be less than zero")
	}

	fmt.Printf("Result: %f\n", Sqrt(num, ep))
}
