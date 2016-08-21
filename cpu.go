package main

import (
	"log"
	"strconv"
	"strings"
)

type cpuUsage struct {
	utime uint64
	stime uint64
}

func TotalCpuTime() uint64 {
	out := ExecAdb("cat", "/proc/stat")
	lines := strings.Split(out, "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			var sum uint64 = 0
			arr := fields[1:]
			for _, num := range arr {
				val, err := strconv.ParseUint(num, 10, 64)
				if err != nil {
					log.Fatal(err)
				}
				sum += val
			}
			return sum
		}
	}
	log.Fatal("No CPU fields found")
	return 0
}
