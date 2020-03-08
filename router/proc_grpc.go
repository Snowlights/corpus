package router

import (
	corpus "github.com/Snowlights/pub/grpc"
	rocserv "github.com/shawnfeng/roc/util/service"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	"log"
	"net/http"
)


type ProcGrpc struct{}
var clientGrpc *rocserv.ClientGrpc

func (m *ProcGrpc) Init() error {
	fun := "ProcGrpc.Init -->"

	log.Println("%v success ",fun)
	return nil
}

func (m *ProcGrpc) Driver() (string, interface{}) {
	err := http.ListenAndServe(":8000",nil) //prot
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	corpus.RegisterCorpusServiceServer(server, new(CorpusServiceTmp))
	// 使用随机端口
	reflection.Register(server)

	return "", server
}

