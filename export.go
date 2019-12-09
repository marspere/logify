package logify

import "fmt"

// Debug log DEBUG level message.
func Debug(v interface{}) {
	fmt.Println(parseOutput(baseLog.defaultFormatLog(debugFormat, v, baseLog.CallerSkipDepth)))
}

// Warn log WARN level message.
func Warn(v interface{}) {
	fmt.Println(parseOutput(baseLog.defaultFormatLog(warnFormat, v, baseLog.CallerSkipDepth)))
}

// Info log INFO level message.
func Info(v interface{}) {
	fmt.Println(parseOutput(baseLog.defaultFormatLog(infoFormat, v, baseLog.CallerSkipDepth)))
}

// Error log ERROR level message.
func Error(v interface{}) {
	fmt.Println(parseOutput(baseLog.defaultFormatLog(errorFormat, v, baseLog.CallerSkipDepth)))
}
