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
	for {
		usage := CpuUsage(*packageName)
		fmt.Println("User:", usage.user)
		fmt.Println("System:", usage.system)
	}

}

func parseArgs() {
	packageName = flag.String("package", "", "package to monitor")
	flag.Parse()
}
