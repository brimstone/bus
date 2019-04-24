package main

import "github.com/brimstone/bus"

func subscribe(b *bus.Bus) error {
	c, err := b.Subscribe(request{})
	if err != nil {
		return err
	}
	log.Info("Ready")
	m := <-c
	r := m.Payload.(request)
	log.Info("Message",
		log.Field("m", m),
		log.Field("name", r.name),
	)
	return nil
}
