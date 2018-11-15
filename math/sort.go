package math

// Sort an array of floats by ascending order
// The specifications didn't mention not to use a library for sorting but for the fun of the exercise,
// I decided to implement to the well known merge sort algorithm
func Sort(numbers []float64) {
	l := len(numbers)
	if l > 1 {
		m := l / 2
		Sort(numbers[0:m])
		Sort(numbers[m:l])
		mergeSort(numbers, m)
	}
}

// mergeSort two sub-arrays,
// the left sub-array being 0 to m-1 elements,
// and the right sub-array being from m to last elements
func mergeSort(numbers []float64, m int) {
	l := len(numbers)
	tmp := make([]float64, l)
	copy(tmp, numbers)

	i, j, k := 0, m, 0
	for i < m && j < l {
		if tmp[i] <= tmp[j] {
			numbers[k] = tmp[i]
			i++
		} else {
			numbers[k] = tmp[j]
			j++
		}
		k++
	}

	if i < m {
		k += copy(numbers[k:], tmp[i:m])
	}

	if j < l {
		copy(numbers[k:], tmp[j:])
	}
}
