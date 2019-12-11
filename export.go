package logify

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

var curGOOS string
var mu sync.Mutex
var wg sync.WaitGroup

func init() {
	curGOOS = runtime.GOOS
}

// The EnableOutputToFile function lets us output log to file
func EnableOutputToFile() {
	baseLog.EnableToFile = true
}

// Debug log DEBUG level message.
func Debug(v interface{}) {
	outputLog(baseLog.defaultFormatLog(debugFormat, v, baseLog.CallerSkipDepth))
}

// Warn log WARN level message.
func Warn(v interface{}) {
	outputLog(baseLog.defaultFormatLog(warnFormat, v, baseLog.CallerSkipDepth))
}

// Info log INFO level message.
func Info(v interface{}) {
	outputLog(baseLog.defaultFormatLog(infoFormat, v, baseLog.CallerSkipDepth))
}

// Error log ERROR level message.
func Error(v interface{}) {
	outputLog(baseLog.defaultFormatLog(errorFormat, v, baseLog.CallerSkipDepth))
}

func outputLog(content string) {
	if baseLog.EnableToFile {
		wg.Add(1)
		if curGOOS == "windows" {
			go syncLogToFile(content)
			parsePrint(content)
			wg.Wait()
			return
		}
		fmt.Println(parseOutput(content))
		go syncLogToFile(content)
		return
	}
	if curGOOS == "windows" {
		parsePrint(content)
		return
	}
	fmt.Println(parseOutput(content))
}

func syncLogToFile(content string) {
	mu.Lock()
	f, err := os.OpenFile(time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_APPEND, 0x666)
	defer f.Close()
	if err != nil {
		fmt.Println(baseLog.defaultFormatLog(errorFormat, err, baseLog.CallerSkipDepth))
		return
	}
	if _, err = fmt.Fprintln(f, content); err != nil {
		fmt.Println(baseLog.defaultFormatLog(errorFormat, err, baseLog.CallerSkipDepth))
		return
	}
	mu.Unlock()
	wg.Done()
	return
}
