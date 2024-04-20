package main

import (
	"context"
	"github.com/44nbud1/golang/pub-sub-emulator/domain"
	"github.com/44nbud1/golang/pub-sub-emulator/model"
)

func main() {

	c := domain.NewPubSubClient(domain.Config{
		Ctx:       context.Background(),
		ProjectID: "pub-sub-id-test-3",
	})

	c.Publish(context.Background(), model.PubRequest{
		Topic: "create-message",
		Data: map[string]interface{}{
			"name": "aan",
			"age":  29,
		},
	})

	c.Subscribe(context.Background(), model.SubRequest{
		Topic:  "create-message",
		Subs:   "create-message-subs",
		IsNack: true,
	})
}
