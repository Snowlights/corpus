package controller

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

/**
 * 语音听写流式 WebAPI 接口调用示例 接口文档（必看）：https://doc.xfyun.cn/rest_api/语音听写（流式版）.html 用于1分钟内的即时语音转文字技术
 * webapi 听写服务参考帖子（必看）：http://bbs.xfyun.cn/forum.php?mod=viewthread&tid=38947&extra=
 * 语音听写流式WebAPI 服务，热词使用方式：登陆开放平台https://www.xfyun.cn/后，找到控制台--我的应用---语音听写---服务管理--上传热词
 * 注意：热词只能在识别的时候会增加热词的识别权重，需要注意的是增加相应词条的识别率，但并不是绝对的，具体效果以您测试为准。
 * 错误码链接：https://www.xfyun.cn/document/error-code （code返回错误码时必看）
 * @author iflytek
 */
var (
	audioTextHostUrl   = "wss://iat-api.xfyun.cn/v2/iat"
	audioTextapiKey    = "c91c413b9f2df7d086dec0ba89a12223"
	audioTextapiSecret = "17c405f65fe9815e737e21b0183d9036"
	//file      = "C:\\Users\\wei11\\Desktop\\iat_pcm_16k.pcm" //请填写您的音频文件路径
	audioTextappid     = "5e1c39a3"
)

const (
	Audio_Text_STATUS_FIRST_FRAME    = 0
	Audio_Text_STATUS_CONTINUE_FRAME = 1
	Audio_Text_STATUS_LAST_FRAME     = 2
)

func audioText(file string) string{
	fmt.Println(HmacWithShaTobase64Picture("hmac-sha256","hello\nhello","hello"))
	st:=time.Now()
	d := websocket.Dialer{
		HandshakeTimeout: 5 * time.Second,
	}
	//握手并建立websocket 连接
	conn, resp, err := d.Dial(assembleAuthUrlPicture(audioTextHostUrl, audioTextapiKey, audioTextapiSecret), nil)
	if err != nil {
		panic(readResp(resp)+err.Error())
		return ""
	}else if resp.StatusCode !=101{
		panic(readResp(resp)+err.Error())
	}
	//打开音频文件

	var frameSize = 1280              //每一帧的音频大小
	var intervel = 40 * time.Millisecond //发送音频间隔
	//开启协程，发送数据
	ctx,_:= context.WithCancel(context.Background())
	defer conn.Close()
	var status = 0
	go func() {
		//	start:
		audioFile, err := os.Open(file)
		if err != nil {
			panic(err)
		}
		status = Audio_Text_STATUS_FIRST_FRAME      //音频的状态信息，标识音频是第一帧，还是中间帧、最后一帧
		//		time.Sleep(20*time.Second)
		var buffer = make([]byte, frameSize)
		for {
			len, err := audioFile.Read(buffer)
			if err != nil {
				if err == io.EOF { //文件读取完了，改变status = Audio_Text_STATUS_LAST_FRAME
					status = Audio_Text_STATUS_LAST_FRAME
				} else {
					panic(err)
				}
			}
			select {
			case <-ctx.Done():
				fmt.Println("session end ---")
				return
			default:
			}
			switch status {
			case Audio_Text_STATUS_FIRST_FRAME: //发送第一帧音频，带business 参数
				frameData := map[string]interface{}{
					"common": map[string]interface{}{
						"app_id": audioTextappid, //audioTextappid 必须带上，只需第一帧发送
					},
					"business": map[string]interface{}{ //business 参数，只需一帧发送
						"language":"en_us", //todo language
						"domain":"iat",
						"accent":"mandarin",
					},
					"data": map[string]interface{}{
						"status":    Audio_Text_STATUS_FIRST_FRAME,
						"format":    "audio/L16;rate=16000",
						"audio":     base64.StdEncoding.EncodeToString(buffer[:len]),
						"encoding":  "raw",
					},
				}
				fmt.Println("send first")
				conn.WriteJSON(frameData)
				status = Audio_Text_STATUS_CONTINUE_FRAME
			case Audio_Text_STATUS_CONTINUE_FRAME:
				frameData := map[string]interface{}{
					"data": map[string]interface{}{
						"status":    Audio_Text_STATUS_CONTINUE_FRAME,
						"format":    "audio/L16;rate=16000",
						"audio":     base64.StdEncoding.EncodeToString(buffer[:len]),
						"encoding":  "raw",
					},
				}
				conn.WriteJSON(frameData)
			case Audio_Text_STATUS_LAST_FRAME:
				frameData := map[string]interface{}{
					"data": map[string]interface{}{
						"status":    Audio_Text_STATUS_LAST_FRAME,
						"format":    "audio/L16;rate=16000",
						"audio":     base64.StdEncoding.EncodeToString(buffer[:len]),
						"encoding":  "raw",
					},
				}
				conn.WriteJSON(frameData)
				fmt.Println("send last")
				return
				//	goto start
			}

			//模拟音频采样间隔
			time.Sleep(intervel)
		}

	}()
	res := ""
	//获取返回的数据
	//var decoder Decoder
	for {
		var resp = PictureRespData{}
		_,msg,err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read message error:", err)
			break
		}
		json.Unmarshal(msg,&resp)

		//fmt.Println(string(msg))
		fmt.Println(resp.Data.Result.String(),resp.Sid)
		res = res+resp.Data.Result.String()
		if resp.Code!=0{
			fmt.Println(resp.Code,resp.Message,time.Since(st))
			return ""
		}
		//decoder.Decode(&resp.Data.Result)
		if resp.Data.Status == 2{
			//cf()
			//fmt.Println("final:",decoder.String())
			fmt.Println(resp.Code,resp.Message,time.Since(st))
			break
			//return
		}

	}

	time.Sleep(1*time.Second)
	return res
}

type PictureRespData struct {
	Sid 	string `json:"sid"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    PictureData  `json:"data"`
}

type PictureData struct {
	Result PictureResult `json:"result"`
	Status int         `json:"status"`
}

//创建鉴权url  audioTextapiKey 即 hmac username
func assembleAuthUrlPicture(audioTextHostUrl string, audioTextapiKey, audioTextapiSecret string) string {
	ul, err := url.Parse(audioTextHostUrl)
	if err != nil {
		fmt.Println(err)
	}
	//签名时间
	date := time.Now().UTC().Format(time.RFC1123)
	//date = "Tue, 28 May 2019 09:10:42 MST"
	//参与签名的字段 host ,date, request-line
	signString := []string{"host: " + ul.Host, "date: " + date, "GET " + ul.Path + " HTTP/1.1"}
	//拼接签名字符串
	sgin := strings.Join(signString, "\n")
	fmt.Println(sgin)
	//签名结果
	sha := HmacWithShaTobase64Picture("hmac-sha256", sgin, audioTextapiSecret)
	fmt.Println(sha)
	//构建请求参数 此时不需要urlencoding
	authUrl := fmt.Sprintf("hmac username=\"%s\", algorithm=\"%s\", headers=\"%s\", signature=\"%s\"", audioTextapiKey,
		"hmac-sha256", "host date request-line", sha)
	//将请求参数使用base64编码
	authorization := base64.StdEncoding.EncodeToString([]byte(authUrl))

	v := url.Values{}
	v.Add("host", ul.Host)
	v.Add("date", date)
	v.Add("authorization", authorization)
	//将编码后的字符串url encode后添加到url后面
	callurl := audioTextHostUrl + "?" + v.Encode()
	return callurl
}

func HmacWithShaTobase64Picture(algorithm, data, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	encodeData := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(encodeData)
}

func readResp(resp *http.Response) string {
	if resp == nil {
		return ""
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("code=%d,body=%s", resp.StatusCode, string(b))
}
// 解析返回数据，仅供demo参考，实际场景可能与此不同。
type Decoder struct {
	results []*PictureResult
}

func (d *Decoder)Decode(result *PictureResult)  {
	if len(d.results)<=result.Sn{
		d.results = append(d.results,make([]*PictureResult,result.Sn-len(d.results)+1)...)
	}
	if result.Pgs == "rpl"{
		for i:=result.Rg[0];i<=result.Rg[1];i++{
			d.results[i]=nil
		}
	}
	d.results[result.Sn] = result
}

func (d *Decoder)String()string  {
	var r string
	for _,v:=range d.results{
		if v== nil{
			continue
		}
		r += v.String()
	}
	return r
}

type PictureResult struct {
	Ls bool `json:"ls"`
	Rg []int `json:"rg"`
	Sn int `json:"sn"`
	Pgs string `json:"pgs"`
	Ws []Ws `json:"ws"`
}

func (t *PictureResult)String() string {
	var wss string
	for _,v:=range t.Ws{
		wss+=v.String()
	}
	return wss
}

type Ws struct {
	Bg int `json:"bg"`
	Cw []Cw `json:"cw"`
}

func (w *Ws)String()string  {
	var wss string
	for _,v:=range w.Cw{
		wss+=v.W
	}
	return wss
}
type Cw struct {
	Sc int `json:"sc"`
	W string `json:"w"`
}