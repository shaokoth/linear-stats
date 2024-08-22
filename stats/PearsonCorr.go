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

	for i := 0; i <= len(x)-1; i++ {
		n := len(x)
		sumx += float64(x[i])
		sumy += float64(y[i])
		sumxy += float64(x[i]) * float64(y[i])
		sumx2 += float64(x[i]) * float64(x[i])
		sumy2 += float64(y[i]) * float64(y[i])
		num = (float64(n)*sumxy - sumx*sumy)
		den = math.Sqrt((float64(n)*sumx2 - sumx*sumx) * (float64(n)*sumy2 - sumy*sumy))
		pearsoncorr = num / den
	}
	return pearsoncorr
}
