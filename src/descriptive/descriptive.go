package descriptive

import (
	"github.com/abinashphulkonwar/stats/src/services"
)

/*
Mean: Σx / n
*/
func Mean(values []float64) float64 {
	var sum float64
	for _, v := range values {
		sum += v
	}
	sum = sum / float64(len(values))
	return sum
}

func Frequency(values []float64, item float64) int {
	count := 0
	for _, v := range values {
		if v == item {
			count++
		}
	}
	return count
}
func Frequencys(values []float64) map[float64]int {
	f := make(map[float64]int)
	for _, v := range values {
		_, ok := f[v]
		if ok {
			f[v] += 1
		} else {
			f[v] = 1
		}
	}
	return f
}

/*
mean = ΣF(x) / ΣF(x)
*/
func WidhtedMean(values []float64) float64 {
	var sum float64
	f := make(map[float64]int)
	for _, v := range values {
		_, ok := f[v]
		if ok {
			f[v] += 1
		} else {
			f[v] = 1
		}
	}
	count := 0
	for k, v := range f {
		sum += k * float64(v)
		count += v
	}

	return sum / float64(count)
}

/*
portion = ΣF(x i) / n
*/
func Portion(values []float64, item float64) float64 {
	var sum float64
	for _, v := range values {
		if v == item {
			sum++
		}
	}
	return sum / float64(len(values))
}

/*
portion = ΣF(x i) / n
*/
func PortionMean(values []float64, item float64) float64 {
	var sum float64
	for _, v := range values {
		if v == item {
			sum += v
		}
	}
	return sum / float64(len(values))
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
		return (first + second) / 2
	}

	return sortedValues[len(sortedValues)/2]
}

func Mode(input []float64) []float64 {

	values := services.Sort(input)

	var mode []float64

	count, maxCount := 0, 0
	for i := 1; i < len(values); i++ {
		switch {
		case values[i] == values[i-1]:
			count++
		case count > maxCount:
			mode = append(mode[:0], values[i-1])
			maxCount = count
			count = 0
		case count == maxCount && maxCount != 0:
			mode = append(mode, values[i-1])
			count = 0
		default:
			count = 0
		}

	}
	if maxCount > 0 && count == maxCount {
		mode = append(mode, values[len(values)-1])
	} else if maxCount > 0 && count > maxCount {
		mode = append(mode[:0], values[len(values)-1])
	}

	return mode

}
