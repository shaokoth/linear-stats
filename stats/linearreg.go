package linear_stats

func LinearReg(x, y []float64) (float64, float64)  {
	sumx := 0.0
	sumy := 0.0
	sumxy := 0.0
	sumx2 := 0.0
	sumy2 := 0.0
	a := 0.0
	b := 0.0

	for i := 0; i <= len(x)-1; i++ {
		N := len(x)
		sumx += float64(x[i])
		sumy += float64(y[i])
		sumxy += float64(x[i]) * float64(y[i])
		sumx2 += float64(x[i]) * float64(x[i])
		sumy2 += float64(y[i]) * float64(y[i])
		b = (float64(N)*sumxy - sumx*sumy) / (float64(N)*sumx2 - sumx*sumx)
		a = (sumy - b*sumx) / float64(N)
	}
	return a, b
}
