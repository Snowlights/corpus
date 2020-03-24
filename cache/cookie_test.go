package cache

import (
	"context"
	"fmt"
	"github.com/Snowlights/corpus/model"
	"github.com/Snowlights/corpus/model/daoimpl"
	"github.com/astaxie/beego/logs"
	"testing"
	"time"
)
func initenv() context.Context{
	ctx := context.Background()
	model.Prepare(ctx)
	daoimpl.PrePare(ctx)
	Prepare(ctx)
	return ctx
}
func TestAddCookieToList(t *testing.T) {
	initenv()
	AddCookieToList("JnL3gxsI402j4hs4")
	AddCookieToList("ONkSy6HRfzRX7w1n")
	ListCookieList()
	DelCookieFromList("JnL3gxsI402j4hs4")
	fmt.Printf("--------------\n")
	ListCookieList()
}

func TestCheckOnLine(t *testing.T) {
	ctx := initenv()
	err := ReoladAdmin(ctx)
	if err != nil{
		logs.Error(err)
	}
}

func TestAddPhoneCode(t *testing.T) {
	initenv()
	AddPhoneCode("18846082154","111111")
	AddPhoneCode("15546023152","222222")
	AddPhoneCode("19999999999","333333")

	DelPhoneCode("18846082154")
	fmt.Printf("-------------")

	time.Sleep(time.Second)
}