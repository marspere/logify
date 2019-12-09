package logify

import (
	"fmt"
	"log"
	"strings"
	"syscall"
)

const (
	infoColor  = 1 | 1
	warnColor  = 2 | 8
	debugColor = 3 | 1
	errorColor = 4 | 4
)

func parsePrint(output string) {
	if strings.Contains(output, "[ERROR]") {
		callWindowsPrint(output, "[ERROR]", errorColor)
		return
	}
	if strings.Contains(output, "[WARN]") {
		callWindowsPrint(output, "[WARN]", warnColor)
		return
	}
	if strings.Contains(output, "[DEBUG]") {
		callWindowsPrint(output, "[DEBUG]", debugColor)
		return
	}
	callWindowsPrint(output, "[INFO]", infoColor)
}

func callWindowsPrint(output, level string, colour int) {
	str := strings.Split(output, level)
	fmt.Print(str[0])
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("SetConsoleTextAttribute")
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(colour))
	fmt.Print(level)
	handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
	CloseHandle := kernel32.NewProc("CloseHandle")
	_, _, err := CloseHandle.Call(handle)
	log.Println("callWindowsPrint error:", err)
	fmt.Print(str[1] + "\n")
}
