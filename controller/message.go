package controller

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	corpus "github.com/Snowlights/pub/grpc"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

func SendMessage(ctx context.Context,req* corpus.SendMessageReq) *corpus.SendMessageRes{
	fun := "Controller.SendMessage -- >"
	res := &corpus.SendMessageRes{}

	err := message(req.Phone)
	if err!= nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  err.Error(),
		}
		return res
	}

	return res
}

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
func message(mobile string) (error){
	v := url.Values{}
	_now := strconv.FormatInt(time.Now().Unix(), 10)
	//fmt.Printf(_now)
	_account := "C76170131"//查看用户名 登录用户中心->验证码通知短信>产品总览->API接口信息->APIID
	_password := "ea3d48374ad86241459043765313b45b" //查看密码 登录用户中心->验证码通知短信>产品总览->API接口信息->APIKEY
	_mobile := mobile
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))

	AddPhoneCode(_mobile,vcode)
	_content := fmt.Sprintf("您的验证码是：%v。请不要把验证码泄露给其他人。",vcode)
	v.Set("account", _account)
	v.Set("password", GetMd5String(_account+_password+_mobile+_content+_now))
	v.Set("mobile", _mobile)
	v.Set("content", _content)
	v.Set("time", _now)
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://106.ihuyi.com/webservice/sms.php?method=Submit&format=json", body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//fmt.Printf("%+v\n", req) //看下发送的结构

	resp, err := client.Do(req) //发送
	if err!= nil{
		return err
	}
	defer resp.Body.Close()     //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)
	return nil
}

type MessageManager struct {
	mu sync.Mutex
	PhoneCode map[string]string
}

var messageManage MessageManager

func Prepare(ctx context.Context){
	messageManage.PhoneCode = make(map[string]string,1)
}

func AddPhoneCode (phone string,code string) bool{
	messageManage.mu.Lock()
	defer messageManage.mu.Unlock()

	messageManage.PhoneCode[phone] = code
	return true
}

func DelPhoneCode(phone string) bool{
	messageManage.mu.Lock()
	defer messageManage.mu.Unlock()

	delete(messageManage.PhoneCode,phone)

	return true
}