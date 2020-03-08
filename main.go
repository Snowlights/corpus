package main

import (
	"context"
	"log"
)

func InitFunc() (err error) {
	fun := "main.InitFunc -->"
	ctx := context.Background()


	// You can put more preparations here


	log.Println("%s usccess %s",ctx,fun)
	return nil
}

func main() {

	//ctx := context.Background()
	//// 给自己定义的procssor分配一个给力的名字
	//// 如果只有一个processor，就沿用这个命名吧，http的命名为proc_http, thrift 命名为proc_thrift
	//ps := map[string]rocserv.Processor{
	//	"proc_grpc": &router.ProcGrpc{},
	//}

	// 如果初始化过程发生错误，会直接panic
	// 正常情况这个调用会直接阻塞
	//err := rocserv.Serve(CLUSTER, BASE_LOCTION, InitFunc, ps)
	//if err != nil {
	//	slog.Errorf(ctx,"serve err:%s", err)
	//}
}