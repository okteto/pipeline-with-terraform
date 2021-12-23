package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
)

var (
	topic              *pubsub.Topic
	googleCloudProject = mustGetenv("GCP_PROJECT_ID")
	topicName          = mustGetenv("PUBSUB_TOPIC")
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

	for counter := 0; ; counter++ {
		publish(ctx, counter)
		time.Sleep(1 * time.Second)
	}
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("%s environment variable not set.", k)
	}
	return v
}

func publish(ctx context.Context, counter int) {

	msg := &pubsub.Message{
		Data: []byte(fmt.Sprintf("counter %d", counter)),
	}

	log.Printf("Publishing message %d...", counter)
	if _, err := topic.Publish(ctx, msg).Get(ctx); err != nil {
		log.Fatalf("Could not publish message %d: %v", counter, err)
	}

	log.Print("Done!")
}
