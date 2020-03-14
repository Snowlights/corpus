package controller

import (
	"context"
	"github.com/Snowlights/corpus/cache"
	"github.com/Snowlights/corpus/model"
	"github.com/Snowlights/corpus/model/daoimpl"
	corpus "github.com/Snowlights/pub/grpc"
	"log"
	"testing"
	"time"
)

func initenv() context.Context{
	ctx := context.Background()
	model.Prepare(ctx)
	daoimpl.PrePare(ctx)
	cache.Prepare(ctx)
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

func TestAddAdminUser(t *testing.T) {
	ctx := initenv()
	req := &corpus.AddAdminUserReq{
		UserId:               4,
		Cookie:               "",
	}
	r := AddAdminUser(ctx,req)

	log.Printf("%v",r)
}

func TestAddAuth(t *testing.T) {
	ctx := initenv()
	req := &corpus.AddAuthReq{
		AuthCode:             "Evaluation_Service",
		AuthDescription:      "打分服务权限码",
		ServiceName:          "http://service/evaluation/11",
		Cookie:               "JnL3gxsI402j4hs4",
	}
	r := AddAuth(ctx,req)
	log.Printf("%v",r)
}

func TestEvaluation(t *testing.T) {
	ctx := initenv()
	req := &corpus.EvaluationReq{
		Audio:                "C:\\Users\\华硕\\Desktop\\pr\\evaluation\\cn\\cn_sentence.wav",
		Text: 				  "不管我的梦想能否成为事实.",
		Cookie:               "JnL3gxsI402j4hs4",
	}
	r := Evaluation(ctx,req)

	log.Printf("%v",r)

	time.Sleep(time.Second*5)
}

func TestGetKeyWord(t *testing.T) {
	ctx := initenv()
	req := &corpus.GetKeyWordReq{
		Text:                 "C:\\Users\\华硕\\Desktop\\pr\\1.txt",
		Cookie:               "",
	}
	r := GetKeyWord(ctx, req)

	log.Printf("%v",r)

	time.Sleep(time.Second*5)
}