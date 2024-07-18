package descriptive_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/abinashphulkonwar/stats/src/descriptive"
	"github.com/montanaflynn/stats"
)

func TestWidghtedMean(t *testing.T) {
	values := []float64{1, 8, 2, 3, 7, 9, 10, 4, 5, 6, 9}
	statsMean, _ := stats.Mean(values)

	mean := descriptive.WidhtedMean(values)
	if mean != statsMean {
		t.Errorf("Expected %v, got %v", statsMean, mean)
	}

}

func TestMedian(t *testing.T) {
	values := []float64{1, 8, 2, 3, 7, 9, 10, 4, 5, 6, 9}
	statsMedian, _ := stats.Median(values)

	median := descriptive.Median(values)

	if median != statsMedian {
		t.Errorf("Expected %v, got %v", statsMedian, median)
	}

}

func TestMode(t *testing.T) {
	values := []float64{1, 8, 2, 3, 7, 9, 10, 4, 5, 6, 6, 7, 9, 9}
	statsMode, _ := stats.Mode(values)

	mode := descriptive.Mode(values)
	if len(mode) != len(statsMode) {
		t.Errorf("Expected %v, got %v", statsMode, mode)
	}
	max := math.Max(float64(len(statsMode)), float64(len(mode)))

	for i := 0; i < int(max)-1; i++ {
		if mode[i] != statsMode[i] {
			t.Errorf("Expected %v, got %v", statsMode, mode)
		}
	}
}

func TestQuantials(t *testing.T) {

	values := []float64{}

	for i := 1; i <= 100; i++ {
		values = append(values, float64(i))
	}
	statsQuantials, _ := stats.Quartile(values)
	quantials := descriptive.Quantile(values)
	fmt.Printf("stats Q1: %v; Q2: %v; Q3: %v\n", statsQuantials.Q1, statsQuantials.Q2, statsQuantials.Q3)
	fmt.Printf("descriptive Q1: %v; Q2: %v; Q3: %v; Q4: %v\n", quantials.Q1, quantials.Q2, quantials.Q3, quantials.Q4)

}
