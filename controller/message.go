package controller

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/Snowlights/corpus/cache"
	corpus "github.com/Snowlights/pub/grpc"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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
	audit := formMessageAudit(req)
	auditLastInsertId, err := addAudit(ctx,audit)
	if err!= nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
	}
	log.Printf("%v %v success ,auditLastInsertId %d",ctx,fun,auditLastInsertId)

	log.Printf("%v %v success",ctx,fun)
	return res
}

func formMessageAudit(req *corpus.SendMessageReq) map[string]interface{}{
	now := time.Now().Unix()
	audit := map[string]interface{}{
		"table_name" : "user_login_out",
		"content" : fmt.Sprintf("send phone code"),
		"history" : fmt.Sprintf("send phone code time %v phone %v",now,req.Phone),
		"activity" : fmt.Sprintf("%v phone code send",req.Phone),
		"created_at" : now,
		"created_by" : req.Phone,
		"is_deleted" : false,
	}
	return audit
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

	cache.AddPhoneCode(_mobile,vcode)
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

