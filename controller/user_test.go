package controller

import (
	"context"
	"fmt"
	"github.com/Snowlights/corpus/cache"
	"github.com/Snowlights/corpus/model"
	"github.com/Snowlights/corpus/model/daoimpl"
	"github.com/Snowlights/pub/adapter"
	corpus "github.com/Snowlights/pub/grpc"
	"github.com/astaxie/beego/logs"
	"log"
	"reflect"
	"syscall"
	"testing"
	"time"
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

func TestRecognizeAge(t *testing.T) { //不能超过10s
	ctx := initenv()
	req := &corpus.RecognizeAgeReq{
		Audio:                "C:\\Users\\华硕\\Desktop\\pr\\evaluation\\aa\\Fido Gets Dressed\\raz_fidogetsdressed_p3_text.mp3",
		Cookie:               "zhangwei",
	}
	res := RecognizeAge(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}
}

func TestSaveImage2(t *testing.T) {
	SaveImage("https://ss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=169785117,3207160551&fm=26&gp=0.jpg")
}

func TestRecognizeImage(t *testing.T) {
	ctx := initenv()

	//File :="https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=4251611548,332805913&fm=11&gp=0.jpg"
	//file := "https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=4251611548,332805913&fm=11&gp=0.jpg"
	req := &corpus.RecognizeImageReq{
		File:                 "https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=4251611548,332805913&fm=11&gp=0.jpg",
		Cookie: "zhangwei",
	}
	res := RecognizeImage(ctx,req)
	fmt.Printf("1111111\n")
	fmt.Printf("%v",res)

	time.Sleep(time.Second*30)
}


func TestTransAudioToText(t *testing.T) {
	ctx := initenv()
	req := &corpus.TransAudioToTextReq{
		Audio:                "http://xia2.kekenet.com/Sound/2015/11/Nov25_5700351F4Y.mp3",//"http://xia2.kekenet.com/Sound/2015/11/16_1710970T6a.mp3",//"http://mp3.en8848.com/kouyu/240huihua/13-4.mp3",
		Cookie:               "zhangwei",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	res := TransAudioToText(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}
	time.Sleep(time.Second)
}

func TestDelTransAudio(t *testing.T) {
	ctx := initenv()

	out,err := transformToPCM("http://xia2.kekenet.com//Sound/yousheng/Story365/ns365.mp3")
	if err != nil{
		logs.Error(err)
	}
	fmt.Printf("%v %v\n",ctx,out)

}

func TestEvaluation(t *testing.T) {
	ctx :=initenv()
	cache.AddCookieToList("lihua")
	req := &corpus.EvaluationReq{
		Audio:                "C:\\Users\\华硕\\Desktop\\pr\\evaluation\\aa\\In\\raz_in_p9_text.mp3",
		Text:                 "in the mud.",
		Cookie:               "lihua",
	}
	res := Evaluation(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}

}

func TestGetKeyWord(t *testing.T) {
	ctx := initenv()
	req := &corpus.GetKeyWordReq{
		Text:   "C:\\Users\\华硕\\Desktop\\keyreq.txt",
		Cookie:               "zhangwei",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	res := GetKeyWord(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}

}

func TestListTransAudio(t *testing.T) {
	ctx := initenv()
	req := &corpus.AddTransAudioReq{
		OriginAudio:          "http://xia2.kekenet.com/Sound/2015/11/Nov25_5700351F4Y.mp3",
		AudioType:            corpus.AudioType_ACC,
		Cookie:               "zhangwei",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	res := AddTransAudio(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}
}

func TestGetKeyWord2(t *testing.T) {
	ctx := initenv()
	req := &corpus.GetKeyWordReq{
		Text:                 "谁的人生没有一场声势浩大的暗恋呢？当爱情的种子在懵懂的心中萌了芽，它就会不顾暴雨狂风恣意生长，它是这样顽强的东西，除非连着心头肉连根拔去，不然就不会停止生长，任凭多少时间的消磨，它始终在那里，只需一个眼神就能傲雪盛放。简琛从七岁起就认识顾晓楠了，在简琛眼里顾晓楠是个笨蛋，在别的女孩子都喜欢扎辫子玩洋娃娃的年纪她偏要起早贪黑地去和男孩子争沙地玩泥巴。简琛最讨厌脏兮兮的泥巴地，每次在小区里碰到顾晓楠灰头土脸，简琛总是在心中好一番嫌弃。可一到饭点还是免不了要乖乖到顾晓楠家里去吃饭。简琛的父母和顾晓楠的妈妈是初中同学，经常出差，所以难得有时间照顾简琛，于是顾晓楠的妈妈就把简琛当半个儿子一样顺带养了。顾晓楠的妈妈常常教育自家女儿：“你看看人家简琛，一样的年纪，怎么你就不像人那么让爸妈省心呢。”顾晓楠看看在饭桌一边安静扒拉米饭的简琛，无奈地摇了摇头，在七岁的顾晓楠眼里，做人最重要的就是硬气，就像挖泥巴这种事情一样，不能因为王二小多挖一勺就轻易认输。可简琛真是太不硬气了，做什么事都柔柔弱弱的像个女孩子。彼时的顾晓楠还是颇想照顾这个白白瘦瘦的小男孩的，毕竟她向来重情义。顾晓楠扯了扯简琛的衣袖，“简琛，你要不要和我一起去沙地？我带你玩！我玩得最好了！”简琛看了一眼面颊上还带着灰的顾晓楠在内心绝望地叹了口气，“不，不了吧，我看书就好了。”于是七岁那年的友谊便如此产生了裂缝。顾晓楠常常怀疑，简琛是不是不喜欢和人一块儿呆着，十二岁那年顾晓楠甚至郑重其事地问过妈妈，简琛是不是有孤僻症或者自闭倾向，当然顾母没好气地翻了自家女儿一个白眼。但有一次顾晓楠看到简琛和他们班的一个女生讲话，讨论着令人摸不着头脑的奥数竞赛题目，顾晓楠生平最讨厌奥数，五花八门的题目从填空到应用题她都不感冒。她不关心他们谈的题目，只是这时候顾晓楠才明白，原来简琛不是不喜欢和人呆着，他只是不喜欢和自己一起玩罢了。十二岁的顾晓楠似乎有一些难过，但做人嘛最重要的是硬气，输什么都不能输架势。就算他简琛明摆着不喜欢和自己玩，那又有什么关系呢？没有简琛顾晓楠自然还有一大群狐朋狗友，班里面想跟着自己玩儿的人可多着呢。她依旧没心没肺地我行我素，只是再也不主动找简琛一起去街上周记小馄饨吃麻辣小馄饨，也不去找他一道放学后顺路去王阿姨的零食店买一块钱的冰棍儿，甚至连碰到令人头疼的数学题都懒得找简琛帮忙写。在十二岁的顾晓楠眼里，这就是一个人硬气的最佳证明。简琛每日仍旧同顾晓楠一起吃饭，一起上学，一起放学，他觉得顾晓楠一切的变化都不能按照常理去推断，与其揣测不如顺其自然，毕竟按照以往的规律，顾晓楠若是和人生气吵架，特别是和自己生气的话，只要保持安静，顾晓楠就气不过三天。而恰巧简琛最擅长做的事情就是保持安静，或者说长久的宛如失去生命一般的沉默。于是简琛安静地吃饭，安静地做作业，安静地上学。他并不需要怀疑自己安静的理由，只是他并不知道，顾晓楠这一次的生气原来会需要这么久。高三下学期，正值四月，春风和煦，吹得人困倦无比。城南的海棠落了一地，小区里三四岁的顽童兀自追逐，玩着经久不衰拾泥巴的游戏。顾晓楠坐在自家那栋楼下的长椅上，望着一群无忧无虑的小孩儿很是忧伤。联考的成绩单躺在皱巴巴的书包里，随之而来的还有一月一次的家长会。虽然这一次她确实小有进步，但终究还是在班里吊了车尾，而该死的简琛竟然又稳当当地名列前茅，仿佛脑门儿上刻了理科状元，金榜题名的大字。暖洋洋的阳光洒在水泥地上，金色让人温暖也让人困倦，快要吃中饭了，顾晓楠一边烦恼着，一边打哈欠，一边忍受着饥饿。远远的看见拐角处简琛骑着自行车驶过来，春季校服松松垮垮的搭在肩上，清爽的蓝白色还有清爽的黑色短发，在和煦的春风里飘动着，顾晓楠呆呆地看着，心想班里面这么多女生喜欢简琛也不是没道理的，还没缓过神来，简琛已经把车停在了顾晓楠面前。“还没吃饭？”简琛抬眼看了一脸痴呆的顾晓楠。“你又没回来，咋开饭啊？”顾晓楠没好气地瞥了一眼对面的冤家，“我妈哪次不等你一起吃了，我要是吃独食岂不是要被批斗。”“也是，学术搞不好，吃饭还要抢第一的话真的要去劳改了。”简琛笑弯了眼，一对卧蚕漾在长睫毛下，少见的有了生气。“哼！你不许跟我妈打小报告！”“好啊，可你得答应我一个条件。”“什么条件？”“和我一起自习。”“什么？！”顾晓楠惊得差点从楼梯上摔下去，说实话自打十二岁那年起除了吃饭上学她可不再和简琛有什么交集，况且做人要有自知之明，简琛带着拖油瓶自习？这显然是报复！这一定有阴谋！顾晓楠自然是个硬气的姑娘，宁做车尾，不当陪读，她可是有原则的，“我不干！”“你这次月考年级排名多少来着？阿姨是不是上次也不知道要去开家长会啊？”简琛说得慢条斯理。“不是吧？你要打小报告？！这么阴…”，“险”字还没说出口，顾晓楠已经看到面前的人慢慢拿出自己的成绩单开始装模作样。“哎，正巧我妈不在呢，要不然索性让阿姨一起帮我去了得了。”简琛说着把自己的成绩单晃了晃，阳光透过楼道的矮窗落在白纸黑字上，顾晓楠看得分明。“我……我去……”顾晓楠终于在一生中屈指可数地低了头。她愤恨地鼓着腮帮，恨不能回到七岁的时候，那时候简琛要是敢在她面前这么耍无赖她一定会揪着他的衣领跟他开打。于是乎六年里，简琛第一次名正言顺让顾晓楠乖乖听话，这确实是一件令人开心的事情，仿佛长久以来心中的石头终于落了地，简琛第一次觉得四月真是春暖花开令人开心的好时光。但是这样轻松欢快的日子是转瞬即逝的，当顾晓楠几次三番当着简琛的面在书桌前睡得哈喇子流一桌之后，简琛不得不说有些许后悔，让顾晓楠来一起学习恐怕是个奢望。弹了一指对面又快睡着的姑娘的脑门，简琛幽怨地把她望着。“你干嘛？”顾晓楠磕着的半边面颊还留着印子，摸了摸脑门，显然是非常不满。“你的数学卷子都要被你的口水浸湿了。”简琛拿笔指了指，一脸嫌弃。“数学就是让人犯困，我把会做的都做了，别的等明天老师讲呗。”“你这一半都是空的吧。”简琛叹气，撵过沾了口水的数学卷，“你坐过来。”“不会吧？你帮我做？”顾晓楠顿时笑逐颜开。“我讲一遍，记不住，月考成绩，你懂的。”少年不怀好意地微笑，顾晓楠一脸绝望。这大概是打小第一次吧，顾晓楠离简琛这么近，近得仿佛能感受到对方的呼吸，顾晓楠仔细地瞧着一边苦心孤诣地讲解着的简琛，突然觉得面前的人似乎不一样了，不在是以前那个柔柔弱弱的小孩子，仿佛这一刻开始以后的每一天，他都会是那样一个干净爽朗的少年。简琛自然知道面前的傻丫头在盯着自己，他从来不会主动拆穿别人，仍由着顾晓楠出神。他尽量装作若无其事地讲着数学题，说实话，这些题目简琛早就在各种参考书上做过百十来次，各种公式和解题技巧他理应烂熟于心，可这一次讲给顾晓楠这个笨蛋听，他竟然有些紧张，他不太自然地呼吸着，故作镇定地搬着一套套公式，终于一不小心对上了顾晓楠空洞的眼神。聪明如简琛，他明知顾晓楠根本没在听他的话，他也清楚顾晓楠发什么呆，但这一次他却较了真。“顾晓楠，你到底在想什么？”“啊？我，我没有啊，我在听你讲题啊。”“那好，你把我刚才讲的重新写一遍。”“我……我看一下啊。”“说实话，你都不知道我讲到哪里了吧。”“……”“你到底在想什么，顾晓楠，为什么你总是和别人不一样？”简琛生了气，这似乎是他迄今为止的人生中屈指可数地生气，他向来安静沉稳，纵使遇人不淑也不会无端失了风度，可这一次他却着实对面前的小姑娘发了火。 “你干嘛这么凶啊，我，我又不是故意的。”顾晓楠不懂简琛为什么生气，她突然感到很委屈，似乎从七岁起受的冷落一下子累积到达了一个顶点，她竟然哭起来，还是当着简琛的面，抽抽嗒嗒，很没有骨气，“简琛，你凭什么对我发火？我，我……”顾晓楠自顾自语无伦次着。简琛人生中第一次看到顾晓楠哭，仿佛是太阳打西边出来了，此时此刻他完全慌了神，说实话，像顾晓楠这么倔强的人怎么可能在自己的面前哭呢。简琛向来认为顾晓楠是个简单的人，但面前的场景着实让他措手不及，难道是她被什么妖魔鬼怪附体了？简琛若有所思，伸手对着面前抹眼泪的小姑娘的额头又弹了一指，这下顾晓楠不哭了，两个人僵持了大概半分钟，空气中的尴尬似乎都快溢出来。“欸，顾晓楠。”“干嘛？”“你怎么不哭了？”“……”“要不然，我给你讲个笑话？”“你当我三岁，还哄我呢！”顾晓楠又羞愧又生气，被弹了一脑瓜之后真想立刻把面前一脸无辜的人打一顿。“那……要不然……你想听什么？”简琛不知道该怎么哄一个哭鼻子的人，他甚至不明白自己为什么要自讨苦吃。大概这就是孽缘吧，碰上了顾晓楠，他总是没辙的。“要不然，你说点八卦我听听。”“八卦？你想听谁的？”“你还真知道？我以为你两耳不闻窗外事呢。”顾晓楠略有吃惊。“别人的我可不知道，但我知道你的。”“什么？我的？你说来听听，我倒想知道编排我些什么呢！”顾晓楠更加吃惊了，她睁圆了眼睛盯着简琛，倒映在她眼珠子里的少年笑盈盈的，倒像是一副看好戏的模样。“你不是喜欢我吗？”简琛说得漫不经心，甚至带了一丝嘲讽。四月的春风拂过万花，吹迷了人眼，吹乱了人心，困倦的午后那些本不以为意的闲言碎语，此时此刻从简琛嘴里面吐出来，却像是一个巴掌，让顾晓楠瞠目结舌。顾晓南甚至来不及思考，她只是告诉自己不要发怒，不要发努，但终究她还是十分及其，“你什么意思？”愤怒反而使她冷静下来，她沉了嗓音。“我的意思是……别人说你喜欢我，但是我知道这是假的。”简琛也突然一本正经起来。顾晓楠着实看不懂面前的人到底在想什么，“你不说，我就走了。”，顾晓楠伸手去整理自己的习题和文具，刚够到放在简琛面前的数学卷，手却被一把按住。顾晓楠再一次睁圆了眼睛，她看着简琛，不知道他想做什么。“我知道这是假的，可是我喜欢你啊。”简琛笑着松了手。这或许是一种冲击，顾晓楠从没想过事情会这么发展，以至于若干年后她回想起来当时的场景还总是要质问简琛到底哪里学来的心眼，自然这都是后话。这一次顾晓楠足足呆滞了十分钟，她的大脑一切空白，然后她做了一个决定，一个很简单的决定。“既然这样，那我也喜欢你好了。”“那我也喜欢你好了”，这句话在简琛的大脑中反复地循环起来，少年笑着，似乎是春光确实明媚，也似乎是因为眼前明媚的少女，还有……大概是因为生涩的种子终于破土生芽，慢慢长大，开了一朵明媚的花……",
		Cookie:               "zhangwei",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	res := GetKeyWord(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}

}

func Test_ListImageByUserCookie(t *testing.T) {
	ctx := initenv()
	req := &corpus.ListImageByUserCookieReq{
		Cookie:               "zhangwei",
		Limit:                10,
		Offset:               0,
	}
	res := ListImageByUserCookie(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}

}

func TestSendMessage2(t *testing.T) {
	ctx := initenv()
	req := &corpus.SendMessageReq{
		Phone:                "18846082154",
	}
	res := adapter.SendMessage(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}
}

func TestSendEmail(t *testing.T) {
	ctx := initenv()
	SendEmail(ctx,[]string{"858777157@qq.com"})
}

func TestUpdateUserPhone(t *testing.T) {
	dll := syscall.MustLoadDLL("../dct.dll")
	log.Printf("%v",dll)
}