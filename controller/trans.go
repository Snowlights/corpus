package controller

import (
	"context"
	"fmt"
	"github.com/Snowlights/corpus/cache"
	"github.com/Snowlights/corpus/model/daoimpl"
	"github.com/Snowlights/corpus/model/domain"
	corpus "github.com/Snowlights/pub/grpc"
	"github.com/astaxie/beego/logs"
	"github.com/facebookgo/runcmd"
	"log"
	"os/exec"
	"time"
)

func AddTransAudio(ctx context.Context,req *corpus.AddTransAudioReq) *corpus.AddTransAudioRes{
	fun := "Controller.AddTransAudio -->"
	res := &corpus.AddTransAudioRes{}
	pass := cache.CheckOnLine(req.Cookie)
	if !pass{
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  "未检测到登陆信息",
		}
		return res
	}

	pass = cache.CheckUserAuth(ctx,cache.TransAuthCode,req.Cookie) || cache.CheckSuperAdmin(ctx,req.Cookie)
	if !pass{
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  "用户未享有该权限",
		}
		return res
	}
	var output string
	var err error
	switch req.AudioType{
	case corpus.AudioType_Mp3:
		output,err = transformToMP3(req.OriginAudio)
	case corpus.AudioType_ACC:
		output,err = transformToAAC(req.OriginAudio)
	case corpus.AudioType_WAV:
		output,err = transformToWAV(req.OriginAudio)
	}
	if err != nil{
		logs.Error("%v error %v",fun,err)
	}

	data,audit := toAddTransAudio(ctx,req,output)

	lastInsertId,err := daoimpl.AudioDao.AddAudio(ctx,data)
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

func toAddTransAudio(ctx context.Context,req *corpus.AddTransAudioReq,output string)(map[string]interface{},map[string]interface{}){
	now := time.Now().Unix()
	data := map[string]interface{}{
		"audio_src" : req.OriginAudio,
		"audio_des" : output,
		"audio_type" : int64(req.AudioType),
		"created_at" : now,
		"created_by" : req.Cookie,
		"updated_at" : now,
		"updated_by" : req.Cookie,
		"is_deleted" : false,
	}
	audit := map[string]interface{}{
		"table_name" : domain.EmptyAudio.TableName(),
		"history" : fmt.Sprintf("%s add trans audio %v time %d ",req.Cookie,req.OriginAudio,now),
		"activity" : fmt.Sprintf("%s add trans audio",req.Cookie),
		"content" : fmt.Sprintf("%v",data),
		"created_at" : now,
		"created_by" : req.Cookie,
		"is_deleted" : false,
	}
	return data,audit
}


func DelTransAudio(ctx context.Context,req *corpus.DelTransAudioReq) *corpus.DelTransAudioRes{
	fun := "Controller.DelTransAudio -->"
	res := &corpus.DelTransAudioRes{}

	pass := cache.CheckOnLine(req.Cookie)
	if !pass{
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  "未检测到登陆信息",
		}
		return res
	}
	limit := map[string]interface{}{
		"limit" : 1,
		"offset" : 0,
	}
	cond := map[string]interface{}{
		"id" : req.AudioId,
		"is_deleted" : false,
	}
	audioInfo,err := daoimpl.AudioDao.ListAudio(ctx,limit,cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  err.Error(),
		}
		return res
	}
	if len(audioInfo) == 0{
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  "未找到该转码信息",
		}
		return res
	}

	if req.Cookie != audioInfo[0].CreatedBy{
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  "非本人不可删除",
		}
		return res
	}

	data, conds ,audit := toDelTransAudio(ctx,req)

	rowsAffected, err := daoimpl.AudioDao.DelAudio(ctx,data,conds)
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

	log.Printf("%v %v success ,rowsAffected %d",ctx,fun,rowsAffected)

	return res
}

func toDelTransAudio(ctx context.Context,req *corpus.DelTransAudioReq) (map[string]interface{},map[string]interface{},map[string]interface{}){
	now := time.Now().Unix()
	data := map[string]interface{}{
		"is_deleted": true,
		"updated_at" : now,
		"updated_by" : req.Cookie,
	}
	conds := map[string]interface{}{
		"id" : req.AudioId,
	}

	audit := map[string]interface{}{
		"table_name" : domain.EmptyAudio.TableName(),
		"history" : fmt.Sprintf("%s del trans audio id %v time %d ",req.Cookie,req.AudioId,now),
		"activity" : fmt.Sprintf("%s del trans audio",req.Cookie),
		"content" : fmt.Sprintf("%v",data),
		"created_at" : now,
		"created_by" : req.Cookie,
		"is_deleted" : false,
	}
	return data,conds,audit
}


func ListTransAudio(ctx context.Context,req *corpus.ListTransAudioReq) *corpus.ListTransAudioRes{
	fun := "Controller.ListTransAudio -->"
	res := &corpus.ListTransAudioRes{}

	limit, conds := toListTransAudio(ctx,req)
	dataList, err := daoimpl.AudioDao.ListAudio(ctx,limit,conds)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  err.Error(),
		}
		return res
	}
	total, err := daoimpl.AudioDao.CountAudio(ctx,conds)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  err.Error(),
		}
		return res
	}

	data := formCorpusAudioList(dataList)
	res.Data = &corpus.ListTransAudioData{
		Items:                data,
		Total:                total,
		Offset:               req.Offset + int64(len(dataList)),
		More:                 (req.Offset + int64(len(dataList))) < total,
	}

	return res
}

func formCorpusAudioList(input []*domain.AudioInfo) []*corpus.Audio{
	var output []*corpus.Audio

	for _, item := range input{
		output = append(output,&corpus.Audio{
			AudioId:              item.Id,
			OriginAudio:          item.AudioSrc,
			TransAudio:           item.AudioDes,
			AudioType:            item.AudioType,
		})
	}
	return output
}


func toListTransAudio(ctx context.Context,req *corpus.ListTransAudioReq) (map[string]interface{},map[string]interface{}){
	limit := map[string]interface{}{
		"limit" : req.Limit,
		"offset" : req.Offset,
	}
	conds := map[string]interface{}{
		"is_deleted" : false,
	}
	if req.AudioType != 0 {
		conds["audio_type"] = req.AudioType
	}

	return limit,conds
}


func transformToAAC(fpath string) (string, error) {
	fun := "logic.transformToAAC -->"

	md5Str, err := MD5Audio(fpath)
	if err != nil{
		logs.Error(err)
		return "错误的MD5值",err
	}
	// transform to wav
	output := fmt.Sprintf("C:\\Users\\华硕\\Desktop\\pr\\age\\%v.aac",md5Str)
	cmdSlice := []string{}
	cmdSlice = append(cmdSlice, "-i", fpath, output)
	cmd := exec.Command("C:\\Users\\华硕\\Desktop\\ffmpeg\\bin\\ffmpeg.exe", cmdSlice...)
	//cmd := exec.Command("ffmpeg", cmdSlice...)  //linux依赖 使用系统命令ffmpeg
	err = cmd.Start()
	if err != nil{
		logs.Error(fun,err)
		return "",err
	}
	//streams, err := runcmd.Run(cmd)
	//if err != nil {
	//	log.Printf("%s trans to AAC failed, error: %s, stdout: %s, stderr: %s", fun, err, streams.Stderr().String(), streams.Stdout().String())
	//	return "", err
	//}
	return output, nil
}

func transformToWAV(fpath string) (string, error) {
	fun := "logic.transformToWAV -->"

	md5Str, err := MD5Audio(fpath)
	if err != nil{
		logs.Error(err)
		return "错误的MD5值",err
	}
	// transform to wav
	output := fmt.Sprintf("C:\\Users\\华硕\\Desktop\\pr\\age\\%v.wav",md5Str)

	cmdSlice := []string{}
	cmdSlice = append(cmdSlice, "-y", "-i", fpath, "-b:a", "16k", "-ar", "16000", "-ac", "1", output)
	//cmdSlice = append(cmdSlice, "-i", fpath, "-f", "wav", output)
	cmd := exec.Command("C:\\Users\\华硕\\Desktop\\ffmpeg\\bin\\ffmpeg.exe", cmdSlice...)
	//cmd := exec.Command("ffmpeg", cmdSlice...)  //linux依赖 使用系统命令ffmpeg

	streams, err := runcmd.Run(cmd)
	if err != nil {
		log.Printf("%s trans to wav failed, error: %s, stdout: %s, stderr: %s", fun, err, streams.Stderr().String(), streams.Stdout().String())
		return "", err
	}
	return output, nil
}

func transformToMP3(fpath string) (string, error) {
	fun := "logic.transformToMp3 -->"

	// transform to wav
	md5Str, err := MD5Audio(fpath)
	if err != nil{
		logs.Error(err)
		return "错误的MD5值",err
	}
	// transform to wav
	output := fmt.Sprintf("C:\\Users\\华硕\\Desktop\\pr\\age\\%v.mp3",md5Str)

	cmdSlice := []string{}
	cmdSlice = append(cmdSlice,"-i", fpath, "-f", "mp3", "-acodec", "libmp3lame", "-y", output)
	//cmdSlice = append(cmdSlice, "-i", fpath, "-f", "wav", output)
	cmd := exec.Command("C:\\Users\\华硕\\Desktop\\ffmpeg\\bin\\ffmpeg.exe",cmdSlice...)
	//cmd := exec.Command("ffmpeg", cmdSlice...)  //linux依赖 使用系统命令ffmpeg

	streams, err := runcmd.Run(cmd)
	if err != nil {
		log.Printf("%s trans to wav failed, error: %s, stdout: %s, stderr: %s", fun, err, streams.Stderr().String(), streams.Stdout().String())
		return "", err
	}
	return output, nil
}

func transformToPCM(fpath string) (string, error) {
	fun := "logic.transformToPCM -->"

	// transform to wav
	md5Str, err := MD5Audio(fpath)
	if err != nil{
		logs.Error(err)
		return "错误的MD5值",err
	}
	// transform to wav
	output := fmt.Sprintf("C:\\Users\\华硕\\Desktop\\pr\\age\\%v.pcm",md5Str)

	cmdSlice := []string{}
	cmdSlice = append(cmdSlice,"-y","-i", fpath, "-acodec", "pcm_s16le", "-f", "s16le", "-ac","1","-ar","16000", output)
	//cmdSlice = append(cmdSlice, "-i", fpath, "-f", "wav", output)
	cmd := exec.Command("C:\\Users\\华硕\\Desktop\\ffmpeg\\bin\\ffmpeg.exe",cmdSlice...)
	//cmd := exec.Command("ffmpeg", cmdSlice...)  //linux依赖 使用系统命令ffmpeg

	streams, err := runcmd.Run(cmd)
	if err != nil {
		log.Printf("%s trans to wav failed, error: %s, stdout: %s, stderr: %s", fun, err, streams.Stderr().String(), streams.Stdout().String())
		return "", err
	}
	return output, nil
}