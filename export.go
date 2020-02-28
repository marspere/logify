package logify

import (
	"fmt"
	"runtime"
	"sync"
)

var curGOOS string
var mu sync.Mutex
var wg sync.WaitGroup

func init() {
	curGOOS = runtime.GOOS
}

// Debug log DEBUG level message.
func Debug(v ...interface{}) {
	outputLog(baseLog.defaultFormatLog(debugFormat, baseLog.CallerSkipDepth, v...))
}

// Warn log WARN level message.
func Warn(v ...interface{}) {
	outputLog(baseLog.defaultFormatLog(warnFormat, baseLog.CallerSkipDepth, v...))
}

// Info log INFO level message.
func Info(v ...interface{}) {
	outputLog(baseLog.defaultFormatLog(infoFormat, baseLog.CallerSkipDepth, v...))
}

// Error log ERROR level message.
func Error(v ...interface{}) {
	outputLog(baseLog.defaultFormatLog(errorFormat, baseLog.CallerSkipDepth, v...))
}

func outputLog(content string) {
	if baseLog.EnableToFile {
		wg.Add(1)
		defer wg.Done()
		if curGOOS == "windows" {
			go syncLogToFile(content)
			parsePrint(content)
			wg.Wait()
			return
		}
		fmt.Println(parseOutput(content))
		go syncLogToFile(content)
		wg.Wait()
		return
	}
	if curGOOS == "windows" {
		parsePrint(content)
		return
	}
	fmt.Println(parseOutput(content))
}
