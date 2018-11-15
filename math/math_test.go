package math

import (
	"reflect"
	"testing"
)

func TestAverage(t *testing.T) {
	t.Run("Multiple integers", func(t *testing.T) {
		numbers := []float64{4, 6, 329, 5, 2}
		avg, err := Average(numbers)

		if err != nil {
			t.Fatalf("Unexpected error: %s", err.Error())
		}

		if avg != 69.2 {
			t.Fatalf("Expected average to be 69.2 but got %.2f", avg)
		}
	})

	t.Run("Multiple floats", func(t *testing.T) {
		numbers := []float64{4.56, 7.80, 3.05}
		avg, err := Average(numbers)

		if err != nil {
			t.Fatalf("Unexpected error: %s", err.Error())
		}

		if avg != 5.14 {
			t.Fatalf("Expected average to be 5.14 but got %.2f", avg)
		}
	})

	t.Run("Mixed floats and integers", func(t *testing.T) {
		numbers := []float64{4.56, 7}
		avg, err := Average(numbers)

		if err != nil {
			t.Fatalf("Unexpected error: %s", err.Error())
		}

		if avg != 5.78 {
			t.Fatalf("Expected average to be 5.78 but got %.2f", avg)
		}
	})

	t.Run("Single number", func(t *testing.T) {
		numbers := []float64{294.67}
		avg, err := Average(numbers)

		if err != nil {
			t.Fatalf("Unexpected error: %s", err.Error())
		}

		if avg != 294.67 {
			t.Fatalf("Expected average to be 294.67 but got %.2f", avg)
		}
	})

	t.Run("Empty list", func(t *testing.T) {
		numbers := []float64{}
		avg, err := Average(numbers)

		if err == nil || err.Error() != "List of numbers cannot be empty" {
			t.Fatalf("Expected empty list error but got: %v", err)
		}

		if avg != 0 {
			t.Fatalf("Unexpected non zero avg: %.2f", avg)
		}
	})
}

func TestValidateList(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		if err := validateList([]float64{}); err == nil || err.Error() != "List of numbers cannot be empty" {
			t.Fatalf("Expected empty list error but got: %v", err)
		}
	})

	t.Run("Not empty", func(t *testing.T) {
		if err := validateList([]float64{4, 5}); err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
	})
}

func TestRound(t *testing.T) {
	t.Run("Integer", func(t *testing.T) {
		round := Round(4)
		if round != 4 {
			t.Fatalf("Expected number to be rounded to 4 but got %f", round)
		}
	})

	t.Run("Two decimal float", func(t *testing.T) {
		round := Round(6.78)
		if round != 6.78 {
			t.Fatalf("Expected number to be rounded to 6.78 but got %f", round)
		}
	})

	t.Run("More than two decimal float, floor expected", func(t *testing.T) {
		round := Round(78.4839238739293)
		if round != 78.48 {
			t.Fatalf("Expected number to be rounded to 78.48 but got %f", round)
		}
	})

	t.Run("More than two decimal float, ceil expected", func(t *testing.T) {
		round := Round(45.738900)
		if round != 45.74 {
			t.Fatalf("Expected number to be rounded to 45.74 but got %f", round)
		}
	})

	t.Run("More than two decimal float, ceil expected with .5", func(t *testing.T) {
		round := Round(67.45555)
		if round != 67.46 {
			t.Fatalf("Expected number to be rounded to 67.46 but got %f", round)
		}
	})
}

func TestMedian(t *testing.T) {
	t.Run("Even number of elements", func(t *testing.T) {
		numbers := []float64{6.78, 3.44, 8, 56, 2.7, 3}
		median, err := Median(numbers)

		if err != nil {
			t.Fatalf("Unexpected error: %s", err.Error())
		}

		if median != 5.11 {
			t.Fatalf("Incorrect median: %.2f", median)
		}
	})

	t.Run("Odd number of elements", func(t *testing.T) {
		numbers := []float64{98.3, 4, 5}
		median, err := Median(numbers)

		if err != nil {
			t.Fatalf("Unexpected error: %s", err.Error())
		}

		if median != 5 {
			t.Fatalf("Incorrect median: %.2f", median)
		}
	})

	t.Run("One number", func(t *testing.T) {
		numbers := []float64{78.67}
		median, err := Median(numbers)

		if err != nil {
			t.Fatalf("Unexpected error: %s", err.Error())
		}

		if median != 78.67 {
			t.Fatalf("Incorrect median: %.2f", median)
		}
	})

	t.Run("Empty list", func(t *testing.T) {
		numbers := []float64{}
		median, err := Median(numbers)

		if err == nil || err.Error() != "List of numbers cannot be empty" {
			t.Fatalf("Expected empty list error but got: %v", err)
		}

		if median != 0 {
			t.Fatalf("Unexpected non zero avg: %.2f", median)
		}
	})
}

func TestMode(t *testing.T) {
	t.Run("Single mode element", func(t *testing.T) {
		numbers := []float64{5.5, 6.7, 7, 6.7}
		mode, err := Mode(numbers)

		if err != nil {
			t.Fatalf("Unexpected error: %s", err.Error())
		}

		if !reflect.DeepEqual(mode, []float64{6.7}) {
			t.Fatalf("Incorrect mode: %v", mode)
		}
	})

	t.Run("Multiple mode element", func(t *testing.T) {
		numbers := []float64{5.5, 3, 67.89, 3, 5.5, 9}
		mode, err := Mode(numbers)

		if err != nil {
			t.Fatalf("Unexpected error: %s", err.Error())
		}

		if !reflect.DeepEqual(mode, []float64{3, 5.5}) {
			t.Fatalf("Incorrect mode: %v", mode)
		}
	})

	t.Run("Empty list", func(t *testing.T) {
		numbers := []float64{}
		mode, err := Mode(numbers)

		if err == nil || err.Error() != "List of numbers cannot be empty" {
			t.Fatalf("Expected empty list error but got: %v", err)
		}

		if mode != nil {
			t.Fatalf("Unexpected mode: %v", mode)
		}
	})
}

func TestMax(t *testing.T) {
	t.Run("Multiple numbers", func(t *testing.T) {
		numbers := []float64{579.5, 3, 89.5}
		max, err := Max(numbers)

		if err != nil {
			t.Fatalf("Unexpected error: %s", err.Error())
		}

		if max != 579.5 {
			t.Fatalf("Incorrect max: %.2f", max)
		}
	})

	t.Run("One number", func(t *testing.T) {
		numbers := []float64{20}
		max, err := Max(numbers)

		if err != nil {
			t.Fatalf("Unexpected error: %s", err.Error())
		}

		if max != 20 {
			t.Fatalf("Incorrect max: %.2f", max)
		}
	})

	t.Run("Empty list", func(t *testing.T) {
		numbers := []float64{}
		max, err := Max(numbers)

		if err == nil || err.Error() != "List of numbers cannot be empty" {
			t.Fatalf("Expected empty list error but got: %v", err)
		}

		if max != 0 {
			t.Fatalf("Unexpected max: %v", max)
		}
	})
}
