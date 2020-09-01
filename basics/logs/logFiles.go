package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	programName := filepath.Base(os.Args[0])
	syslog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL0, programName)

	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(syslog)
	}

	log.Println("LOG_INFO + LOG_LOCAL0: Logging in Go!")
	fmt.Println("Logged sucessfully")
}
