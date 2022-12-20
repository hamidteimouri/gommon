package htapplife

import (
	"fmt"
	colog "github.com/hamidteimouri/gommon/htcolog"
	envier "github.com/hamidteimouri/gommon/htenvier"
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
