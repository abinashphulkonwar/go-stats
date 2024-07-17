package services

func Sort(input []float64) []float64 {
	values := make([]float64, len(input))
	copy(values, input)

	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				temp := values[i]
				values[i] = values[j]
				values[j] = temp
			}
		}
	}
	return values
}
