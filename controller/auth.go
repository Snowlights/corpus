package controller

import "sync"

type cookie struct {
	mu sync.Mutex
	cookieList []string
}

var cookieList cookie

func AddCookieToList(cookie string) bool{
	cookieList.mu.Lock()
	defer cookieList.mu.Unlock()

	cookieList.cookieList = append(cookieList.cookieList,cookie)
	return true
}

func DelCookieFromList(cookie string) bool {
	cookieList.mu.Lock()
	defer cookieList.mu.Unlock()
	for i,item := range cookieList.cookieList{
		if item == cookie{
			cookieList.cookieList = append(cookieList.cookieList[:i],cookieList.cookieList[i+1:]...)
		}
	}

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
	return false
}