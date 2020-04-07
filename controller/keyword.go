package controller

import (
	"bufio"
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Snowlights/corpus/cache"
	"github.com/Snowlights/corpus/model/daoimpl"
	"github.com/Snowlights/corpus/model/domain"
	corpus "github.com/Snowlights/pub/grpc"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetKeyWord(ctx context.Context,req*corpus.GetKeyWordReq) *corpus.GetKeyWordRes{
	fun := "Controller.GetKeyWord -->"
	res := &corpus.GetKeyWordRes{}
	pass := cache.CheckOnLine(req.Cookie)
	if !pass{
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  "未检测到登陆信息",
		}
		return res
	}

	pass = cache.CheckUserAuth(ctx,cache.KeyWordAuthCode,req.Cookie) || cache.CheckSuperAdmin(ctx,req.Cookie)
	if !pass{
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  "用户未享有该权限",
		}
		return res
	}

	outfile,err := writeText(req.Text)
	if err != nil{
		log.Printf("保存用户的文本失败，请重新尝试")
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  "服务器错误，请重新尝试",
		}
		return res
	}

	r, err := geyWord(outfile)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  err.Error(),
		}
		return res
	}
	var info KeyWordResponse
	err = json.Unmarshal(r,&info)
	keyWordData,keyData,audit := toFormKeyWord(&info,req,outfile)

	err = daoimpl.KeyWordTxDao.AddKeyWordTx(ctx,keyWordData,keyData)
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

	keywordData := toFormKeywordData(info)
	res.Data = &corpus.GetKeyWordData{
		Items:                keywordData,
		Total:                int64(len(keywordData)),
	}
	log.Printf("%v %v success ,auditLastInsertId %d",ctx,fun,auditLastInsertId)
	return res
}

func toFormKeywordData(info KeyWordResponse) []*corpus.KeyWord{
	var data []*corpus.KeyWord
	for _, item := range info.Data.Ke{
		score,_ := strconv.ParseFloat(item.Score,32)
		data = append(data,&corpus.KeyWord{
			Text:                 item.Word,
			Score:                float32(score),
		})
	}
	return data
}

func toFormKeyWord(keyword *KeyWordResponse,req *corpus.GetKeyWordReq,filename string) (map[string]interface{},[]map[string]interface{},map[string]interface{}){
	now := time.Now().Unix()
	keyWordData := map[string]interface{}{
		"origin_text" : req.Text,
		"created_at" : now,
		"created_by" : req.Cookie,
		"is_deleted" : false,
	}
	keyData := []map[string]interface{}{}

	for _, item := range keyword.Data.Ke{
		ite := map[string]interface{}{
			"word" : item.Word,
			"score" : item.Score,
			"created_at" : now,
			"created_by" : req.Cookie,
			"is_deleted" : false,
		}
		keyData = append(keyData,ite)
	}

	auditData := map[string]interface{}{
		"table_name" : domain.EmptyKeyWord.TableName(),
		"history" : fmt.Sprintf("%s add KeyWord time %d file name %v",req.Cookie,now,filename),
		"activity" : fmt.Sprintf("%s  add KeyWord info",req.Cookie),
		"content" : fmt.Sprintf("%v",keyWordData),
		"created_at" : now,
		"created_by" : req.Cookie,
		"is_deleted" : false,
	}

	return keyWordData,keyData,auditData
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeText(text string) (string,error){
	md5String := GetMd5String(text)
	var err1 error
	var f *os.File
	filename := fmt.Sprintf("C:\\Users\\华硕\\Desktop\\pr\\text\\%v.txt",md5String)
	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在,创建文件")
	}
	check(err1)
	n, err1 := io.WriteString(f, text) //写入文件(字符串)
	check(err1)
	fmt.Printf("%v 写入 %d 个字节n", filename, n)

	return filename,err1
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func geyWord(fileName string ) ([]byte,error){
	appid := "5e1c39a3"
	apikey := "86df52d25033f6404b5306c96abc48c6"
	post_url := "http://ltpapi.xfyun.cn/v1/ke"
	curtime := strconv.FormatInt(time.Now().Unix(), 10)

	param := make(map[string]string)
	param["type"] = "dependent"
	tmp, _ := json.Marshal(param)
	base64_param := base64.StdEncoding.EncodeToString(tmp)

	w := md5.New()
	io.WriteString(w, apikey+curtime+base64_param)
	checksum := fmt.Sprintf("%x", w.Sum(nil))

	var filestr string
	//file, err := os.Open("C:\\Users\\华硕\\Desktop\\finalproject\\test.txt")
	file, err := os.Open(fileName)
	if err != nil{
		fmt.Println("open file err ",err.Error())
	}
	defer file.Close()

	r := bufio.NewReader(file)

	for {
		str, err := r.ReadBytes('\n')
		filestr = string(str)

		if err == io.EOF{
			break
		}
		if err != nil{
			fmt.Println("read error ",err.Error())
			break
		}
	}

	data := url.Values{}
	data.Add("text", filestr)
	body := data.Encode()

	client := &http.Client{}
	req, _ := http.NewRequest("POST", post_url, strings.NewReader(body))
	req.Header.Set("X-Appid", appid)
	req.Header.Set("X-CurTime", curtime)
	req.Header.Set("X-Param", base64_param)
	req.Header.Set("X-CheckSum", checksum)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, _ := client.Do(req)
	defer res.Body.Close()
	res_body, err := ioutil.ReadAll(res.Body)
	fmt.Print(string(res_body))

	if err != nil{
		return nil,err
	}

	return (res_body),nil
}


type Ke struct {
	Score string `json:"score"`
	Word string `json:"word"`
}

type KeyWordData struct {
	Ke []Ke `json:"ke"`
} 

type KeyWordResponse struct {
	Code int64 `json:"code"`
	Data KeyWordData `json:"data"`
	Desc string `json:"desc"`
	Sid string `json:"sid"`
}
