package main

import (
	"log"
	"redisgoexample/utils"
	"time"
)

func main() {

	//Publishing events
	client := utils.CreateClient()
	for i := 0; i < 100; i++ {
		utils.PublishMessage(client, "monitor", "This is Message from Monitor")
		log.Println("Published message from Monitor Subject")
		time.Sleep(3 * time.Second)
	}
}
