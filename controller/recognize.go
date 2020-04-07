package controller

import (
	"context"
	"fmt"
	"github.com/Snowlights/corpus/cache"
	"github.com/Snowlights/corpus/model/daoimpl"
	"github.com/Snowlights/corpus/model/domain"
	corpus "github.com/Snowlights/pub/grpc"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func RecognizeImage(ctx context.Context,req *corpus.RecognizeImageReq) *corpus.RecognizeImageRes{
	fun := "Controller.RecognizeImage --> "
	res :=&corpus.RecognizeImageRes{}
	pass := cache.CheckOnLine(req.Cookie)
	if !pass{
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  "未检测到登陆信息",
		}
		return res
	}

	pass = cache.CheckUserAuth(ctx,cache.ImageAuthCode,req.Cookie) || cache.CheckSuperAdmin(ctx,req.Cookie)
	if !pass{
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  "用户未享有该权限",
		}
		return res
	}

	rand.Seed(time.Now().Unix())
	number := rand.Intn(10)
	var pis []byte
	var md5 string
	tmp := req.File[0:4]
	if tmp == "http"{
		resp, err := http.Get(req.File)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		pix, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		outPath, err := SaveImage(req.File)
		if err != nil{
			logs.Error(err)
		}
		go MyDCT(outPath)
		md5,err = Md5(outPath)
		if err != nil{
			logs.Error(err)
		}
		pis = pix
	} else {
		dat, err := ioutil.ReadFile(req.File)
		if err != nil {
			logs.Error(err)
			return nil
		}
		md5,err = Md5(req.File)
		if err != nil{
			logs.Error(err)
		}
		pis = dat
	}


	var data,audit map[string]interface{}
	pictureTrans := picture_trans(pis)

	if number < 5{
		resp := handwriting_baidu(pis)
		data, audit = toFormRecognizeImageBaidu(req,resp,pictureTrans,md5)
	} else {
		resp := handwriting_xunfei(pis)
		data, audit = toFormRecognizeImageXf(req,resp,pictureTrans,md5)
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

	res.Data = &corpus.RecognizeImageData{
		Text:                 data["picture_text"].(string),
	}
	log.Printf("%v %v success ,lastInsertId %d",ctx,fun,lastInsertId)
	return res
}

func toFormRecognizeImageBaidu(req *corpus.RecognizeImageReq,resp *BaiduPicture,presp *PictureTransData,md5 string)(map[string]interface{},map[string]interface{}){
	now := time.Now().Unix()
	text := ""
	for _,item := range resp.WordsResult{
		text = text + item.Words
	}
	data := map[string]interface{}{
		"picture_src" : req.File,
		"picture_des" : presp.Output.Url,
		"picture_text" : text,
		"md5" : md5,
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

func toFormRecognizeImageXf(req *corpus.RecognizeImageReq,resp *XfPictureResp,presp *PictureTransData,md5 string)(map[string]interface{},map[string]interface{}){
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
		"md5":  md5,
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
	pass := cache.CheckOnLine(req.Cookie)
	if !pass{
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  "未检测到登陆信息",
		}
		return res
	}

	pass = cache.CheckUserAuth(ctx,cache.AudioAuthCode,req.Cookie) || cache.CheckSuperAdmin(ctx,req.Cookie)
	if !pass{
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  "用户未享有该权限",
		}
		return res
	}

	out,err := transformToPCM(req.Audio)
	if err != nil{
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  "错误的音频格式",
		}
		return res
	}
	fmt.Printf("%v\n",out)
	resString := audioText(out)
	data ,audit := toTransAudioToText(ctx,req,resString,out)
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

	res.Data = &corpus.TransAudioToTextData{
		Text:                 data["audio_text"].(string),
	}
	log.Printf("%v %v success ,lastInsertId %d",ctx,fun,lastInsertId)
	return res
}

func toTransAudioToText(ctx context.Context, req *corpus.TransAudioToTextReq,message string,out string)(map[string]interface{},map[string]interface{}){
	now := time.Now().Unix()
	data := map[string]interface{}{
		"audio_src" :req.Audio,
		"audio_trans_from" : fmt.Sprintf("xfyun.com %v",out),
		"audio_text" : message,
		"created_at" : now,
		"created_by" : req.Cookie,
		"updated_at" : now,
		"updated_by" : req.Cookie,
		"is_deleted" : false,
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
	pass := cache.CheckOnLine(req.Cookie)
	if !pass{
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  "未检测到登陆信息",
		}
		return res
	}

	pass = cache.CheckUserAuth(ctx,cache.AgeAuthCode,req.Cookie) || cache.CheckSuperAdmin(ctx,req.Cookie)
	if !pass{
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  "用户未享有该权限",
		}
		return res
	}

	tmp := req.Audio[len(req.Audio)-3:len(req.Audio)]
	var out string
	if tmp != "wav"{
		output,err := transformToWAV(req.Audio)
		if err != nil{
			res.Errinfo = &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  "错误的音频类型",
			}
			return res
		}
		out = output
	} else {
		out = req.Audio
	}

	resp := recognizeAge(out)
	data ,audit := toRecognizeAge(ctx,req,resp,out)
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

	res.Data = &corpus.RecognizeAgeData{
		AgeChild:             data["child_score"].(string),
		AgeMiddle:            data["middle_score"].(string),
		AgeOld:               data["child_score"].(string),
		Gender:               data["gender_type"].(string),
		GenderMale:           data["gender_female"].(string),
		GenderFemale:         data["gender_male"].(string),
	}
	log.Printf("%v %v success ,lastInsertId %d",ctx,fun,lastInsertId)
	return res
}
func toRecognizeAge(ctx context.Context, req *corpus.RecognizeAgeReq,resp *RespData,out string) (map[string]interface{},map[string]interface{}){
	now := time.Now().Unix()
	data := map[string]interface{}{
		"audio_src" : fmt.Sprintf("src %v des %v",req.Audio,out),
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

