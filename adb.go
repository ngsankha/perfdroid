package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

var adbPath string
var appPid uint64

func SetupAdb(packageName string) error {
	path, err := exec.LookPath("adb")
	if err != nil {
		return err
	}
	adbPath = path
	pid, err := queryPid(packageName)
	if err != nil {
		return err
	}
	appPid = pid
	return nil
}

func AdbPath() string {
	return adbPath
}

func Pid() uint64 {
	return appPid
}

func ExecAdb(arg ...string) string {
	cmd := exec.Command(AdbPath(), append([]string{"shell"}, arg...)...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	readBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
	}
	return string(readBytes)
}

func queryPid(packageName string) (uint64, error) {
	out := ExecAdb("ps")
	lines := strings.Split(out, "\n")
	for _, line := range lines {
		if strings.Contains(line, packageName) {
			fields := strings.Fields(line)
			val, err := strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			return val, nil
		}
	}
	return 0, fmt.Errorf("Target app (%s) is not running", packageName)
}
