package main

import (
	"./radius-go"

	"fmt"
	"os"
)


var progName = "main"

func main() {
	err := radius_go.ServerStartUp()
	if err != nil{
		fmt.Fprintf(os.Stderr, "[%s][%s]\t%s\n", progName, "ERROR", err)
		os.Exit(1)
	}
	defer fmt.Fprintf(os.Stdout, "[%s][%s]\t%s\n", progName, "NORMAL", "Server exit")
}

