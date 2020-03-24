package controller

import (
	corpus "github.com/Snowlights/pub/grpc"
	"testing"
)

func TestSendMessage(t *testing.T) {
	ctx := initenv()
	req := &corpus.SendMessageReq{
		Phone:                "",
	}
	SendMessage(ctx,req)

}