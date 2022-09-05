package htapplife

import (
	"fmt"
	colog "github.com/hamidteimouri/htutils/htcolog"
	envier "github.com/hamidteimouri/htutils/htenvier"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	HTTimeZone   = "Asia/Tehran"
	HTTimeFormat = "2006-01-02 15:04:05"
)

type HTInitiator struct {
	t time.Time
}

func Start() HTInitiator {
	t := time.Now()
	tz := envier.EnvOrDefault("HT_TIMEZONE", HTTimeZone)
	loc, err := time.LoadLocation(tz)
	if err != nil {
		loc, _ = time.LoadLocation(HTTimeZone)
		tz = HTTimeZone
	}
	colog.DoPurple("***************************")
	colog.DoPurple("* Application is starting *")
	colog.DoPurple("***************************")
	n := time.Now().In(loc).Format(HTTimeFormat)
	startedTime := fmt.Sprintf(" - Started at : %v (%v)", n, tz)
	colog.DoGray(startedTime)
	colog.DoGreen(" - APP_ENV is : " + envier.GetAppEnv())
	return HTInitiator{t}
}

func (ci HTInitiator) Stop() {
	endtime := fmt.Sprintf(" - Uptime : %v", time.Since(ci.t).String())
	colog.DoGray(endtime)
	colog.DoGray(" - The closing process is over")
	colog.DoPurple("********* Stopped *********")
}

func ClosingSignalListener() (isForceToStop bool) {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	switch <-ch {
	case os.Interrupt:
		return true
	case syscall.SIGTERM:
		return false
	default:
		return false
	}
}
