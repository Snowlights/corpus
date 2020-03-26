package controller

import (
	"fmt"
	"testing"
)

func TestSendMessage(t *testing.T) {
	ctx := initenv()
	//req := &corpus.SendMessageReq{
	//	Phone:                "18846085051",
	//}
	code,err := message("18846082154")
	if err != nil{
		return
	}
	fmt.Printf("%v %v",ctx,code)
}