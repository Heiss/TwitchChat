package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"github.com/Heiss/TwitchChat/handler"
	"github.com/Heiss/TwitchChat/subscriber"

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
	TwitchChat.RegisterTwitchChatHandler(service.Server(), new(handler.TwitchChat))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.TwitchChat", service.Server(), new(subscriber.TwitchChat))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
