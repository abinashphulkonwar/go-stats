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

func Quantile(inputs []float64) Quantiles {
	values := services.Sort(inputs)
	l := len(values)
	q1 := values[services.Position(25, l)]
	q2 := values[services.Position(50, l)]
	q3 := values[services.Position(75, l)]
	q4 := values[services.Position(100, l)]

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
		detailds = append(detailds, values[services.Position(i, l)])
	}

	return detailds
}

func Percentile(input []float64, percent float64) float64 {

	values := services.Sort(input)
	i := services.PositionF(percent, len(input))
	if percent == float64(int(percent)) {
		return values[i]
	} else {
		return (values[i-1] + values[i]) / 2

	}
}
