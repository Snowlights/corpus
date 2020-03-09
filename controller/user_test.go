package controller

import (
	"context"
	"github.com/Snowlights/corpus/model"
	corpus "github.com/Snowlights/pub/grpc"
	"testing"
)

func initenv() context.Context{
	ctx := context.Background()
	model.Prepare(ctx)

	return ctx
}

func TestLoginUser(t *testing.T) {
	ctx:= initenv()
	req := &corpus.LoginUserReq{
		EMail:                "wei1109942647",
		UserPassword:         "",
		Phone:                "",
		Code:                 0,
		Cookie:               "",
	}


	LoginUser(ctx,req)
}