package log

import (
	"fmt"
	gtime "github.com/astrolink/gutils/time"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const EmptyRequestObjectErrorMessage = "no request object was set to get %s"

func getCallerFullInfo() string {
	var info string
	now := time.Now().Format(gtime.DateTimeFormat)
	pc, file, line, ok := runtime.Caller(1)

	if !ok {
		return info
	}

	info = fmt.Sprintf("[(%s) %s:#%d method: %v]", now, file, line, runtime.FuncForPC(pc).Name())
	return info
}

func getBinName() string {
	binName := os.Args[0]
	binName = filepath.Base(binName)
	return binName
}
