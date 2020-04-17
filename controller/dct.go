package controller

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego/logs"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

var Pi = 3.141596

func Md5(filePath string) (string,error){
	f, err := os.Open(filePath) //打开文件
	if nil != err {
		fmt.Println(err)
		return "", err
	}
	defer f.Close()

	md5Handle := md5.New()      //创建 md5 句柄
	_, err = io.Copy(md5Handle, f)  //将文件内容拷贝到 md5 句柄中
	if nil != err {
		fmt.Println(err)
		return "", err
	}
	md := md5Handle.Sum(nil)    //计算 MD5 值，返回 []byte

	md5str := fmt.Sprintf("%x", md) //将 []byte 转为 string
	return md5str, nil
}

func MD5Audio(filePath string) (string,error){
	tmp := filePath[0:4]
	var pis []byte
	if tmp == "http"{
		resp, err := http.Get(filePath)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		pix, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		pis = pix
	} else {
		dat, err := ioutil.ReadFile(filePath)
		if err != nil {
			logs.Error(err)
			return "",err
		}
		pis = dat
	}
	md5 := md5.New()
	md5.Write(pis)
	MD5Str := hex.EncodeToString(md5.Sum(nil))
	return MD5Str,nil
}

func MD5Text(filePath string) string{
	md5 := md5.New()
	md5.Write([]byte(filePath))
	MD5Str := hex.EncodeToString(md5.Sum(nil))
	return MD5Str
}

func StrPtr(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringBytePtr(s)))
}

func MyDCT(filename string) error{
	time.Sleep(time.Second*5)
	dctDll := syscall.MustLoadDLL("../dct.dll")
	fmt.Println("+++++++syscall.LoadLibrary:", dctDll, "+++++++")
	mydct := dctDll.MustFindProc("mydct")
	fmt.Println("GetProcAddress", mydct)
	fmt.Printf("%v",StrPtr(filename))
	md5Str, err := Md5(filename)
	if err != nil{
		logs.Error(err)
		return err
	}
	//output := fmt.Sprintf("C:\\Users\\华硕\\Desktop\\pr\\age\\%v.wav",number)
	outPath := fmt.Sprintf("C:\\image\\xml\\%v.xml", md5Str)
	ret, _, err := mydct.Call(StrPtr(filename),StrPtr(outPath))
	if err != nil {
		fmt.Println("DllTestDef.dll运算结果为:", ret)
	}
	return nil
}

func MyIDCT(filename string,md5 string) error{
	dctDll := syscall.MustLoadDLL("../dct.dll")
	fmt.Println("+++++++syscall.LoadLibrary:", dctDll, "+++++++")
	myidct := dctDll.MustFindProc("myidct")
	fmt.Println("GetProcAddress", myidct)

	ret, _, err := myidct.Call(StrPtr(filename),StrPtr(md5))
	if err != nil {
		fmt.Println("DllTestDef.dll运算结果为:", ret)
	}
	return nil
}

func MyDCTRGB(filename string) error{
	time.Sleep(time.Second*5)
	dctDll := syscall.MustLoadDLL("../dct.dll")
	fmt.Println("+++++++syscall.LoadLibrary:", dctDll, "+++++++")
	mydctrgb := dctDll.MustFindProc("mydctRGB")
	fmt.Println("GetProcAddress", mydctrgb)
	fmt.Printf("%v",StrPtr(filename))
	md5Str, err := Md5(filename)
	if err != nil{
		logs.Error(err)
		return err
	}
	//output := fmt.Sprintf("C:\\Users\\华硕\\Desktop\\pr\\age\\%v.wav",number)
	outPath := fmt.Sprintf("C:\\image\\yuv\\%v.jpg", md5Str)
	ret, _, err := mydctrgb.Call(StrPtr(filename),StrPtr(outPath))
	if err != nil {
		fmt.Println("DllTestDef.dll运算结果为:", ret)
	}
	return nil
}

func MyiDCTRGB(filename string,md5 string)error{
	dctDll := syscall.MustLoadDLL("../dct.dll")
	fmt.Println("+++++++syscall.LoadLibrary:", dctDll, "+++++++")
	myidctrgb := dctDll.MustFindProc("myidctRGB")
	fmt.Println("GetProcAddress", myidctrgb)

	ret, _, err := myidctrgb.Call(StrPtr(filename),StrPtr(md5))
	if err != nil {
		fmt.Println("DllTestDef.dll运算结果为:", ret)
	}
	return nil
}

func SaveImage(imgUrl string) (string,error){
	var file *os.File
	var err error
	var filePath string
	imgPath := "C:\\image\\"
	fileName := path.Base(imgUrl)
	res, err := http.Get(imgUrl)
	if err != nil{
		logs.Error(err)
		return "",err
	}
	defer res.Body.Close()
	if strings.Contains(imgUrl,".jpg") || strings.Contains(imgUrl,".jpeg"){
		fmt.Printf("%v photo is jpg or jpeg \n",imgUrl)
		filePath = imgPath+ fileName
		file, err = os.Create(filePath)
		if err != nil{
			logs.Error(err)
			return "",err
		}
	} else if strings.Contains(imgUrl,".png"){
		fmt.Printf("%v photo png \n",imgUrl)
		filePath = imgPath+ fileName
		file, err = os.Create(filePath)
		if err != nil{
			logs.Error(err)
			return "",err
		}
	} else if strings.Contains(imgUrl,".gif"){
		fmt.Printf("%v photo gif \n",imgUrl)
		filePath = imgPath+ fileName
		file, err = os.Create(filePath)
		if err != nil{
			logs.Error(err)
			return "",err
		}
	} else {
		fmt.Printf("%v default photo is jpg or jpeg \n",imgUrl)
		filePath = imgPath+ fileName +".jpg"
		file, err = os.Create(filePath)
		if err != nil{
			logs.Error(err)
			return "",err
		}
	}
	reader := bufio.NewReaderSize(res.Body,32*1024)

	writer := bufio.NewWriter(file)
	written, err := io.Copy(writer,reader)
	if err != nil{
		logs.Error(err)
		return "",err
	}

	fmt.Printf("total length %d", written)
	return filePath,nil
}

//流程 灰度化 二值化 缩小图片 膨胀 高斯模糊 查找轮廓 绘制轮廓 绘制轮廓的最小外接矩阵
func myMat(filePath string) error{
	window := gocv.NewWindow("Hello")
	img := gocv.IMRead(filePath, gocv.IMReadColor)
	grayImage := gocv.NewMat()
	defer grayImage.Close()

	gocv.CvtColor(img, &grayImage, gocv.ColorBGRToGray)
	destImage := gocv.NewMat()
	gocv.Threshold(grayImage, &destImage, 100, 255, gocv.ThresholdBinaryInv)
	resultImage := gocv.NewMatWithSize(500, 400, gocv.MatTypeCV8U)

	gocv.Resize(destImage, &resultImage, image.Pt(resultImage.Rows(), resultImage.Cols()), 0, 0, gocv.InterpolationCubic)
	gocv.Dilate(resultImage, &resultImage, gocv.NewMat())
	gocv.GaussianBlur(resultImage, &resultImage, image.Pt(5, 5), 0, 0, gocv.BorderWrap)
	results := gocv.FindContours(resultImage, gocv.RetrievalTree, gocv.ChainApproxSimple)
	imageForShowing := gocv.NewMatWithSize(resultImage.Rows(), resultImage.Cols(), gocv.MatChannels4)
	for index, element := range results {
		fmt.Println(index)
		gocv.DrawContours(&imageForShowing, results, index, color.RGBA{R: 0, G: 0, B: 255, A: 255}, 1)
		gocv.Rectangle(&imageForShowing,
			gocv.BoundingRect(element),
			color.RGBA{R: 0, G: 255, B: 0, A: 100}, 1)
	}

	if img.Empty() {
		fmt.Printf("Error reading image from: %v\n", filePath)
		return nil
	}

	for {
		window.IMShow(imageForShowing)
		if window.WaitKey(1) >= 0 {
			break
		}
	}

	return nil
}


