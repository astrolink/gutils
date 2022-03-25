package log

import (
	"fmt"
	gtime "github.com/astrolink/gutils/time"
	"log"
	"net/http"
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

func GetUserAgent(r *http.Request) string {
	var agent string

	if r == nil {
		log.Println(fmt.Sprintf(EmptyRequestObjectErrorMessage, "user agent"))
		return agent
	}

	return r.UserAgent()
}

func GetCurrentRoute(r *http.Request) string {
	var route string

	if r == nil {
		log.Println(fmt.Sprintf(EmptyRequestObjectErrorMessage, "current route"))
		return route
	}

	route = fmt.Sprintf("%s/%s", r.Host, r.URL.Path)
	return route
}
