package utils

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var RedisClient *redis.Client

// create a redis client
func CreateClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

// create a redis pipeline and publish a message to a channel
func PublishMessage(client *redis.Client, channel string, message string) {
	pipe := client.Pipeline()
	pipe.Publish(ctx, channel, message)
	_, err := pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}
}

// subscribe to a channel from the pipeline and print the message
func SubscribeToChannel(client *redis.Client, channel string) {
	pubsub := client.Subscribe(ctx, channel)
	defer pubsub.Close()
	ch := pubsub.Channel()
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}
}
