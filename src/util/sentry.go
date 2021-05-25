package util

import (
	"log"
	"time"

	"github.com/getsentry/sentry-go"
)

func Sentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "http://1fd9eb78958b40db880b283523750f49@192.168.1.10:9000/4",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("It works!")
}
