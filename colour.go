package logify

import (
	"fmt"
	"strings"
)

const (
	errorColour = "%c[1;40;31m%s%c[0m"
	debugColour = "%c[1;40;32m%s%c[0m"
	warnColour  = "%c[1;40;33m%s%c[0m"
	infoColour  = "%c[1;40;34m%s%c[0m"
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

