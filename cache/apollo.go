package cache

import (
	"context"
	"fmt"
	"github.com/ZhengHe-MD/agollo/v4"
	"github.com/ZhengHe-MD/properties"
	"github.com/astaxie/beego/logs"
	"log"
	"sync"
)
func apollo(ctx context.Context){
	conf := &agollo.Conf{
		AppID:          "corpus",
		Cluster:        "default",
		NameSpaceNames: []string{"application"},
		CacheDir:       "/tmp/agollo",
		IP:             "localhost:8080",
	}
	err := agollo.StartWithConf(conf)
	if err != nil{
		log.Println(err)
	}
	recall := agollo.RegisterObserver(&observer{})
	defer recall()

	err = LoadCorpusConfig(ctx)
	if err != nil{
		log.Print(err)
	}
}

type CorpusConfig struct {
	mu sync.Mutex
	Admin []string `properties:"admin"`
}

var corpusCfg CorpusConfig

func LoadCorpusConfig(ctx context.Context) (err error){
	var corpusConfig CorpusConfig

	if err = Unmarshal(ctx,&corpusConfig);err != nil{
		logs.Error(err)
	}
	corpusCfg.mu.Lock()
	defer corpusCfg.mu.Unlock()
	corpusCfg.Admin = corpusConfig.Admin

	return
}

func testout(){
	fmt.Print(corpusCfg.Admin)
}

func Unmarshal(ctx context.Context,v interface{}) error{
	var kv = map[string]string{}
	data := agollo.GetAllKeys("application")
	for _, k := range data{
		if v, ok := agollo.GetStringWithNamespace("application",k); ok{
			kv[k] = v
		}
	}
	return properties.UnmarshalKV(kv, v)
}

func CheckSuperAdmin(ctx context.Context,cookie string) bool{
	corpusCfg.mu.Lock()
	defer corpusCfg.mu.Unlock()
	for _, item := range corpusCfg.Admin{
		if item == cookie{
			return true
		}
	}
	return false
}

type observer struct {
}

func (m *observer) HandleChangeEvent(ce *agollo.ChangeEvent){
	ctx := context.Background()
	log.Printf("%v start to pull data from apollo \n",ctx)
	_ = LoadCorpusConfig(ctx)
}

