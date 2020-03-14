package controller

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Snowlights/corpus/model/daoimpl"
	corpus "github.com/Snowlights/pub/grpc"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func Evaluation(ctx context.Context,req *corpus.EvaluationReq) *corpus.EvaluationRes{
	fun := "Controller.Evaluation -->"
	res := &corpus.EvaluationRes{}
	//todo check cookie

	data := evaluation(req.Audio,req.Text)

	dataList,_ := analizeEvaluationData(data,req)

	lastInsertId,err := daoimpl.EvaluationDao.AddEvaluation(ctx,dataList)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  err.Error(),
		}
		return res
	}
	//todo items word
	res.Data = &corpus.EvaluationData{
		Items:                nil,
		TotalNumber:          0,
		TotalScore:           0,
	}

	log.Printf("%v",data)
	log.Printf("%v %v success ,lastInsertId %d",ctx,fun,lastInsertId)
	return res
}

func analizeEvaluationData(data []byte,req *corpus.EvaluationReq) (map[string]interface{},*EvaluationResponse){
	dataList := map[string]interface{}{}
	dataList["audio_src"] = req.Audio
	dataList["audio_text"] = req.Text
	now := time.Now().Unix()

	var res EvaluationResponse
	err  := json.Unmarshal(data,&res)
	if err != nil{
		log.Printf("analizeEvaluationData failed ")
		return dataList,&res
	}
	//if res.ReadSentence.RecPaper.ReadWord.TotalScore != 0{
	//	dataList["total_score"] = res.ReadSentence.RecPaper.ReadWord.TotalScore
	//} else {
	//	dataList["total_score"] = 0
	//}
	origindata := string(data)
	dataList["total_score"] = int64(0)
	dataList["original_data"] = origindata
	dataList["created_at"] = now
	dataList["created_by"] = req.Cookie
	dataList["is_deleted"] = false

	return dataList,&res
}


func evaluation(filename string,text string) []byte{

	appid := "5e1c39a3"
	apikey := "6ffbf970ad40212e91e01d979c36a09c"
	curtime := strconv.FormatInt(time.Now().Unix(),10)

	param := make(map[string]string)
	param["aue"] = "raw"
	param["language"] = "zh_cn"
	param["category"] = "read_sentence"

	tmp, _ := json.Marshal(param)
	base64_param := base64.StdEncoding.EncodeToString(tmp)

	w := md5.New()
	io.WriteString(w, apikey+curtime+base64_param)
	checksum := fmt.Sprintf("%x", w.Sum(nil))

	//f, _ := ioutil.ReadFile("C:\\evaluation\\test\\chinese\\cn_chapter.wav")
	f, _ := ioutil.ReadFile(filename)
	audio := base64.StdEncoding.EncodeToString(f)
	//text := "将文字信息转化为声音信息，给应用配上嘴巴。我们提供了众多极具特色的发音人供您选择。"

	data := url.Values{}
	data.Add("audio", audio)
	data.Add("text", text)
	body := data.Encode()

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://api.xfyun.cn/v1/service/v1/ise", strings.NewReader(body))
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	req.Header.Set("X-Appid", appid)
	req.Header.Set("X-CurTime", curtime)
	req.Header.Set("X-Param", base64_param)
	req.Header.Set("X-CheckSum", checksum)

	res, _ := client.Do(req)
	defer res.Body.Close()
	resp_body, _ := ioutil.ReadAll(res.Body)
	fmt.Print(string(resp_body))

	return resp_body
}
type Word struct {
	BegPos int64 `json:"beg_pos"`
	Content string `json:"content"`
	EndPos int64 `json:"end_pos"`
	Symbol string `json:"symbol"`
	TimeLen int64 `json:"time_len"`
	TotalScore string `json:"total_score"`
}

type Sentence struct {
	BegPos int64 `json:"beg_pos"`
	Content string `json:"content"`
	EndPos int64 `json:"end_pos"`
	TimeLen int64 `json:"time_len"`
	TotalScore float64 `json:"total_score"`
	Word Word `json:"word"`
}

type ReadWord struct {
	ExceptInfo string `json:"except_info"`
	IsRejected bool `json:"is_rejected"`
	TimeLen int64 `json:"time_len"`
	TotalScore float64 `json:"total_score"`
	Sentence []Sentence `json:"sentence"`
	BegPos int64 `json:"beg_pos"`
	Content string `json:"content"`
	EndPos int64 `json:"end_pos"`
}

type RecPaper struct {
	ReadWord ReadWord `json:"read_sentence"`
}

type ReadSentence struct {
	Lan string `json:"lan"`
	Type string `json:"type"`
	Version string `json:"version"`
	RecPaper RecPaper `json:"rec_paper"`
}

type EvaluationData struct {
	ReadSentence ReadSentence `json:"read_sentence"`
}

type EvaluationResponse struct {
	Code int64 `json:"code"`
	Desc string `json:"desc"`
	Sid string `json:"sid"`
	Data EvaluationData `json:"data"`
}