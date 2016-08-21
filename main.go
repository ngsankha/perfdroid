package main

import (
	"flag"
	"fmt"
	"log"
)

var packageName *string

func main() {
	parseArgs()

	err := SetupAdb(*packageName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Package:", *packageName)
	fmt.Println("PID:", Pid())
	fmt.Println("Total CPU Time:", TotalCpuTime())
}

func parseArgs() {
	packageName = flag.String("package", "", "package to monitor")
	flag.Parse()
}
