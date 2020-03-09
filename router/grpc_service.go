package router

import (
	"context"
	corpus "github.com/Snowlights/pub/grpc"
	"github.com/Snowlights/corpus/controller"
)

type CorpusServiceTmp struct {
}

func (m *CorpusServiceTmp) LoginUser(ctx context.Context, req *corpus.LoginUserReq) (*corpus.LoginUserRes,error){
	return controller.LoginUser(ctx,req),nil
}

func (m *CorpusServiceTmp) LogoutUserInfo(ctx context.Context, req *corpus.LogoutUserInfoReq) (*corpus.LogoutUserInfoRes,error){
	return controller.LoginOutUserInfo(ctx,req),nil
}

func (m *CorpusServiceTmp) SendMessage(ctx context.Context,req *corpus.SendMessageReq)(*corpus.SendMessageRes,error){
	return controller.SendMessage(ctx,req),nil
}

func (m *CorpusServiceTmp) UpdateUserInfo(ctx context.Context, req *corpus.UpdateUserInfoReq) (*corpus.UpdateUserInfoRes,error){
	return controller.UpdateUserInfo(ctx,req),nil
}

func (m *CorpusServiceTmp) DelUserInfo(ctx context.Context,req *corpus.DelUserInfoReq) (*corpus.DelUserInfoRes,error){
	return controller.DelUserInfo(ctx,req),nil
}

func (m *CorpusServiceTmp) ListUserInfo(ctx context.Context, req *corpus.ListUserInfoReq) (*corpus.ListUserInfoRes,error){
	return controller.ListUserInfo(ctx,req),nil
}

func (m *CorpusServiceTmp) AddTransAudio(ctx context.Context, req *corpus.AddTransAudioReq) (*corpus.AddTransAudioRes,error){

}

func (m *CorpusServiceTmp) DelTransAudio(ctx context.Context, req *corpus.DelTransAudioReq) (*corpus.DelTransAudioRes,error){

}

func (m *CorpusServiceTmp) ListTransAudio(ctx context.Context, req *corpus.ListTransAudioReq) (*corpus.ListTransAudioRes,error){

}

func (m *CorpusServiceTmp) GetKeyWord(ctx context.Context,req *corpus.GetKeyWordReq) (*corpus.GetKeyWordRes,error){

}

func (m *CorpusServiceTmp) Evaluation(ctx context.Context, req* corpus.EvaluationReq) (*corpus.EvaluationRes,error){
}

func (m *CorpusServiceTmp) RecognizeImage(ctx context.Context, req* corpus.RecognizeImageReq) (*corpus.RecognizeImageRes,error){

}

func (m *CorpusServiceTmp) TransAudioToText(ctx context.Context,req *corpus.TransAudioToTextReq) (*corpus.TransAudioToTextRes,error){

}

func (m *CorpusServiceTmp) RecognizeAge(ctx context.Context,req* corpus.RecognizeAgeReq) (*corpus.RecognizeAgeRes,error){

}