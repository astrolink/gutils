package log

import (
	"fmt"
	ghttp "github.com/astrolink/gutils/http"
	gtime "github.com/astrolink/gutils/time"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
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

func GetRequestRealIp(r *http.Request) string {
	var ip string

	if r == nil {
		log.Println(fmt.Sprintf(EmptyRequestObjectErrorMessage, "ip"))
		return ip
	}

	ip = r.Header.Get(ghttp.CustomRealIpHeaderKey)

	if ip == "" {
		forwardedFor := r.Header.Get(ghttp.CustomForwardedForKey)
		ips := strings.Split(forwardedFor, ",")

		if len(ips) > 0 {
			ip = ips[0]
		}

		if ip == "" {
			ip = r.RemoteAddr
		}
	}

	return ip
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
