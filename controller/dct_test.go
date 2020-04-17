package controller

import (
	"github.com/astaxie/beego/logs"
	"gocv.io/x/gocv"
	"log"
	"testing"
	"time"
)

func  TestSaveImage(t *testing.T) {
	file , err := SaveImage("https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=4251611548,332805913&fm=11&gp=0.jpg")
	//err := SaveImage("http://hbimg.b0.upaiyun.com/32f065b3afb3fb36b75a5cbc90051b1050e1e6b6e199-Ml6q9F_fw320")
	if err != nil{
		logs.Error(err)
	}
	log.Printf("%v",file)
}

func TestGetMd5String(t *testing.T) {
	err := myMat("C:\\Users\\image\\u=4251611548,332805913&fm=11&gp=0.jpg")
	if err != nil{
		logs.Error(err)
	}
}

func TestAddTransAudio(t *testing.T) {
	MyDCTRGB("C:\\image\\32f065b3afb3fb36b75a5cbc90051b1050e1e6b6e199-Ml6q9F_fw320.jpg")
	//MyiDCTRGB("C:\\image\\yuv\\","eafddcb6c451bdfa64cfe8148c340469")
	//MyIDCT("C:\\image\\xml\\","5cc45fcecbab03869afba0551242d3d0")
	//dat, err := ioutil.ReadFile("C:\\image\\32f065b3afb3fb36b75a5cbc90051b1050e1e6b6e199-Ml6q9F_fw320.jpg")
	//if err != nil {
	//	logs.Error(err)
	//}
	//picture_trans(dat)
}

func TestListAuth(t *testing.T) {
	img := gocv.IMRead("https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=4251611548,332805913&fm=11&gp=0.jpg",0664)
	defer img.Close()
	rows := img.Rows()
	cols:= img.Cols()

	window := gocv.NewWindow("origin")
	defer window.Close()

	//flag := gocv.IMWrite("yuantu",img)
	//if !flag{
	//	fmt.Println("原图失败")
	//}
	//Mat grayimage; CV_BGR2GRAY=6
	//cvtColor(img, grayimage, CV_BGR2GRAY);
	////浮点变换
	//Mat fgrayimage(img.rows, img.cols, CV_64FC1);
	//grayimage.convertTo(fgrayimage, CV_64FC1);
	//
	////c++ dct函数实现
	//Mat B(rows, cols, CV_64FC1);
	//mydct(fgrayimage, B, rows, cols);
	//imshow("dct后图像", B);
	//waitKey(0);
	//c++ dct函数实现
	//Mat B(rows, cols, CV_64FC1);
	//mydct(fgrayimage, B, rows, cols);
	//imshow("dct后图像", B);
	//waitKey(0);
	//
	////idct实现
	////Mat C(rows, cols, CV_64FC1);
	//Mat C(rows, cols, CV_8UC1);;
	//myidct(B, C, rows, cols);
	//imshow("idct后图像", C);
	//waitKey(0);
	grayimage := gocv.Mat{}
	gocv.CvtColor(img,&grayimage,gocv.ColorBGRToGray)

	fgrayimage := gocv.NewMatWithSize(rows,cols,gocv.MatTypeCV64FC1)
	grayimage.ConvertTo(&fgrayimage,gocv.MatTypeCV64FC1)

	dctRes := gocv.NewMatWithSize(rows,cols,gocv.MatTypeCV64FC1)
	window.IMShow(dctRes)
	//gocv.DCT(fgrayimage,&dctRes,gocv.DftComplexOutput)
	//flag = gocv.IMWrite("dct",img)
	//if !flag{
	//	fmt.Println("dct")
	//}
	gocv.WaitKey(0)
	time.Sleep(time.Second*20)
}
