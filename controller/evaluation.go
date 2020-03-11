package controller

import (
	"context"
	corpus "github.com/Snowlights/pub/grpc"
)

func Evaluation(ctx context.Context,req *corpus.EvaluationReq) *corpus.EvaluationRes{
	res := &corpus.EvaluationRes{}
	return res
}