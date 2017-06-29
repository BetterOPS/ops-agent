package main

import (
	"os"
	"syscall"
	"os/signal"
	"ops-agent/lib/log"
)

func InitSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for {
		c := <-ch
		log.Infof("receive signal: %s", c.String())
		switch c {
		case syscall.SIGHUP:
			return
		case syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM:
			cancel()
			return
		default:
			cancel()
			return
		}
	}
}
