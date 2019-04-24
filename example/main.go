package main

import (
	"time"

	"github.com/brimstone/bus"
	"github.com/brimstone/logger"
)

var (
	log = logger.New()
)

type request struct {
	name string
}
type response struct{}

func main() {

	b := bus.NewBus()

	go subscribe(b)

	time.Sleep(time.Second)
	log.Info("Publishing")
	publish(b)
	time.Sleep(time.Second)
	log.Info("Done")
}
