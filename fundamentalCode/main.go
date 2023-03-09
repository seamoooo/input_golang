package main

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/pubsub"
)

const topicName = "projects/inputgcp2022/topics/goTest"
const projectId = "inputgcp2022"
const topicId = "goTest"

func main() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectId)

	if err != nil {
		fmt.Print("test ", err)
		os.Exit(1)
	}

	for i := 0; i < 10; i++ {
		t := client.Topic(topicId)
		result := t.Publish(ctx, &pubsub.Message{
			Data: []byte("hello"),
		})

		id, err := result.Get(ctx)
		if err != nil {
			fmt.Print(err)
		}

		fmt.Printf("Published a message; msg ID: %v\n", id)
	}
}
