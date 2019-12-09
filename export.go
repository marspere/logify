package logify

import (
	"fmt"
	"runtime"
)

var curGOOS string

func init() {
	curGOOS = runtime.GOOS
}

// Debug log DEBUG level message.
func Debug(v interface{}) {
	if curGOOS == "windows" {
		parsePrint(baseLog.defaultFormatLog(debugFormat, v, baseLog.CallerSkipDepth))
		return
	}
	fmt.Println(parseOutput(baseLog.defaultFormatLog(debugFormat, v, baseLog.CallerSkipDepth)))
}

// Warn log WARN level message.
func Warn(v interface{}) {
	if curGOOS == "windows" {
		parsePrint(baseLog.defaultFormatLog(warnFormat, v, baseLog.CallerSkipDepth))
		return
	}
	fmt.Println(parseOutput(baseLog.defaultFormatLog(warnFormat, v, baseLog.CallerSkipDepth)))
}

// Info log INFO level message.
func Info(v interface{}) {
	if curGOOS == "windows" {
		parsePrint(baseLog.defaultFormatLog(infoFormat, v, baseLog.CallerSkipDepth))
		return
	}
	fmt.Println(parseOutput(baseLog.defaultFormatLog(infoFormat, v, baseLog.CallerSkipDepth)))
}

// Error log ERROR level message.
func Error(v interface{}) {
	if curGOOS == "windows" {
		parsePrint(baseLog.defaultFormatLog(errorFormat, v, baseLog.CallerSkipDepth))
		return
	}
	fmt.Println(parseOutput(baseLog.defaultFormatLog(errorFormat, v, baseLog.CallerSkipDepth)))
}
