package cache

import (
	"context"
	"fmt"
	"github.com/Snowlights/corpus/model/daoimpl"
	"github.com/astaxie/beego/logs"
	"log"
	"sync"
	"time"
)

type Timer struct {
	code string
	timer int64
}

type MessageManager struct {
	mu sync.Mutex
	PhoneCode map[string]*Timer
}

var messageManage MessageManager

func Prepare(ctx context.Context){
	messageManage.PhoneCode = make(map[string]*Timer)
	cookieList.cookieList = make([]string,500)
	cookieList.adminCookieList = make([]string,500)
	go messageManage.syncTimer(ctx)
}
func (m *MessageManager) syncTimer(ctx context.Context){
SyncLoop:
	for {
		err := Reload(ctx)
		if err != nil{
			logs.Error(err)
		}
		err = ReoladAdmin(ctx)
		if err != nil{
			logs.Error(err)
		}

		select {
		case <-time.After(20*time.Second):
			log.Printf("sync message code")
		case <-ctx.Done():
			log.Printf("bale to exit ")
			break SyncLoop
		}
	}
}

func ReoladAdmin(ctx context.Context) error{

	var admin []string
	limit := map[string]interface{}{
		"offset" : 0,
		"limit" : 10,
	}
	conds := map[string]interface{}{
		"is_deleted":false,
	}
	dataList,err := daoimpl.AdminUserDao.ListAdminUser(ctx,limit,conds)
	if err != nil{
		logs.Error(err)
	}

	var idList []int64
	for _, item := range dataList{
		idList = append(idList,item.UserId)
	}
	userConds := map[string]interface{}{
		"id" : idList,
	}
	adminData, err := daoimpl.UserDao.ListUserInfo(ctx,limit,userConds)
	if err != nil{
		logs.Error(err)
	}

	for _, item := range adminData{
		admin = append(admin,item.Token)
	}
	cookieList.adminCookieList = admin
	log.Printf("%v",cookieList.adminCookieList)
	return nil
}


func Reload(ctx context.Context) error{
	messageManage.mu.Lock()
	defer messageManage.mu.Unlock()
	now := time.Now().Unix()

	for k, v := range messageManage.PhoneCode{
		if now > v.timer{
			delete(messageManage.PhoneCode,k)
		}
	}
	return nil
}

func AddPhoneCode(phone string,code string) bool{
	messageManage.mu.Lock()
	defer messageManage.mu.Unlock()
	now := time.Now().Unix()
	time := &Timer{
		code:  code,
		timer: now+300,
	}
	messageManage.PhoneCode[phone] = time
	ListPhoneCode()
	return true
}

func DelPhoneCode(phone string) bool{
	messageManage.mu.Lock()
	defer messageManage.mu.Unlock()

	delete(messageManage.PhoneCode,phone)
	ListPhoneCode()
	return true
}

func CheckPhoneCode(phone string,code string) bool{
	for k, v := range messageManage.PhoneCode{
		if k == phone && v.code == code{
			return true
		}
	}
	return false
}


func ListPhoneCode(){
	fmt.Printf("ListPhoneCode-----------------------\n" )
	for k,v := range messageManage.PhoneCode{
		fmt.Printf("%v  %v \n",k,v)
	}
}