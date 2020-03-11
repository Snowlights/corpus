package controller

import (
	"context"
	corpus "github.com/Snowlights/pub/grpc"
)

func GetKeyWord(ctx context.Context,req*corpus.GetKeyWordReq) *corpus.GetKeyWordRes{
	res := &corpus.GetKeyWordRes{}
	return res
}
