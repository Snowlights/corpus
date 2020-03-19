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

func analizeEvaluationData(data []byte,req *corpus.EvaluationReq) (map[string]interface{},*EvaluationResp){
	dataList := map[string]interface{}{}
	dataList["audio_src"] = req.Audio
	dataList["audio_text"] = req.Text
	now := time.Now().Unix()

	var res EvaluationResp
	err  := json.Unmarshal(data,&res)
	if err != nil{
		log.Printf("analizeEvaluationData failed ")
		return dataList,&res
	}
	if res.Data.ReadSentence.RecPaper.ReadChapter.TotalScore != ""{
		dataList["total_score"] = res.Data.ReadSentence.RecPaper.ReadChapter.TotalScore
	} else {
		dataList["total_score"] = 0
	}
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
	param["language"] = "en_us"
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

type Phone struct {
	BegPos string `json:"beg_pos"`
	EndPos string `json:"end_pos"`
	DpMessage string `json:"dp_message"`
	Content string `json:"content"`
}

type Syll struct {
	Phone []Phone `json:"phone"`
	BegPos string `json:"beg_pos"`
	EndPos string `json:"end_pos"`
	Content string `json:"content"`
	SyllAccent string `json:"syll_accent"`
	SyllScore string `json:"syll_score"`
}

type EvaluationWord struct {
	Index string `json:"index"`
	BegPos string `json:"beg_pos"`
	EndPos string `json:"end_pos"`
	DpMessage string `json:"dp_message"`
	Content string `json:"content"`
	GlobalIndex string `json:"global_index"`
	TotalScore string `json:"total_score"`
	Property string `json:"property"`
	Syll []Syll `json:"syll"`
}

type EvaluationSentence struct {
	Index string `json:"index"`
	BegPos string `json:"beg_pos"`
	EndPos string `json:"end_pos"`
	Content string `json:"content"`
	TotalScore string `json:"total_score"`
	WordCount string `json:"word_count"`
	Word []EvaluationWord `json:"word"`
}

type ReadChapter struct {
	BegPos string `json:"beg_pos"`
	EndPos string `json:"end_pos"`
	ExceptInfo string `json:"except_info"`
	IsRejected string `json:"is_rejected"`
	Content string `json:"content"`
	TotalScore string `json:"total_score"`
	WordCount string `json:"word_count"`
	Sentence []EvaluationSentence `json:"sentence"`
}

type RecPaper struct {
	ReadChapter ReadChapter `json:"read_chapter"`
}

type ReadSentence struct {
	RecPaper RecPaper `json:"rec_paper"`
	Lan string `json:"lan"`
	Type string `json:"type"`
	Version string `json:"version"`
}

type EvaluationData struct {
	ReadSentence ReadSentence `json:"read_sentence"`
}
type EvaluationResp struct {
	Data EvaluationData `json:"data"`
	Desc string `json:"desc"`
	Sid string `json:"sid"`
	Code string `json:"code"`
}