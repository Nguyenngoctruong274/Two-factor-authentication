package main

import (
	"os"
	"os/signal"
	"strings"
	"time"
)

const (
	ReadTimeoutSecond  time.Duration = 30
	WriteTimeoutSecond               = 30
	IdleTimeoutSecond                = 60
)

func listenSignal() chan os.Signal {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	return s
}

func bindAddressPretty(s string) string {
	if !strings.Contains(s, ":") {
		s = "0.0.0.0:" + s
	}
	bindArr := strings.Split(s, ":")
	if !strings.Contains(bindArr[0], ".") { // include len(bindArr[0]) == 0
		bindArr[0] = "0.0.0.0"
	}
	if len(bindArr[1]) == 0 {
		bindArr[1] = "8080"
	}
	return strings.Join(bindArr, ":")
}
