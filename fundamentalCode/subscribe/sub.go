package main

import (
	"context"
	"fmt"
	"os"
	"sync"

	"cloud.google.com/go/pubsub"
)

const projectId = ""
const subName = ""

func main() {
	var mu sync.Mutex
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, projectId)

	if err != nil {
		fmt.Print("test ", err)
		os.Exit(1)
	}

	// サブスクリプションの参照作成
	sub := client.Subscription(subName)

	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()
		fmt.Printf("Got message: %q\n", string(msg.Data))
		mu.Lock()
		defer mu.Unlock()
	})
	if err != nil {
		fmt.Print(err)
	}
}
