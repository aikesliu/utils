package signals

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/aikesliu/utils/log"
)

func WaitSignals(signals ...os.Signal) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, signals...)
	sig := <-c
	log.I("stop, rcv sig: %v", sig)
}

func DefaultWait() {
	WaitSignals(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGABRT)
}
