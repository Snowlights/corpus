package controller

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
	"time"
)

var (
	hostUrl   = "wss://ws-api.xfyun.cn/v2/igr"
	apiKey    = "c91c413b9f2df7d086dec0ba89a12223"
	apiSecret = "17c405f65fe9815e737e21b0183d9036"
	//file      = "C:\\Users\\wei11\\Desktop\\igr_pcm_16k.pcm"
	appid     = "5e1c39a3"
)
const (
	STATUS_FIRST_FRAME    = 0
	STATUS_CONTINUE_FRAME = 1
	STATUS_LAST_FRAME     = 2
	HttpCodeSuccessHandshake = 101  //握手成功返回的httpcode
)
func recognizeAge(file string) *RespData{
	d := websocket.Dialer{
		HandshakeTimeout: 1 * time.Second,
	}
	//握手并建立websocket 连接
	conn, resp, err := d.Dial(assembleAuthUrl(hostUrl, apiKey, apiSecret), nil)
	if err != nil {
		panic(err)
		if resp.StatusCode != HttpCodeSuccessHandshake {
			b, _ := ioutil.ReadAll(resp.Body)
			fmt.Printf("handshake failed:message=%s,httpCode=%d\n", string(b), resp.StatusCode)
		}
		return nil
	}
	//打开音频文件
	audioFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	//开启协程，发送数据
	go func() {
		var frameSize = 1280                 //每一帧的音频大小
		var intervel = 40 * time.Millisecond //发送音频间隔
		var status = STATUS_FIRST_FRAME      //音频的状态信息，标识音频是第一帧，还是中间帧、最后一帧
		var buffer = make([]byte, frameSize)
		for {
			len, err := audioFile.Read(buffer)
			if err != nil {
				if err == io.EOF { //文件读取完了，改变status = STATUS_LAST_FRAME
					status = STATUS_LAST_FRAME
				} else {
					panic(err)
				}
			}
			switch status {
			case STATUS_FIRST_FRAME: //发送第一帧音频，带business 参数
				frameData := map[string]interface{}{
					"common": map[string]interface{}{
						"app_id": appid, //appid 必须带上，只需第一帧发送
					},
					"business": map[string]interface{}{ //business 参数，只需一帧发送
						"rate": 16000,
						"aue": "raw",
					},
					"data": map[string]interface{}{
						"status": STATUS_FIRST_FRAME, //第一帧音频status要为 0
						"audio":  base64.StdEncoding.EncodeToString(buffer[:len]),
					},
				}
				conn.WriteJSON(frameData)
				status = STATUS_CONTINUE_FRAME
			case STATUS_CONTINUE_FRAME:
				frameData := map[string]interface{}{
					"data": map[string]interface{}{
						"status": STATUS_CONTINUE_FRAME, // 中间音频status 要为1
						"audio":  base64.StdEncoding.EncodeToString(buffer[:len]),
					},
				}
				conn.WriteJSON(frameData)
			case STATUS_LAST_FRAME:
				frameData := map[string]interface{}{
					"data": map[string]interface{}{
						"status": STATUS_LAST_FRAME, // 最后一帧音频status 一定要为2 且一定发送
						"audio":  base64.StdEncoding.EncodeToString(buffer[:len]),
					},
				}
				conn.WriteJSON(frameData)
				goto end
			}
			//模拟音频采样间隔
			time.Sleep(intervel)
		}
	end:
	}()
	//获取返回的数据
	var resData *RespData
	for {
		var resp = &RespData{}
		err := conn.ReadJSON(resp)
		if err != nil {
			fmt.Println("read message error:", err)
			break
		}
		fmt.Println(resp)
		if resp.Code == 0 {
			if resp.Data != nil {
				if result:=resp.Data.Result ;result!= nil {
					fmt.Printf("result is :%+v \n",result)
					// todo
				}

				if resp.Data.Status == 2 { //当返回的数据status=2时，表示数据已经全部返回，这时候应该结束本次会话
					resData = resp
					break
				}
			}
		} else {
			fmt.Println("Error:",resp.Code, "|", resp.Message)
		}
		conn.Close()
		return resp
	}
	conn.Close()
	return resData
}
type RespData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *Data  `json:"data"`
}
type Data struct {
	Result *Result `json:"result"`
	Status int     `json:"status"`
}
type Result struct {
	Age    Age    `json:"age"`
	Gender Gender `json:"gender"`
}
type Age struct {
	AgeType string `json:"age_type"`
	Child   string `json:"child"`
	Middle  string `json:"middle"`
	Old     string `json:"old"`
}
type Gender struct {
	Female      string `json:"female"`
	Gender_type string `json:"gender_type"`
	Male        string `json:"male"`
}
//创建鉴权url
func assembleAuthUrl(hosturl string, apiKey, apiSecret string) string {
	ul, err := url.Parse(hosturl)
	if err != nil {
		fmt.Println(err)
	}
	//签名时间
	date := time.Now().UTC().Format(time.RFC1123)
	//参与签名的字段 host ,date, request-line
	signString := []string{"host: " + ul.Host, "date: " + date, "GET " + ul.Path + " HTTP/1.1"}
	//拼接签名字符串
	sgin := strings.Join(signString, "\n")
	//签名结果
	sha := HmacWithShaTobase64("hmac-sha256", sgin, apiSecret)
	//构建请求参数 authorization
	authUrl := fmt.Sprintf("api_key=\"%s\", algorithm=\"%s\", headers=\"%s\", signature=\"%s\"", apiKey,
		"hmac-sha256", "host date request-line", sha)
	//将请求参数使用base64编码
	authorization:= base64.StdEncoding.EncodeToString([]byte(authUrl))
	v := url.Values{}
	v.Add("host", ul.Host)
	v.Add("date", date)
	v.Add("authorization", authorization)
	//将编码后的字符串url encode后添加到url后面
	callurl := hosturl + "?" + v.Encode()
	return callurl
}
func HmacWithShaTobase64(algorithm, data, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	encodeData := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(encodeData)
}