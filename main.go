package main

import (
	"context"

	"github.com/kurtpeek/my-cloud-function/helloworld"
)

func main() {
	helloworld.HelloPubSub(context.Background(), helloworld.PubSubMessage{
		Data: []byte("foobar"),
	})
}
