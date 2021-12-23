package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/pubsub"
)

var (
	topic              *pubsub.Topic
	googleCloudProject = mustGetenv("GCP_PROJECT_ID")
	topicName          = mustGetenv("PUBSUB_TOPIC")
	subName            = fmt.Sprintf("sub-%s", topicName)
)

func main() {
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, googleCloudProject)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	topic = client.Topic(topicName)

	// Check topic already exists
	exists, err := topic.Exists(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if !exists {
		log.Fatalf("Topic %s doesn't exist", topicName)
	}

	sub, err := client.CreateSubscription(ctx, subName, pubsub.SubscriptionConfig{Topic: topic})
	if err != nil {
		if !strings.Contains(err.Error(), "AlreadyExists") {
			log.Fatalf("error creating subscription: %s", err.Error())
		}
		sub = client.SubscriptionInProject(subName, googleCloudProject)
	}

	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		log.Printf("Received %s", m.Data)
		m.Ack()
	})
	if err != nil {
		log.Fatalf("error reading message: %s", err.Error())
	}
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("%s environment variable not set.", k)
	}
	return v
}
