package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	TwitchChat "github.com/Heiss/TwitchChat/proto/TwitchChat"
)

type TwitchChat struct{}

func (e *TwitchChat) Handle(ctx context.Context, msg *TwitchChat.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *TwitchChat.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
