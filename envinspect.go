package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		printProcessData()
		time.Sleep(5 * time.Second)
	}
}

func printProcessData() {
	uid := os.Getuid()
	gid := os.Getgid()

	fmt.Printf("User data: \nUID: %v\nGID: %v\n", uid, gid)

	pid := os.Getpid()
	parentPid := os.Getppid()
	fmt.Printf("\nProcess info:\nPID: %v\nParent PID: %v\n", pid, parentPid)
	workdir, _ := os.Getwd()
	fmt.Printf("Working Dir: %v\n", workdir)

	fmt.Printf("\nEnvironment:\n")
	env := os.Environ()
	for _, v := range env {
		fmt.Printf("%v;  ", v)
	}

	fmt.Printf("\n\n\n")
}
