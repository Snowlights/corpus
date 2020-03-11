package controller

import (
	"context"
	"github.com/Snowlights/corpus/model"
	"github.com/Snowlights/corpus/model/daoimpl"
	corpus "github.com/Snowlights/pub/grpc"
	"log"
	"testing"
)

func initenv() context.Context{
	ctx := context.Background()
	model.Prepare(ctx)
	daoimpl.PrePare(ctx)
	return ctx
}

func TestLoginUser(t *testing.T) {
	ctx:= initenv()
	req := &corpus.LoginUserReq{
		EMail:                "xueyeup@163.com",
		UserPassword:         "woaini12",
		Phone:                "",
		Code:                 0,
	}

	r := LoginUser(ctx,req)

	log.Printf("%v",r)

}