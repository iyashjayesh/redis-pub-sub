package main

import (
	"log"
	"redisgoexample/utils"
	"time"
)

func main() {

	//Publishing events
	client := utils.CreateClient()
	log.Println("Publishing to the test channel")
	for i := 0; i < 100; i++ {
		utils.PublishMessage(client, "test", "This is Message from test Subject")
		log.Println("Published message from test Subject")
		time.Sleep(3 * time.Second)
	}
}
