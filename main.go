package main

import (
	"flag"
	"fmt"
	"log"
)

var packageName *string
var timeInterval *int

func main() {
	parseArgs()

	err := SetupAdb(*packageName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Package:", *packageName)
	for {
		usage := CpuUsage(*packageName, *timeInterval)
		fmt.Println("User:", usage.user)
		fmt.Println("System:", usage.system)
	}

}

func parseArgs() {
	packageName = flag.String("package", "", "package to monitor")
	timeInterval = flag.Int("timeInterval", 1, "Time interval in which to gather data (in seconds)")
	flag.Parse()
}
