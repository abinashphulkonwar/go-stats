package descriptive

import (
	"github.com/abinashphulkonwar/stats/src/services"
)

type Quantiles struct {
	Q1 float64
	Q2 float64
	Q3 float64
	Q4 float64
}

func Position(p int, l int) int {
	return int((float64(p) / float64(100)) * float64(l-1))
}

func Quantile(inputs []float64) Quantiles {
	values := services.Sort(inputs)
	l := len(values)
	q1 := values[Position(25, l)]
	q2 := values[Position(50, l)]
	q3 := values[Position(75, l)]
	q4 := values[Position(100, l)]

	return Quantiles{
		Q1: q1,
		Q2: q2,
		Q3: q3,
		Q4: q4,
	}

}

func Detailed(input []float64) []float64 {

	values := services.Sort(input)
	l := len(values)
	detailds := []float64{}
	for i := 10; i <= 100; i = i + 10 {
		detailds = append(detailds, values[Position(i, l)])
	}

	return detailds
}
