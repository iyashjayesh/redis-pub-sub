package main

import (
	"log"
	"redisgoexample/utils"
	"time"
)

func main() {

	// subcribe to the pipeline and listen for messages
	client := utils.CreateClient()
	log.Println("Subscribing to the test channel")
	utils.SubscribeToChannel(client, "test")

	time.Sleep(10 * time.Minute)
}
