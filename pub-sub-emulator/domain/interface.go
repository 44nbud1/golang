package domain

import (
	"context"
	"github.com/44nbud1/golang/pub-sub-emulator/model"

	"google.golang.org/api/option"
	"log"

	googlePubSub "cloud.google.com/go/pubsub"
)

type (
	Client interface {
		Publish(ctx context.Context, request model.PubRequest)
		Subscribe(ctx context.Context, request model.SubRequest)
	}

	client struct {
		pubSubClient *googlePubSub.Client
	}

	Config struct {
		Ctx       context.Context
		ProjectID string
	}
)

func NewPubSubClient(c Config) Client {

	var err error
	var pubSubClient *googlePubSub.Client

	if pubSubClient, err = googlePubSub.NewClient(
		c.Ctx,
		c.ProjectID,
		option.WithoutAuthentication(),
	); err != nil {

		log.Fatalf("failed to construct pubSubClient, err: %v", err)
	}

	return &client{
		pubSubClient: pubSubClient,
	}
}
