package main

import (
	"log"
	"redisgoexample/utils"
	"time"
)

func main() {

	// subcribe to the pipeline and listen for messages
	client := utils.CreateClient()
	log.Println("Subscribing to the monitor channel")
	utils.SubscribeToChannel(client, "monitor")

	time.Sleep(10 * time.Minute)
}
