package main

import (
	"io"
	"os"
)

func main() {
	/*
		get user input by reading the command-line arguments of a program
	*/
	myString := ""
	arguments := os.Args

	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}
