package linear_stats

func LinearReg(x, y []float64) (float64, float64) {
	sumx := 0.0
	sumy := 0.0
	sumxy := 0.0
	sumx2 := 0.0
	sumy2 := 0.0
	a := 0.0
	b := 0.0

	for i := 0; i < len(x); i++ {
		N := float64(len(x))
		sumx += float64(x[i])
		sumy += float64(y[i])
		sumxy += float64(x[i]) * float64(y[i])
		sumx2 += float64(x[i]) * float64(x[i])
		sumy2 += float64(y[i]) * float64(y[i])
		a = (sumy / N) - (((N*sumxy - sumx*sumy) / (N*sumx2 - sumx*sumx)) * ((sumx) / N))
		b = (N*sumxy - sumx*sumy) / (N*sumx2 - sumx*sumx)
	}
	return a, b
}
