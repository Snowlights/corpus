package controller

import (
	"context"
	"fmt"
	"github.com/Snowlights/corpus/model/daoimpl"
	"github.com/Snowlights/corpus/model/domain"
	corpus "github.com/Snowlights/pub/grpc"
	"log"
	"math/rand"
	"time"
)

func RecognizeImage(ctx context.Context,req *corpus.RecognizeImageReq) *corpus.RecognizeImageRes{
	fun := "Controller.RecognizeImage --> "
	res :=&corpus.RecognizeImageRes{}
	rand.Seed(time.Now().Unix())
	number := rand.Intn(10)

	var data,audit map[string]interface{}
	pictureTrans := picture_trans(req.File)

	if number < 5{
		resp := handwriting_baidu(req.File)
		data, audit = toFormRecognizeImageBaidu(req,resp,pictureTrans)
	} else {
		resp := handwriting_xunfei(req.File)
		data, audit = toFormRecognizeImageXf(req,resp,pictureTrans)
	}
	lastInsertId,err := daoimpl.RecognizeDao.AddRecognizeImage(ctx,data)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  err.Error(),
		}
		return res
	}
	auditLastInsertId, err := addAudit(ctx,audit)
	if err!= nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
	}
	log.Printf("%v %v success ,auditLastInsertId %d",ctx,fun,auditLastInsertId)

	log.Printf("%v %v success ,lastInsertId %d",ctx,fun,lastInsertId)
	return res
}

func toFormRecognizeImageBaidu(req *corpus.RecognizeImageReq,resp *BaiduPicture,presp *PictureTransData)(map[string]interface{},map[string]interface{}){
	now := time.Now().Unix()
	data := map[string]interface{}{
		"picture_src" : req.File,
		"picture_des" : presp.Output.Url,
		"picture_text" : resp.WordsResult,
		"created_at" : now,
		"created_by" : req.Cookie,
		"is_deleted" : false,
	}
	audit := map[string]interface{}{
		"table_name" : domain.EmptyRecognize.TableName(),
		"history" : fmt.Sprintf("%s recognize picture to text picture src %v time %d ",req.Cookie,req.File,now),
		"activity" : fmt.Sprintf("%s recognize picture",req.Cookie),
		"content" : fmt.Sprintf("%v",data),
		"created_at" : now,
		"created_by" : req.Cookie,
		"is_deleted" : false,
	}
	return data,audit
}

func toFormRecognizeImageXf(req *corpus.RecognizeImageReq,resp *XfPictureResp,presp *PictureTransData)(map[string]interface{},map[string]interface{}){
	now := time.Now().Unix()

	str := ""
	for _, item := range resp.Data.Block{

		for _ ,ite := range item.Line{
			for _, wo := range ite.Word{
				str = str + wo.Content
			}
		}
	}
	data := map[string]interface{}{
		"picture_src" : req.File,
		"picture_des" : presp.Output.Url,
		"picture_text" : str,
		"created_at" : now,
		"created_by" : req.Cookie,
		"is_deleted" : false,
	}
	audit := map[string]interface{}{
		"table_name" : domain.EmptyRecognize.TableName(),
		"history" : fmt.Sprintf("%s recognize picture to text picture src %v time %d ",req.Cookie,req.File,now),
		"activity" : fmt.Sprintf("%s recognize picture",req.Cookie),
		"content" : fmt.Sprintf("%v",data),
		"created_at" : now,
		"created_by" : req.Cookie,
		"is_deleted" : false,
	}
	return data,audit
}

func TransAudioToText(ctx context.Context,req* corpus.TransAudioToTextReq) *corpus.TransAudioToTextRes{
	fun := "Controller.TransAudioToText -->"
	res := &corpus.TransAudioToTextRes{}

	resp := audioText(req.Audio)
	if resp.Data.Status != 2{
		//cf()
		fmt.Printf("audioToText failed error code %v",resp.Code)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  "识别失败",
		}
		return res
	}
	message := resp.Data.Result.String()

	data ,audit := toTransAudioToText(ctx,req,message)
	lastInsertId,err := daoimpl.AudioDao.AddAudioToText(ctx,data)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  err.Error(),
		}
		return res
	}
	auditLastInsertId, err := addAudit(ctx,audit)
	if err!= nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
	}
	log.Printf("%v %v success ,auditLastInsertId %d",ctx,fun,auditLastInsertId)

	log.Printf("%v %v success ,lastInsertId %d",ctx,fun,lastInsertId)
	return res
}

func toTransAudioToText(ctx context.Context, req *corpus.TransAudioToTextReq,message string)(map[string]interface{},map[string]interface{}){
	now := time.Now().Unix()
	data := map[string]interface{}{
		"audio_src" :req.Audio,
		"audio_trans_from" : "xfyun.com",
		"audio_text" : message,
		"created_at" : now,
		"created_by" : req.Cookie,
		"updated_at" : now,
		"updated_by" : req.Cookie,
		"is_deleted_" : false,
	}
	audit := map[string]interface{}{
		"table_name" : domain.EmptyAudioText.TableName(),
		"history" : fmt.Sprintf("%s audio to text audio %v time %d ",req.Cookie,req.Audio,now),
		"activity" : fmt.Sprintf("%s audio to text",req.Cookie),
		"content" : fmt.Sprintf("%v",data),
		"created_at" : now,
		"created_by" : req.Cookie,
		"is_deleted" : false,
	}
	return data,audit
}



func RecognizeAge(ctx context.Context, req *corpus.RecognizeAgeReq) *corpus.RecognizeAgeRes{
	fun := "Controller.RecognizeAge"
	res := &corpus.RecognizeAgeRes{}

	resp := recognizeAge(req.Audio)
	data ,audit := toRecognizeAge(ctx,req,resp)
	lastInsertId,err := daoimpl.RecognizeDao.AddRecognizeAge(ctx,data)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  err.Error(),
		}
		return res
	}
	auditLastInsertId, err := addAudit(ctx,audit)
	if err!= nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
	}
	log.Printf("%v %v success ,auditLastInsertId %d",ctx,fun,auditLastInsertId)

	log.Printf("%v %v success ,lastInsertId %d",ctx,fun,lastInsertId)

	return res
}
func toRecognizeAge(ctx context.Context, req *corpus.RecognizeAgeReq,resp *RespData) (map[string]interface{},map[string]interface{}){
	now := time.Now().Unix()
	data := map[string]interface{}{
		"audio_src" : req.Audio,
		"recognize_age_type" : resp.Data.Result.Age.AgeType,
		"child_score" : resp.Data.Result.Age.Child,
		"middle_score" : resp.Data.Result.Age.Middle,
		"old_score" : resp.Data.Result.Age.Old,
		"gender_type" : resp.Data.Result.Gender.Gender_type,
		"gender_female" : resp.Data.Result.Gender.Female,
		"gender_male" : resp.Data.Result.Gender.Male,
		"created_at" : now,
		"created_by" : req.Cookie,
		"is_deleted" : false,
	}
	audit := map[string]interface{}{
		"table_name" : domain.EmptyRecognize.TableName(),
		"history" : fmt.Sprintf("%s recognize age audio id %v time %d ",req.Cookie,req.Audio,now),
		"activity" : fmt.Sprintf("%s recognize age",req.Cookie),
		"content" : fmt.Sprintf("%v",data),
		"created_at" : now,
		"created_by" : req.Cookie,
		"is_deleted" : false,
	}
	return data,audit
}

