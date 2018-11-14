package math

import (
	"errors"
	gomath "math"
)

// Average calculates the average of the numbers
func Average(numbers []float64) (float64, error) {
	count := len(numbers)
	if count <= 0 {
		return 0, errors.New("List of numbers cannot be empty")
	}

	var total float64
	for _, n := range numbers {
		total += n
	}

	avg := total / float64(count)
	return Round(avg), nil
}

// Round a number with 2 decimal precision
func Round(number float64) float64 {
	return gomath.Round(number*100) / 100
}
