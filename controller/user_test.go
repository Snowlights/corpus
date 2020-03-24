package controller

import (
	"context"
	"github.com/Snowlights/corpus/cache"
	"github.com/Snowlights/corpus/model"
	"github.com/Snowlights/corpus/model/daoimpl"
	corpus "github.com/Snowlights/pub/grpc"
	"log"
	"reflect"
	"testing"
)

func initenv() context.Context{
	ctx := context.Background()
	model.Prepare(ctx)
	daoimpl.PrePare(ctx)
	cache.Prepare(ctx)
	return ctx
}

func TestLoginUser(t *testing.T) {
	ctx:= initenv()
	req := &corpus.LoginUserReq{
		EMail:                "858777157@qq.com",
		UserPassword:         "woaini12",
		Phone:                "",
		Code:                 "",
	}

	r := LoginUser(ctx,req)

	log.Printf("%v",r)

}

func TestDelUserInfo(t *testing.T) {
	type args struct {
		ctx context.Context
		req *corpus.DelUserInfoReq
	}
	tests := []struct {
		name string
		args args
		want *corpus.DelUserInfoRes
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DelUserInfo(tt.args.ctx, tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DelUserInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListUserInfo(t *testing.T) {
	ctx := initenv()
	req := &corpus.ListUserInfoReq{
		Offset:               0,
		Limit:                10,
		EMail:                "",
		UserName:             "",
		Phone:                "",
		Cookie:               "",
	}
	res := ListUserInfo(ctx,req)
	log.Printf("%v",res)
}

func TestLoginOutUserInfo(t *testing.T) {
	type args struct {
		ctx context.Context
		req *corpus.LogoutUserInfoReq
	}
	tests := []struct {
		name string
		args args
		want *corpus.LogoutUserInfoRes
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoginOutUserInfo(tt.args.ctx, tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginOutUserInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoginUser1(t *testing.T) {
	type args struct {
		ctx context.Context
		req *corpus.LoginUserReq
	}
	tests := []struct {
		name string
		args args
		want *corpus.LoginUserRes
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoginUser(tt.args.ctx, tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateUserInfo(t *testing.T) {
	type args struct {
		ctx context.Context
		req *corpus.UpdateUserInfoReq
	}
	tests := []struct {
		name string
		args args
		want *corpus.UpdateUserInfoRes
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateUserInfo(tt.args.ctx, tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateUserInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
