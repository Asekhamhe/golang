package main

// The os package provides a way to access the command-line arguments of a Go program (os.Args)
// Performs OS functions
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println("$", scanner.Text())
	}

}
