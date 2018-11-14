package math

import (
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
