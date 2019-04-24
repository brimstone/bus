package bus_test

import (
	"bus"
	"testing"
)

type request struct {
	name string
}

func subscribe(b *bus.Bus, n int) error {
	c, err := b.Subscribe(request{})
	if err != nil {
		return err
	}
	for i := 0; i < n; i++ {
		_ = <-c
	}
	return nil
}

func BenchmarkPub1(bench *testing.B) {
	// run the Fib function b.N times
	b := bus.NewBus()
	go subscribe(b, bench.N)
	for n := 0; n < bench.N; n++ {
		b.Publish(request{name: "foo"})
	}
}
