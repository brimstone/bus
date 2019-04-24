package bus

import (
	"reflect"

	"github.com/brimstone/logger"
)

type Message struct {
	Payload interface{}
}
type Bus struct {
	subscriptions []*Subscription
}

var (
	log = logger.New()
)

func NewBus() *Bus {
	return &Bus{}
}

func (b *Bus) Subscribe(t interface{}) (chan Message, error) {
	c := make(chan Message)
	s := &Subscription{
		Ch:   c,
		Type: t,
	}
	b.subscriptions = append(b.subscriptions, s)
	return c, nil
}

func (b *Bus) Publish(p interface{}) {
	m := Message{
		Payload: p,
	}
	for i := range b.subscriptions {
		/*
			log.Debug("Comparing",
				log.Field("i", reflect.TypeOf(b.subscriptions[i].Type)),
				log.Field("p", reflect.TypeOf(p)),
			)
		*/
		if reflect.TypeOf(b.subscriptions[i].Type) == reflect.TypeOf(p) {
			//log.Debug("Found a match")
			b.subscriptions[i].Ch <- m
		}
	}
}

type Subscription struct {
	Ch   chan Message
	Type interface{}
}
