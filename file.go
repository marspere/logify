package logify

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	ticker := time.NewTicker(24 * 60 * 60 * time.Second)
	go func() {
		for t := range ticker.C {
			if baseLog.EnableToFile {
				err := os.Remove(baseLog.Location + time.Now().AddDate(0, 0, -baseLog.MaxSaveDays).Format("2006-01-02") + ".log")
				if err != nil && err.Error() != "The system cannot find the file specified." {
					fmt.Println(baseLog.defaultFormatLog(errorFormat, baseLog.CallerSkipDepth, err))
				}
				fmt.Println(baseLog.defaultFormatLog(debugFormat, baseLog.CallerSkipDepth, strconv.FormatInt(t.Unix(), 10)+":delete log"))
			}
		}
	}()
}

// The EnableOutputToFile function lets us output log to file
func EnableOutputToFile() {
	baseLog.EnableToFile = true
}

// The SetMaxSaveDays function lets us set maximum days of saving logs.
func SetMaxSaveDays(n int) {
	if n < 0 {
		return
	}
	baseLog.MaxSaveDays = n
}

// The SetLogLocation function lets us set the location of saving logs.
// The default save path is the current directory.
func SetLogLocation(loc string) {
	if loc == "" {
		baseLog.Location = "./"
		return
	}
	if !strings.HasSuffix(loc, "/") {
		baseLog.Location = loc + "/"
		return
	}
	baseLog.Location = loc
}

func syncLogToFile(content string) {
	mu.Lock()
	f, err := os.OpenFile(baseLog.Location+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_APPEND, 0x666)
	defer f.Close()
	if err != nil {
		fmt.Println(baseLog.defaultFormatLog(errorFormat, baseLog.CallerSkipDepth, err))
		return
	}
	if _, err = fmt.Fprintln(f, content); err != nil {
		fmt.Println(baseLog.defaultFormatLog(errorFormat, baseLog.CallerSkipDepth, err))
		return
	}
	mu.Unlock()
	wg.Done()
	return
}
