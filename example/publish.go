package main

import "github.com/brimstone/bus"

func publish(b *bus.Bus) {
	b.Publish(request{name: "foo"})
}
