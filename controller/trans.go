package controller

import (
	"context"
	corpus "github.com/Snowlights/pub/grpc"
)

func AddTransAudio(ctx context.Context,req *corpus.AddTransAudioReq) *corpus.AddTransAudioRes{
	res := &corpus.AddTransAudioRes{}
	return res
}

func DelTransAudio(ctx context.Context,req *corpus.DelTransAudioReq) *corpus.DelTransAudioRes{
	res := &corpus.DelTransAudioRes{}
	return res
}

func ListTransAudio(ctx context.Context,req *corpus.ListTransAudioReq) *corpus.ListTransAudioRes{
	res := &corpus.ListTransAudioRes{}
	return res
}