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
	MaxSaveDays      int    // the maximum days of saving log
	Location         string // where the log files are saved
}

const (
	infoFormat         = "%s [INFO] [%s]"
	debugFormat        = "%s [DEBUG] [%s]"
	warnFormat         = "%s [WARN] [%s]"
	errorFormat        = "%s [ERROR] [%s]"
	defaultTimeFormat  = "2006-01-02 15:04:05"
	defaultMaxSaveDays = 7
)

var baseLog = newLogs()

func newLogs() *Logs {
	return &Logs{
		EnableCallerSkip: true,
		CallerSkipDepth:  2,
		MaxSaveDays:      defaultMaxSaveDays,
	}
}

func (l *Logs) defaultFormatLog(format string, skip int, v ...interface{}) string {
	_, file, line, ok := runtime.Caller(skip)
	v1 := make([]interface{}, 0)
	v1 = append(v1, time.Now().Format(defaultTimeFormat))
	if ok {
		v1 = append(v1, parseFilename(file)+":"+strconv.FormatInt(int64(line), 10))
	}
	v1 = append(v1, v...)
	return fmt.Sprintf(assembleFormat(format, ok, v...), v1...)
}

func parseFilename(filename string) string {
	str := strings.Split(filename, "/")
	if len(str) > 1 {
		return str[len(str)-1]
	}
	return str[0]
}

func assembleFormat(format string, ok bool, v ...interface{}) string {
	if !ok {
		format = strings.TrimSuffix(format, " [%s]")
	}
	for i := 0; i < len(v); i++ {
		format = format + " " + "%v"
	}
	return format
}

func SetLogCallDepth(n int) {
	baseLog.CallerSkipDepth = n
}
