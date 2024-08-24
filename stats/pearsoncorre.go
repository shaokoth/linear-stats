package linear_stats

import (
	"math"
)

func PearsonCorr(x, y []float64) float64 {
	pearsoncorr := 0.0
	sumx := 0.0
	sumy := 0.0
	sumxy := 0.0
	sumx2 := 0.0
	sumy2 := 0.0
	den := 0.0
	num := 0.0

	for i := 0; i < len(x); i++ {
		N := float64(len(x))
		sumx += float64(x[i])
		sumy += float64(y[i])
		sumxy += float64(x[i]) * float64(y[i])
		sumx2 += float64(x[i]) * float64(x[i])
		sumy2 += float64(y[i]) * float64(y[i])
		// Calculates the numerator of the pearson correlation coefficient formula
		num = (float64(N)*sumxy - sumx*sumy)
		// Calculates the denominator of the pearson correlation coefficient formula
		den = math.Sqrt((N*sumx2 - sumx*sumx) * (N*sumy2 - sumy*sumy))
		pearsoncorr = num / den
	}
	return pearsoncorr
}
