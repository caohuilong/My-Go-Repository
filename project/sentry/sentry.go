package main

import (
	"github.com/getsentry/sentry-go"
	"log"
	"time"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:   "https://0e5c2f67492f4143a97af2123251d423@axe-sentry.ctcdn.cn/10",
		Debug: true,
	})
	if err != nil {
		log.Fatalf("sentry init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("It works!")
}
