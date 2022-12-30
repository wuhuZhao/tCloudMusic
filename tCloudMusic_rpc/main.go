package main

import (
	"net"
	api "tCloudMusic_rpc/kitex_gen/tCloudMisc/biz/api/bizservice"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 手动指定地址
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8889")
	dns := "root:123456@tcp(127.0.0.1:3306)/USER_INFO?charset=utf8mb4&parseTime=True&loc=Local"
	mysql, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		klog.Errorf("mysql connection error, err: %v\n", err.Error())
	}
	svr := api.NewServer(NewBizServiceImpl(mysql), server.WithServiceAddr(addr))
	if err := svr.Run(); err != nil {
		klog.Errorf("kitex server start error, err: %v\n", err.Error())
	}
}
