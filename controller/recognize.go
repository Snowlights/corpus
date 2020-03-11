package controller

import (
	"context"
	corpus "github.com/Snowlights/pub/grpc"
)

func RecognizeImage(ctx context.Context,req *corpus.RecognizeImageReq) *corpus.RecognizeImageRes{
	res :=&corpus.RecognizeImageRes{}
	return res
}

func TransAudioToText(ctx context.Context,req* corpus.TransAudioToTextReq) *corpus.TransAudioToTextRes{
	res := &corpus.TransAudioToTextRes{}
	return res
}

func RecognizeAge(ctx context.Context, req *corpus.RecognizeAgeReq) *corpus.RecognizeAgeRes{
	res := &corpus.RecognizeAgeRes{}
	return res
}