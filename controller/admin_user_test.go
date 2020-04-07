package controller

import (
	"context"
	corpus "github.com/Snowlights/pub/grpc"
	"log"
	"reflect"
	"testing"
)

func TestAddAdminUser(t *testing.T) {
	ctx := initenv()
	req := &corpus.AddAdminUserReq{
		UserId:               7,
		Cookie:               "zhangwei",
	}
	res := AddAdminUser(ctx,req)
	log.Printf("%v",res)
	if res.Errinfo!= nil{
		log.Printf("%v",res.Errinfo.Msg)
	}

}

func TestDelAdminUser(t *testing.T) {
	type args struct {
		ctx context.Context
		req *corpus.DelAdminUserReq
	}
	tests := []struct {
		name string
		args args
		want *corpus.DelAdminUserRes
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DelAdminUser(tt.args.ctx, tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DelAdminUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListAdminUser(t *testing.T) {
	ctx := initenv()
	req := &corpus.ListAdminUserReq{
		Limit:                10,
		Offset:               0,
		Cookie:               "",
	}
	res := ListAdminUser(ctx,req)

	log.Printf("%v",res)
}

func Test_toAddAdminUser(t *testing.T) {
	type args struct {
		ctx context.Context
		req *corpus.AddAdminUserReq
	}
	tests := []struct {
		name  string
		args  args
		want  map[string]interface{}
		want1 map[string]interface{}
		want2 map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := toAddAdminUser(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toAddAdminUser() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("toAddAdminUser() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("toAddAdminUser() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_toDelAdminUser(t *testing.T) {
	type args struct {
		ctx context.Context
		req *corpus.DelAdminUserReq
	}
	tests := []struct {
		name  string
		args  args
		want  map[string]interface{}
		want1 map[string]interface{}
		want2 map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := toDelAdminUser(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toDelAdminUser() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("toDelAdminUser() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("toDelAdminUser() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_toListAdminUser(t *testing.T) {
	type args struct {
		ctx context.Context
		req *corpus.ListAdminUserReq
	}
	tests := []struct {
		name  string
		args  args
		want  map[string]interface{}
		want1 map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := toListAdminUser(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toListAdminUser() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("toListAdminUser() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}