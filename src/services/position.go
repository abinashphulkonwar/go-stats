package services

func Position(p int, l int) int {
	return int((float64(p) / float64(100)) * float64(l-1))
}
func PositionF(p float64, l int) int {
	return int((p / 100) * float64(l-1))
}
