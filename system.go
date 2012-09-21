package go_conf

import (
	"os"
	"os/signal"
	"syscall"
)

func signalCatcher() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP)
	for signal := range ch {
		if signal == syscall.SIGHUP {
			log.Println("received SIGHUP exiting...")
			os.Exit(0)
		}
	}
}

func startSignalCatcher() {
	//react to sighup
	go signalCatcher()
}
