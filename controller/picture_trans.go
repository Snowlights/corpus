package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/astaxie/beego/logs"
)

const (
	CompressingUrl = "https://api.tinify.com/shrink"

	// Email和ApiKey替换成自己的
	Email  = "wei1109942647@qq.com"
	ApiKey = "WBMs0rVL5gNHPcJZM6MTQKLLGFhGZNR4"
)

func init() {
	logs.SetLogFuncCall(true)
	logs.SetLogFuncCallDepth(3)
}

func picture_trans(file string) string{
	// 创建Request
	req, err := http.NewRequest(http.MethodPost, CompressingUrl, nil)
	if err != nil {
		logs.Error(err)
		return ""
	}

	// 将鉴权信息写入Request
	req.SetBasicAuth(Email, ApiKey)

	// 将图片以二进制的形式写入Request
	data, err := ioutil.ReadFile("C:\\Users\\华硕\\Desktop\\11.png")
	if err != nil {
		logs.Error(err)
		return ""
	}
	req.Body = ioutil.NopCloser(bytes.NewReader(data))

	// 发起请求
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		logs.Error(err)
		return ""
	}

	// 解析请求
	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		logs.Error(err)
		return ""
	}

	logs.Info(string(data))
	var pictureData PictureTransData
	err = json.Unmarshal(data,&pictureData)

	logs.Info(pictureData)

	return ""
}

type Input struct {
	Size int `json:"size"`
	Type string `json:"type"`
}

type Output struct {
	Size int `json:"size"`
	Type string `json:"type"`
	Width int `json:"width"`
	Height int `json:"height"`
	Ratio float64 `json:"ratio"`
	Url string `json:"url"`
}

type PictureTransData struct {
	Input Input `json:"input"`
	Output Output `json:"output"`
}