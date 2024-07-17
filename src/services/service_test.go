package services_test

import (
	"testing"

	"github.com/abinashphulkonwar/stats/src/services"
)

func TestSort(t *testing.T) {

	values := []float64{1, 8, 2, 3, 7, 9, 10, 4, 5, 6}
	correctValue := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sortedValues := services.Sort(values)

	for i, v := range sortedValues {
		if v != correctValue[i] {
			t.Errorf("Expected %v, got %v", correctValue[i], v)
		}
	}

}
