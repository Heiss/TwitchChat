package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	TwitchChat "github.com/Heiss/TwitchChat/proto/TwitchChat"
)

type TwitchChat struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *TwitchChat) Call(ctx context.Context, req *TwitchChat.Request, rsp *TwitchChat.Response) error {
	log.Info("Received TwitchChat.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *TwitchChat) Stream(ctx context.Context, req *TwitchChat.StreamingRequest, stream TwitchChat.TwitchChat_StreamStream) error {
	log.Infof("Received TwitchChat.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&TwitchChat.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *TwitchChat) PingPong(ctx context.Context, stream TwitchChat.TwitchChat_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&TwitchChat.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
