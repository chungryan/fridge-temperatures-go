package math

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	t.Run("Even number of elements", func(t *testing.T) {
		numbers := []float64{6.5, 50.65, 4.9, 3}
		Sort(numbers)

		if !reflect.DeepEqual(numbers, []float64{3, 4.9, 6.5, 50.65}) {
			t.Fatalf("Order is incorrect: %v", numbers)
		}
	})

	t.Run("Odd number of elements", func(t *testing.T) {
		numbers := []float64{478.8, 6.8, 10}
		Sort(numbers)

		if !reflect.DeepEqual(numbers, []float64{6.8, 10, 478.8}) {
			t.Fatalf("Order is incorrect: %v", numbers)
		}
	})

	t.Run("One number", func(t *testing.T) {
		numbers := []float64{79.89}
		Sort(numbers)

		if !reflect.DeepEqual(numbers, []float64{79.89}) {
			t.Fatalf("Order is incorrect: %v", numbers)
		}
	})

	t.Run("No numbers", func(t *testing.T) {
		numbers := []float64{}
		Sort(numbers)

		if !reflect.DeepEqual(numbers, []float64{}) {
			t.Fatalf("Order is incorrect: %v", numbers)
		}
	})
}
