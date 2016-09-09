package main

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/etcd"

	"shorturld/handler"

	"crazyant.com/groot/pbd/go/shorturld"
	. "crazyant.com/groot/pkg/consts"
)

func main() {
	service := micro.NewService(
		micro.Name(SERVICE_SHORTURLD),
		micro.RegisterTTL(time.Second*8),
		micro.RegisterInterval(time.Second*3),
	)

	service.Init()

	shorturld.RegisterShorturlServiceHandler(service.Server(), new(handler.Shorturl))

	if err := service.Run(); err != nil {
		logrus.Fatalf("启动shorturld时出错: %v", err)
	}
}
