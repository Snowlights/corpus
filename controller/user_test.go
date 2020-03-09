package controller

import (
	"context"
	"fmt"
	"github.com/Snowlights/corpus/model"
	"log"
	"math/rand"
	"testing"
	"time"
)

func initenv() context.Context{
	ctx := context.Background()
	model.Prepare(ctx)

	return ctx
}

func TestLoginUser(t *testing.T) {
	//ctx:= initenv()
	//req := &corpus.LoginUserReq{
	//	EMail:                "wei1109942647",
	//	UserPassword:         "",
	//	Phone:                "",
	//	Code:                 0,
	//}
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	log.Println(vcode)

	//LoginUser(ctx,req)
}