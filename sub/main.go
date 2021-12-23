package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"cloud.google.com/go/pubsub"
)

var (
	topic *pubsub.Topic

	// Messages received by this instance.
	messagesMu sync.Mutex
	messages   []string

	googleCloudProject = mustGetenv("GCP_PROJECT_ID")
	topicName          = mustGetenv("PUBSUB_TOPIC")
)

const maxMessages = 10

func main() {
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, googleCloudProject)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	topic = client.Topic(topicName)

	// Create the topic if it doesn't exist.
	exists, err := topic.Exists(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if !exists {
		log.Fatal("Topic %v doesn't exist")
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
		Data: []byte(fmt.Sprintf("Counter %d", counter)),
	}

	fmt.Printf("Publishing message %d...\n", counter)
	if _, err := topic.Publish(ctx, msg).Get(ctx); err != nil {
		log.Fatalf("Could not publish message %d: %v", counter, err)
	}

	fmt.Println("Done!")
}
