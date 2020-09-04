package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please enter one or more numbers.")
		os.Exit(1)
	}

	var err error = errors.New("an error")
	var number float64

	arguments := os.Args
	index := 1

	for err != nil {
		if index >= len(arguments) {
			fmt.Println("None of the arguments is a float!")
			return
		}
		number, err = strconv.ParseFloat(arguments[index], 64)
		index++
	}

	min, max := number, number

	for i := 2; i < len(arguments); i++ {
		number, err := strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			if number < min {
				min = number
			}
			if number > max {
				max = number
			}
		}
	}

	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)

}
