package router


//type ProcGrpc struct{}
//
//func (m *ProcGrpc) Init() error {
//	fun := "ProcGrpc.Init -->"
//
//	log.Printf("%v success ",fun)
//	return nil
//}
//
//func (m *ProcGrpc) Driver() (string, interface{}) {
//	err := http.ListenAndServe("127.0.0.1",nil) //prot
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//	server := grpc.NewServer()
//	corpus.RegisterCorpusServiceServer(server, new(CorpusServiceTmp))
//	// 使用随机端口
//	reflection.Register(server)
//
//	return "", server
//}

