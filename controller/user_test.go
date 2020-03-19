package controller

import (
	"context"
	"fmt"
	"github.com/Snowlights/corpus/cache"
	"github.com/Snowlights/corpus/model"
	"github.com/Snowlights/corpus/model/daoimpl"
	corpus "github.com/Snowlights/pub/grpc"
	"log"
	"math/rand"
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
		EMail:                "858777157@qq.com",
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
		Audio:                "C:\\Users\\华硕\\Desktop\\pr\\evaluation\\eng\\en_chapter.wav",
		Text: 				  "Firefighters take part in an emergency rescue drill in a forest in Taian city, Shandong province, on Feb 24, 2019. This is the countrys largest joint air-ground drill with around 2,000 rescuers, seven helicopters and vehicles, and over 1,200 firefighting equipment taking part in the exercise.",
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


func TestTransAudioToText(t *testing.T) {

	file := "C:\\Users\\华硕\\Desktop\\pr\\evaluation\\eng\\en_sentence.wav"

	audioText(file)

}

func TestRecognizeAge(t *testing.T) {
	file := "C:\\Users\\华硕\\Desktop\\pr\\evaluation\\eng\\en_sentence.wav"
	recognizeAge(file)
}

func TestPictureResult_String(t *testing.T) {
	picture_trans("")
}

func TestDelUserInfo(t *testing.T) {
	number := rand.Intn(10)
	fmt.Print(number)

	file := "C:\\Users\\华硕\\Desktop\\file.png"
	handwriting_xunfei(file)
}