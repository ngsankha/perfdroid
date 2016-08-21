package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type cpuUsage struct {
	utime uint64
	stime uint64
}

type CpuMetric struct {
	user   float64
	system float64
}

func totalCpuTime() uint64 {
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

func appCpuUsage(pid uint64) cpuUsage {
	// TODO: App might be killed by the time this is called, which will cause a crash
	out := ExecAdb("cat", fmt.Sprintf("/proc/%v/stat", pid))
	fields := strings.Fields(out)
	utime, err := strconv.ParseUint(fields[13], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	stime, err := strconv.ParseUint(fields[14], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return cpuUsage{utime: utime, stime: stime}
}

func CpuUsage(packageName string) CpuMetric {
	// PID is computed everytime because PID can change if app is killed or screen is locked
	pid, err := Pid(packageName)
	if err != nil {
		log.Fatal(err)
	}
	oldTotalTime := totalCpuTime()
	oldAppTime := appCpuUsage(pid)
	time.Sleep(1 * time.Second)
	newTotalTime := totalCpuTime()
	newAppTime := appCpuUsage(pid)
	user := float64(newAppTime.utime-oldAppTime.utime) / float64(newTotalTime-oldTotalTime) * 100
	system := float64(newAppTime.stime-oldAppTime.stime) / float64(newTotalTime-oldTotalTime) * 100
	return CpuMetric{user: user, system: system}
}
