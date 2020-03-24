package cache

import (
	"fmt"
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
	for _,item := range cookieList.cookieList{
		fmt.Printf("%v-",item)
	}
	fmt.Printf("\n")
}