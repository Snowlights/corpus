package controller

import (
	"bufio"
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
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

	r, err := geyWord(req.Text)
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
	keyWordData,keyData,audit := toFormKeyWord(&info,req)

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
	log.Printf("%v %v success ,auditLastInsertId %d",ctx,fun,auditLastInsertId)
	return res
}

func toFormKeyWord(keyword *KeyWordResponse,req *corpus.GetKeyWordReq) (map[string]interface{},[]map[string]interface{},map[string]interface{}){
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
		"history" : fmt.Sprintf("%s add KeyWord time %d ",req.Cookie,now),
		"activity" : fmt.Sprintf("%s  add KeyWord info",req.Cookie),
		"content" : fmt.Sprintf("%v",keyWordData),
		"created_at" : now,
		"created_by" : req.Cookie,
		"is_deleted" : false,
	}

	return keyWordData,keyData,auditData
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
	res_body, _ := ioutil.ReadAll(res.Body)
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
