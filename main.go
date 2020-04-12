package main

import (
	"github.com/Heiss/TwitchChat/handler"
	"github.com/Heiss/TwitchChat/subscriber"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	TwitchChat "github.com/Heiss/TwitchChat/proto/TwitchChat"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.TwitchChat"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	TwitchChat.RegisterTwitchChatHandler(service.Server(), new(handler.TwitchChatStruct))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.TwitchChat", service.Server(), new(subscriber.TwitchChatStruct))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
