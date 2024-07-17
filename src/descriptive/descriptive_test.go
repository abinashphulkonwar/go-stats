package descriptive_test

import (
	"testing"

	"github.com/abinashphulkonwar/stats/src/descriptive"
	"github.com/montanaflynn/stats"
)

func TestMedian(t *testing.T) {
	values := []float64{1, 8, 2, 3, 7, 9, 10, 4, 5, 6, 9}
	statsMedian, _ := stats.Median(values)

	median := descriptive.Median(values)

	if median != statsMedian {
		t.Errorf("Expected %v, got %v", statsMedian, median)
	}

}
