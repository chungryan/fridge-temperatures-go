package math

import (
	"errors"
	gomath "math"
)

// Average calculates the average of the numbers
func Average(numbers []float64) (float64, error) {
	if err := validateList(numbers); err != nil {
		return 0, err
	}

	var total float64
	for _, n := range numbers {
		total += n
	}

	avg := total / float64(len(numbers))
	return Round(avg), nil
}

// Round a number with 2 decimal precision
// This function could be made more reusable by accepting the number of decimals as a second argument
func Round(number float64) float64 {
	return gomath.Round(number*100) / 100
}

// Check if the list of numbers is valid, ie not empty
func validateList(numbers []float64) error {
	if len(numbers) <= 0 {
		return errors.New("List of numbers cannot be empty")
	}

	return nil
}

// Median returns the median of a list
func Median(numbers []float64) (float64, error) {
	if err := validateList(numbers); err != nil {
		return 0, err
	}

	Sort(numbers)

	l := len(numbers)
	m := l / 2
	if l%2 == 0 {
		return Average(numbers[m-1 : m+1])
	}

	return numbers[m], nil
}

// Mode returns the mode of the series
func Mode(numbers []float64) ([]float64, error) {
	if err := validateList(numbers); err != nil {
		return nil, err
	}

	counts := map[float64]uint{}
	for _, n := range numbers {
		counts[n]++
	}

	countSeries := []float64{}
	for _, c := range counts {
		countSeries = append(countSeries, float64(c))
	}

	max, _ := Max(countSeries)

	modes := []float64{}
	for n, c := range counts {
		if c == uint(max) {
			modes = append(modes, n)
		}
	}

	return Sort(modes), nil
}

// Max returns the max number of the series
func Max(numbers []float64) (float64, error) {
	if err := validateList(numbers); err != nil {
		return 0, err
	}

	Sort(numbers)
	return numbers[len(numbers)-1], nil
}
