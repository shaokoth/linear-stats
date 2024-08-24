
# Overview

This program reads a list of float64s from a file and calculates the linear regression model as well as the pearson correlation coefficient. It then prints these statistical values to the console.
Features:

* Reads float64 values from a file of type string
* Calculates the linear regression model
* Calculates the pearson correlation coefficient

# Prerequisites

To run this program make sure you have GO 1.16 or above installed on your system. You can download and install Go from golang.org.
Clone this repository or download source code. https://learn.zone01kisumu.ke/git/shaokoth/linear-stats.git

# Usage

Ensure you have a file (e.g., data.txt) in the same directory as the program. This file should contain the population data whose statistics we are interested in.

Run the program with the default data.txt file: 
```
go run . main.go data.txt
```

# Code Structure

* main.go: The main file containing the code logic. readfile.go: function where the file is open and read. linearreg.go: calculates the linear regression of the population data. pearsoncorr.go: calculates the pearson correlation coefficient. 

# Error Handling

The program includes basic error handling for:

* Invalid file names or paths.
* Errors during file reading.
* Unsupported characters not defined in the data file.
* Empty files

# Other Innformation

This project is maintained by [ShadrackOkoth](https://github.com/shaokoth)

# License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.