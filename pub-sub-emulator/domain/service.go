package domain

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"github.com/44nbud1/golang/pub-sub-emulator/model"
	"log"
)

func (c *client) Publish(ctx context.Context, request model.PubRequest) {

	var err error
	var isExist bool

	// create topic
	t := c.pubSubClient.Topic(request.Topic)

	isExist, err = t.Exists(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if !isExist {
		log.Printf("Topic %v doesn't exist - creating it\n", request.Topic)

		_, err = c.pubSubClient.CreateTopic(ctx, request.Topic)
		if err != nil {
			log.Fatalf("Create topic, err:%v", err)
		}

		_, err = c.pubSubClient.CreateTopic(ctx, fmt.Sprintf("%s%s", request.Topic, "-dead-letter"))
		if err != nil {
			log.Fatalf("Create topic dead letter, err:%v", err)
		}
	}

	result := t.Publish(ctx, &pubsub.Message{
		Data: request.DataToByteArray(),
	})

	var sID string
	if sID, err = result.Get(ctx); err != nil {
		log.Fatalf("Publish, result.Get, err:%v", err)

		return
	}

	log.Printf("Published a message %v", sID)
}

func (c *client) Subscribe(ctx context.Context, request model.SubRequest) {

	s := c.pubSubClient.Subscription(request.Subs)

	exists, err := s.Exists(ctx)
	if err != nil {
		log.Fatal(err)
	}

	t := c.pubSubClient.Topic(request.Topic)

	if !exists {
		topic := c.pubSubClient.Topic(request.Topic)

		subConfig := pubsub.SubscriptionConfig{
			Topic: topic,
			DeadLetterPolicy: &pubsub.DeadLetterPolicy{
				DeadLetterTopic:     fmt.Sprintf("%v%v", t.String(), "-dead-letter"),
				MaxDeliveryAttempts: 5,
			},
		}

		sub, err := c.pubSubClient.CreateSubscription(ctx, request.Subs, subConfig)

		if err != nil {
			log.Fatalf("CreateSubscription: %v", err)
		}

		log.Printf("Success to create subs %v", sub)

	}

	if err = s.Receive(ctx, func(c context.Context, message *pubsub.Message) {

		request.Data = message.Data

		log.Printf("Message received %s", string(message.Data))

		if message.DeliveryAttempt != nil {
			log.Printf("max attempt %v", *message.DeliveryAttempt)
		}

		if request.IsNack {
			message.Nack()
		} else {
			message.Ack()
		}

	}); err != nil {

	}
}
