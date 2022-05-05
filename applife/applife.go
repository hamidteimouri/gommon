package applife

import (
	"fmt"
	"github.com/hamidteimouri/htutils/colog"
	"github.com/hamidteimouri/htutils/envier"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	TimeZone = "Asia/Tehran"
)

type CrpInitiator struct {
	t time.Time
}

func Start() CrpInitiator {
	t := time.Now()
	tz := envier.EnvOrDefault("HT_TIMEZONE", TimeZone)
	loc, err := time.LoadLocation(tz)
	if err != nil {
		loc, _ = time.LoadLocation(TimeZone)
		tz = TimeZone
	}
	colog.DoPurple("***************************")
	colog.DoPurple("* Application is starting *")
	colog.DoPurple("***************************")
	n := time.Now().In(loc).Format("2006-01-02 15:04:05")
	startedTime := fmt.Sprintf(" - Started at : %v (%v)", n, tz)
	colog.DoGray(startedTime)
	colog.DoGreen("- APP_ENV is :" + envier.GetAppEnv())
	return CrpInitiator{t}
}

func (ci CrpInitiator) Stop() {
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
