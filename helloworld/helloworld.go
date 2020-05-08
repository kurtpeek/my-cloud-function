package helloworld

import (
	"context"
	"log"

	"github.com/kurtpeek/my-private-repo/mypackage"
)

// PubSubMessage is the payload of a Pub/Sub event.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// HelloPubSub2 consumes a Pub/Sub message.
func HelloPubSub2(ctx context.Context, m PubSubMessage) error {
	name := string(m.Data)
	if name == "" {
		name = "World"
	}
	log.Printf("Hello, %s!", name)
	log.Println(mypackage.SayHello())
	return nil
}
