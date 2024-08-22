package main

import (
	"fmt"
	"os"
	"strings"

	linear_stats "linearstats/stats"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}
	filename := os.Args[1]
	if filename == "" {
		return
	}
	if !strings.HasSuffix(filename, ".txt") {
		fmt.Println("Only .txt files allowed")
		return
	}
	x, y := linear_stats.Readfile(filename)
	pearsoncoefficient := linear_stats.PearsonCorr(x, y)
	xval, yval := linear_stats.Readfile(filename)
	a, b := linear_stats.LinearReg(xval, yval)
	fmt.Printf("%.10f\n", pearsoncoefficient)
	fmt.Printf("y = %.6f + %.6fx", a, b)
}
