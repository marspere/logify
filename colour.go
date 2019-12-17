package logify

import (
	"fmt"
	"strings"
	"syscall"
)

const (
	errorColour = "%c[1;40;31m%s%c[0m"
	debugColour = "%c[1;40;32m%s%c[0m"
	warnColour  = "%c[1;40;33m%s%c[0m"
	infoColour  = "%c[1;40;34m%s%c[0m"
)

const (
	infoColor  = 1 | 1
	warnColor  = 2 | 8
	debugColor = 3 | 1
	errorColor = 4 | 4
)

func parseOutput(output string) string {
	if strings.Contains(output, "[ERROR]") {
		str := strings.Split(output, "[ERROR]")
		return str[0] + fmt.Sprintf(errorColour, 0x1B, "[ERROR]", 0x1B) + str[1]
	}
	if strings.Contains(output, "[WARN]") {
		str := strings.Split(output, "[WARN]")
		return str[0] + fmt.Sprintf(warnColour, 0x1B, "[WARN]", 0x1B) + str[1]
	}
	if strings.Contains(output, "[DEBUG]") {
		str := strings.Split(output, "[DEBUG]")
		return str[0] + fmt.Sprintf(debugColour, 0x1B, "[DEBUG]", 0x1B) + str[1]
	}
	str := strings.Split(output, "[INFO]")
	return str[0] + fmt.Sprintf(infoColour, 0x1B, "[INFO]", 0x1B) + str[1]
}

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
	handle, _, err := proc.Call(uintptr(syscall.Stdout), uintptr(7))
	if err != nil {
		fmt.Print(str[1] + "\n")
		return
	}
	_, _, err = kernel32.NewProc("CloseHandle").Call(handle)
	if err != nil {
		fmt.Print(str[1] + "\n")
		return
	}
	fmt.Print(str[1] + "\n")
}
