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

func (m *CorpusServiceTmp) AddAdminUser(ctx context.Context,req *corpus.AddAdminUserReq) (*corpus.AddAdminUserRes,error){
	return controller.AddAdminUser(ctx,req),nil
}

func (m *CorpusServiceTmp) DelAdminUser(ctx context.Context,req *corpus.DelAdminUserReq) (*corpus.DelAdminUserRes,error){
	return controller.DelAdminUser(ctx,req),nil
}

func (m *CorpusServiceTmp) ListAdminUser(ctx context.Context,req *corpus.ListAdminUserReq) (*corpus.ListAdminUserRes,error){
	return controller.ListAdminUser(ctx,req),nil
}

func (m *CorpusServiceTmp) AddAuth(ctx context.Context, req *corpus.AddAuthReq) (*corpus.AddAuthRes,error){
	return controller.AddAuth(ctx,req),nil
}

func (m *CorpusServiceTmp) UpdateAuth(ctx context.Context,req *corpus.UpdateAuthReq) (*corpus.UpdateAuthRes,error){
	return controller.UpdateAuth(ctx,req),nil
}

func (m *CorpusServiceTmp) DelAuth(ctx context.Context,req *corpus.DelAuthReq) (*corpus.DelAuthRes,error){
	return controller.DelAuth(ctx,req),nil
}

func (m *CorpusServiceTmp) ListAuth(ctx context.Context,req *corpus.ListAuthReq) (*corpus.ListAuthRes,error){
	return controller.ListAuth(ctx,req),nil
}

func (m *CorpusServiceTmp) AddUserAuth(ctx context.Context, req* corpus.AddUserAuthReq) (*corpus.AddUserAuthRes,error){
	return controller.AddUserAuth(ctx,req),nil
}

func (m *CorpusServiceTmp) DelUserAuth(ctx context.Context,req* corpus.DelUserAuthReq) (*corpus.DelUserAuthRes,error){
	return controller.DelUserAuth(ctx,req),nil
}

func (m *CorpusServiceTmp) ListUserAuth(ctx context.Context,req *corpus.ListUserAuthReq) (*corpus.ListUserAuthRes,error){
	return controller.ListUserAuth(ctx,req),nil
}

func (m *CorpusServiceTmp) AddTransAudio(ctx context.Context, req *corpus.AddTransAudioReq) (*corpus.AddTransAudioRes,error){
	return controller.AddTransAudio(ctx,req),nil
}

func (m *CorpusServiceTmp) DelTransAudio(ctx context.Context, req *corpus.DelTransAudioReq) (*corpus.DelTransAudioRes,error){
	return controller.DelTransAudio(ctx,req),nil
}

func (m *CorpusServiceTmp) ListTransAudio(ctx context.Context, req *corpus.ListTransAudioReq) (*corpus.ListTransAudioRes,error){
	return controller.ListTransAudio(ctx,req),nil
}

func (m *CorpusServiceTmp) GetKeyWord(ctx context.Context,req *corpus.GetKeyWordReq) (*corpus.GetKeyWordRes,error){
	return controller.GetKeyWord(ctx,req),nil
}

func (m *CorpusServiceTmp) Evaluation(ctx context.Context, req* corpus.EvaluationReq) (*corpus.EvaluationRes,error){
	return controller.Evaluation(ctx,req),nil
}

func (m *CorpusServiceTmp) RecognizeImage(ctx context.Context, req* corpus.RecognizeImageReq) (*corpus.RecognizeImageRes,error){
	return controller.RecognizeImage(ctx,req),nil
}

func (m *CorpusServiceTmp) TransAudioToText(ctx context.Context,req *corpus.TransAudioToTextReq) (*corpus.TransAudioToTextRes,error){
	return controller.TransAudioToText(ctx,req),nil
}

func (m *CorpusServiceTmp) RecognizeAge(ctx context.Context,req* corpus.RecognizeAgeReq) (*corpus.RecognizeAgeRes,error){
	return controller.RecognizeAge(ctx,req),nil
}