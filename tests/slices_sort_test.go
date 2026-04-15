package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	slices"Library/internal"
)

func TestBubbleSort(t *testing.T) {
	intSlice := []int{5, 2, 9, 1, 5, 6}
	expectedIntSlice := []int{1, 2, 5, 5, 6, 9}
	slices.BubbleSort(intSlice)
	assert.Equal(t, expectedIntSlice, intSlice)

	floatSlice := []float64{3.1, 2.2, 5.5, 1.1, 4.4}
	expectedFloatSlice := []float64{1.1, 2.2, 3.1, 4.4, 5.5}
	slices.BubbleSort(floatSlice)
	assert.Equal(t, expectedFloatSlice, floatSlice)
}

func TestInsertionSort(t *testing.T) {
	intSlice := []int{5, 2, 9, 1, 5, 6}
	expectedIntSlice := []int{1, 2, 5, 5, 6, 9}
	slices.InsertionSort(intSlice)
	assert.Equal(t, expectedIntSlice, intSlice)

	floatSlice := []float64{3.1, 2.2, 5.5, 1.1, 4.4}
	expectedFloatSlice := []float64{1.1, 2.2, 3.1, 4.4, 5.5}
	slices.InsertionSort(floatSlice)
	assert.Equal(t, expectedFloatSlice, floatSlice)
}

func TestSelectionSort(t *testing.T) {
	intSlice := []int{5, 2, 9, 1, 5, 6}
	expectedIntSlice := []int{1, 2, 5, 5, 6, 9}
	slices.SelectionSort(intSlice)
	assert.Equal(t, expectedIntSlice, intSlice)

	floatSlice := []float64{3.1, 2.2, 5.5, 1.1, 4.4}
	expectedFloatSlice := []float64{1.1, 2.2, 3.1, 4.4, 5.5}
	slices.SelectionSort(floatSlice)
	assert.Equal(t, expectedFloatSlice, floatSlice)
}

func TestMergeSort(t *testing.T) {
	intSlice := []int{5, 2, 9, 1, 5, 6}
	expectedIntSlice := []int{1, 2, 5, 5, 6, 9}
	sortedIntSlice := slices.MergeSort(intSlice)
	assert.Equal(t, expectedIntSlice, sortedIntSlice)

	floatSlice := []float64{3.1, 2.2, 5.5, 1.1, 4.4}
	expectedFloatSlice := []float64{1.1, 2.2, 3.1, 4.4, 5.5}
	sortedFloatSlice := slices.MergeSort(floatSlice)
	assert.Equal(t, expectedFloatSlice, sortedFloatSlice)
}
