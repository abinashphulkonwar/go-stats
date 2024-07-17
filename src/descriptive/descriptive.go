package descriptive

import (
	"fmt"

	"github.com/abinashphulkonwar/stats/src/services"
)

func Mean(values []float64) float64 {
	var sum float64
	for _, v := range values {
		sum += v
	}
	sum = sum / float64(len(values))
	return sum
}

func Median(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	if len(values) <= 2 {
		return 1
	}
	sortedValues := services.Sort(values)

	if len(sortedValues)%2 == 0 {
		middleIndex := len(sortedValues) / 2
		first := sortedValues[middleIndex-1]
		second := sortedValues[middleIndex]
		fmt.Printf("middleIndex: %v, first: %v, second: %v\n", middleIndex, first, second)

		return (first + second) / 2
	}

	return sortedValues[len(sortedValues)/2]
}
