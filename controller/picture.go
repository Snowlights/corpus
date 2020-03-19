package controller

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//百度api
func handwriting_baidu(file string) *BaiduPicture{
	// 应用APPID(必须为webapi类型应用,并开通手写文字识别服务,参考帖子如何创建一个webapi应用：http://bbs.xfyun.cn/forum.php?mod=viewthread&tid=36481)
	//appid := "18926851"
	// 接口密钥(webapi类型应用开通手写文字识别后，控制台--我的应用---手写文字识别---相应服务的apikey)
	apikey := "dINzyryp6MYCqu2LbcFpyRKU"
	curtime := strconv.FormatInt(time.Now().Unix(), 10)
	param := make(map[string]string)
	// 语种参数和文本位置参数
	param["language_type"] = "CHN_ENG"
	param["detect_direction"] = "true"
	param["detect_language"] = "true"
	param["paragraph"] = "true"
	param["probability"] = "true"
	tmp, _ := json.Marshal(param)
	base64_param := base64.StdEncoding.EncodeToString(tmp)
	w := md5.New()
	io.WriteString(w, apikey+curtime+base64_param)
	checksum := fmt.Sprintf("%x", w.Sum(nil))
	//测试图片地址
	//f, _ := ioutil.ReadFile("C:\\Users\\wei11\\Desktop\\test1.png")
	f, _ := ioutil.ReadFile(file)
	f_base64 := base64.StdEncoding.EncodeToString(f)
	data := url.Values{}
	data.Add("image", f_base64)
	body := data.Encode()
	client := &http.Client{}
	// 手写文字识别OCR接口地址
	req, _ := http.NewRequest("POST", "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic?access_token=24.29d1e6601544f1b0945791a5e270f88e.2592000.1587187332.282335-18926851", strings.NewReader(body))
	//req.Header.Set("X-Appid", appid)
	req.Header.Set("X-CurTime", curtime)
	req.Header.Set("X-Param", base64_param)
	req.Header.Set("X-CheckSum", checksum)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, _ := client.Do(req)
	defer res.Body.Close()
	res_body, _ := ioutil.ReadAll(res.Body)
	fmt.Print(string(res_body))
	var BaiduPicture BaiduPicture
	err := json.Unmarshal(res_body,&BaiduPicture)
	if err != nil{
		logs.Error(err)
	}
	fmt.Printf("%v",BaiduPicture)
	return &BaiduPicture
}

type BaiDuWord struct {
	Words string `json:"words"`
}

type BaiduPicture struct {
	LogId int64 `json:"log_id"`
	WordsResultNum int `json:"words_result_num"`
	WordsResult []BaiDuWord `json:"words_result"`
}

func handwriting_xunfei(file string) *XfPictureResp{
	// 应用APPID(必须为webapi类型应用,并开通手写文字识别服务,参考帖子如何创建一个webapi应用：http://bbs.xfyun.cn/forum.php?mod=viewthread&tid=36481)
	appid := "5e1c39a3"
	// 接口密钥(webapi类型应用开通手写文字识别后，控制台--我的应用---手写文字识别---相应服务的apikey)
	apikey := "3e8fe2c454c339a7c969a21d1e9e44fe"
	curtime := strconv.FormatInt(time.Now().Unix(), 10)
	param := make(map[string]string)
	// 语种参数和文本位置参数
	param["language"] = "cn|en"
	param["location"] = "true"
	tmp, _ := json.Marshal(param)
	base64_param := base64.StdEncoding.EncodeToString(tmp)
	w := md5.New()
	io.WriteString(w, apikey+curtime+base64_param)
	checksum := fmt.Sprintf("%x", w.Sum(nil))
	// 测试图片地址
	//f, _ := ioutil.ReadFile("./tt_1.jpg")
	f, _ := ioutil.ReadFile(file)
	f_base64 := base64.StdEncoding.EncodeToString(f)
	data := url.Values{}
	data.Add("image", f_base64)
	body := data.Encode()
	client := &http.Client{}
	// 手写文字识别OCR接口地址
	req, _ := http.NewRequest("POST", "http://webapi.xfyun.cn/v1/service/v1/ocr/handwriting", strings.NewReader(body))
	req.Header.Set("X-Appid", appid)
	req.Header.Set("X-CurTime", curtime)
	req.Header.Set("X-Param", base64_param)
	req.Header.Set("X-CheckSum", checksum)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, _ := client.Do(req)
	defer res.Body.Close()
	res_body, _ := ioutil.ReadAll(res.Body)
	fmt.Print(string(res_body))
	var XfPictureResp XfPictureResp
	err := json.Unmarshal(res_body,&XfPictureResp)
	if err != nil{
		logs.Error(err)
	}
	fmt.Printf("%v",XfPictureResp)

	return &XfPictureResp
}

type XfTopLeft struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type XfRightBottom struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type XfLocation struct {
	TopLeft XfTopLeft `json:"top_left"`
	RightBottom XfRightBottom `json:"right_bottom"`
}

type XfContent struct {
	Content string `json:"content"`
}

type XfLine struct {
	Confidence int `json:"confidence"`
	Location XfLocation `json:"location"`
	Word []XfContent `json:"word"`
}

type Block struct {
	Type string `json:"type"`
	Line []XfLine `json:"line"`
}

type XfPictureData struct {
	Block []Block `json:"block"`
}

type XfPictureResp struct {
	Code string `json:"code"`
	Data XfPictureData `json:"data"`
	Desc string `json:"desc"`
	Sid string `json:"sid"`
}