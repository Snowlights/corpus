package cache

import (
	"context"
	"fmt"
	"github.com/Snowlights/corpus/model/daoimpl"
	"log"
	"sync"
)

type cookie struct {
	mu sync.Mutex
	cookieList []string
	adminCookieList []string
}

var cookieList cookie

func AddCookieToList(cookie string) bool{
	fun := "AddCookieToList -->"
	cookieList.mu.Lock()
	defer cookieList.mu.Unlock()
	cookieList.cookieList = append(cookieList.cookieList,cookie)
	ListCookieList()
	log.Printf("%v start",fun)
	return true
}

func DelCookieFromList(cookie string) bool {
	fun := "DelCookieFromList -->"
	cookieList.mu.Lock()
	defer cookieList.mu.Unlock()

	for i,item := range cookieList.cookieList{
		if item == cookie{
			cookieList.cookieList = append(cookieList.cookieList[:i],cookieList.cookieList[i+1:]...)
			ListCookieList()
			return true
		}
	}
	ListCookieList()
	log.Printf("%v start",fun)
	return false
}

func CheckOnLine(cookie string) bool{
	cookieList.mu.Lock()
	defer cookieList.mu.Unlock()
	for _, item := range cookieList.cookieList{
		if item == cookie{
			return true
		}
	}
	ListCookieList()
	return false
}

func CheckIsAdmin(cookie string) bool{

	for _, item := range cookieList.adminCookieList{
		if item == cookie{
			return true
		}
	}
	return false
}


func ListCookieList(){
	fmt.Printf("ListCookieList---")
	if len(cookieList.cookieList) == 0{
		return
	}
	for _,item := range cookieList.cookieList{
		fmt.Printf("%v-",item)
	}
	fmt.Printf("\n")
}

func CheckUserAuth(ctx context.Context,code string, cookie string) bool{
	fun := "CheckUserAuth"

	conds := map[string]interface{}{
		"token" : cookie,
		"is_deleted" :false,
	}
	limit := map[string]interface{}{
		"limit" : 1,
		"offset" :0,
	}
	user,err := daoimpl.UserDao.ListUserInfo(ctx,limit,conds)
	if err != nil{
		log.Printf("%v %v error %v",ctx,fun,err)
		return false
	}
	if len(user) == 0{
		log.Printf("未找到用户信息")
		return false
	}

	userConds := map[string]interface{}{
		"user_id" : user[0].Id,
		"auth_code" : code,
		"is_deleted" : false,
	}
	userAuth,err := daoimpl.UserAuthDao.ListUserAuth(ctx,limit,userConds)
	if err != nil{
		log.Printf("%v %v error %v",ctx,fun,err)
		return false
	}
	if len(userAuth) > 0{
		return true
	}
	return false
}