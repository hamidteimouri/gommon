package gommon

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitForExit() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-ch
}
