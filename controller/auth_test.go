package controller

import (
	"fmt"
	corpus "github.com/Snowlights/pub/grpc"
	"testing"
)

func Test_AddUserAuth(t *testing.T) {
	ctx := initenv()
	req := &corpus.AddUserAuthReq{
		UserId:               4,
		AuthCode:             "SERVICE_AUDIO_CODE",
		Cookie:               "JnL3gxsI402j4hs4",
	}
	res := AddUserAuth(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}
}

func TestListUserAuth(t *testing.T) {
	ctx := initenv()
	req := &corpus.ListUserAuthReq{
		UserId:               4,
		Limit:                10,
		Offset:               0,
		Cookie:               "",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	r := ListUserAuth(ctx,req)
	fmt.Printf("%v",r)
}

func TestUpdateAuth(t *testing.T) {
	ctx := initenv()
	req := &corpus.UpdateAuthReq{
		Id:                   1,
		AuthCode:             "SERVICE_AUDIO_CODE",
		AuthDescription:      "音频",
		ServiceName:          "/audio",
		Cookie:               "JnL3gxsI402j4hs4",
	}
	res := UpdateAuth(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v",res.Errinfo.Msg)
	}

}