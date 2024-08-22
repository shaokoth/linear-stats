package linear_stats

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
// Reads the file input 
func Readfile(filename string) ([]float64, []float64)  {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	defer file.Close()
	var yvalues []float64
	var xvalues []float64
	hasContent := false
// Scans the file line by line
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		hasContent = true
		line := scanner.Text()
		values := strings.Split(line, "\n")
		if line == "" {
			continue
		}
		// Converts the string to float64s
		number, err := strconv.ParseFloat(values[0.0], 64)
		if err != nil {
			fmt.Println("Invalid value format in line :",count, err)
			os.Exit(1)
		} else {
			yvalues = append(yvalues, number)
			xvalues = append(xvalues, float64(count))
		}
		count++

	}
	if len(yvalues) == 0 {
		fmt.Println("The file isempty")
			os.Exit(1)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	if !hasContent {
		fmt.Println("Error: The file is empty")
		os.Exit(1)
	}
	return xvalues, yvalues
}
