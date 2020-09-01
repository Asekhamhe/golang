package main

// Program to find the minimum and the maximum of its command-line arguments
import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")

		os.Exit(1)
	}

	arguments := os.Args
	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)

	for i := 2; i < len(arguments); i++ {
		number, _ := strconv.ParseFloat(arguments[i], 64)

		if number < min {
			min = number
		}
		if number > max {
			max = number
		}
	}

	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}
