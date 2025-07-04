package linearclassifier

func binaryScore(w []float32, f []float32) float32 {

	if len(w) == 0 || len(f) == 0 || len(w) != len(f) {
		panic("Invalid input: weights and features must be non-empty and of the same length")
	}
	sum := float32(0.0)
	for i := 0; i < len(w); i++ {
		W_i := w[i]
		F_i := f[i]
		sum = sum + (W_i * F_i)
	}
	return sum
}

func BinaryClassify(w []float32, f []float32) int {
	score := binaryScore(w, f)
	if score > 0 {
		return 1
	} else if score < 0 {
		return -1
	} else {
		return 0
	}
}
