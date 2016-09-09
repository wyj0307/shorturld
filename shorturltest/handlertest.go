package main

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/etcd"
	"golang.org/x/net/context"

	"crazyant.com/groot/pbd/go/shorturld"
	. "crazyant.com/groot/pkg/consts"

	"shorturld/handler"
)

var shorturl = &handler.Shorturl{}

func InitServer() {
	service := micro.NewService(
		micro.Name(SERVICE_SHORTURLD),
		micro.RegisterTTL(time.Second*8),
		micro.RegisterInterval(time.Second*3),
	)

	service.Init()
}

//var shorURL string
var shortUrlArray []string
var originalArr []string

func main() {

	InitServer()

	// 网址测试数组，网址内容随意取，只在乎格式
	originalArr = []string{
		"http://blog.sina.com.cn/s/blog_8882a6260101a4ox.html",
		"http://tech.sina.com.cn/d/s/2016-09-08/doc-ifxvukhv7802989.shtml",
		"http://d1.sina.com.cn/201608/01/1427196_caijing.jpg",
		"http://image.baidu.com/search/detail?ct=503316480&z=0&ipn=d&word=%E5%B0%8F%E6%B8%85%E6%96%B0&step_word=&hs=0&pn=68&spn=0&di=178328213470&pi=&rn=1&tn=baiduimagedetail&is=&istype=2&ie=utf-8&oe=utf-8&in=&cl=2&lm=-1&st=-1&cs=2665846294%2C2104883459&os=3381810428%2C2492197421&simid=3456304449%2C71694063&adpicid=0&ln=1991&fr=&fmq=1459502282690_R&fm=&ic=0&s=undefined&se=&sme=&tab=0&width=&height=&face=undefined&ist=&jit=&cg=&bdtype=0&oriquery=&objurl=http%3A%2F%2Fwww.bz55.com%2Fuploads%2Fallimg%2F150428%2F140-15042Q51324-50.jpg&fromurl=ippr_z2C%24qAzdH3FAzdH3Fooo_z%26e3Byx3xbb_z%26e3Bv54AzdH3Fv-dmdd9_z%26e3Bip4s&gsm=10000001e&rpstart=0&rpnum=0",
		"http://www.crazyant.com/?124skdfhjfsahfkfjslkfjlakfjds89ijsjdflksjfliewfjslkfjsaifljaifnmewf;rkjwfkm;jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjlllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllfffff",

		"https://app.bitly.com/default/bitlinks/2c7qr32",
	}

	for _, v := range originalArr {
		// 短网址转换
		shortUrlTest(v)
	}

	logrus.Infoln("测试完毕！")
}

func shortUrlTest(url string) {
	longURL := shorturld.UrlReq{
		LongUrl:  url,
		Provider: "gggg",
	}

	resp := shorturld.UrlResp{}

	err := shorturl.ConvertUrlShort(context.Background(), &longURL, &resp)
	if err != nil {
		logrus.Infof("短网址转换失败，err:%v", err)
		return
	}

	shorURL := resp.DesUrl
	logrus.Infof("短网址转换成功，url:%s", shorURL)

	expandUrlTest(shorURL)
}

func expandUrlTest(url string) {
	sUrl := shorturld.ShorturlReq{
		Shorturl: url,
	}

	resp := shorturld.UrlResp{}
	err := shorturl.ConvertUrlOriginal(context.Background(), &sUrl, &resp)
	if err == nil {
		logrus.Infof("还原之后的url:%s", resp.DesUrl)
		return
	}
	logrus.Infof("url还原失败，err:%v", err)
}
