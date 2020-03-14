package cache

import (
	"context"
	"sync"
)

type MessageManager struct {
	mu sync.Mutex
	PhoneCode map[string]string
}

var messageManage MessageManager

func Prepare(ctx context.Context){
	messageManage.PhoneCode = make(map[string]string,1)
}

func AddPhoneCode(phone string,code string) bool{
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
