package logify

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Logs struct {
	Prefix           string // prefix of each line when printing
	Level            int    // log level, above level will be exported to file
	EnableCallerSkip bool   // whether to enable print function call stack ascend
	CallerSkipDepth  int    // the depth of ascend
	EnableToFile     bool   // whether print to file
}

const (
	infoFormat        = "%s [INFO] [%s] %v"
	debugFormat       = "%s [DEBUG] [%s] %v"
	warnFormat        = "%s [WARN] [%s] %v"
	errorFormat       = "%s [ERROR] [%s] %v"
	defaultTimeFormat = "2006-01-02 15:04:05"
)

var baseLog = newLogs()

func newLogs() *Logs {
	return &Logs{
		EnableCallerSkip: true,
		CallerSkipDepth:  2,
	}
}

func (l *Logs) defaultFormatLog(format string, v interface{}, skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if ok {
		location := parseFilename(file) + ":" + strconv.FormatInt(int64(line), 10)
		return fmt.Sprintf(format, time.Now().Format(defaultTimeFormat), location, v)
	}
	return fmt.Sprintf(parseFormat(format), time.Now().Format(defaultTimeFormat), v)
}

func parseFilename(filename string) string {
	str := strings.Split(filename, "/")
	if len(str) > 1 {
		return str[len(str)-1]
	}
	return str[0]
}

func parseFormat(format string) string {
	str := strings.Split(format, " ")
	return str[0] + " " + str[1] + " " + str[3]
}
