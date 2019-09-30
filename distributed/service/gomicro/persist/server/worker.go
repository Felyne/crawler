package main

import (
	"log"

	"github.com/astaxie/beego/config"

	"github.com/Felyne/crawler/distributed/service/gomicro/persist/pb"

	"github.com/Felyne/crawler/distributed/service/gomicro/persist"

	"github.com/Felyne/service_launch"
	"github.com/micro/go-micro/server"
)

var (
	Version   = ""
	BuildTime = ""
)

func main() {
	serviceName := pb.ServiceName_name[0]
	service_launch.Start(serviceName, Version, BuildTime, setup)
}

func setup(s server.Server, cfgContent string) error {
	cfg, err := config.NewConfigData("ini", []byte(cfgContent))
	if err != nil {
		log.Printf("NewConfigData() failed: %v", err)
		return err
	}
	h, err := persist.NewItemSaverService(cfg)
	if err != nil {
		return err
	}
	return pb.RegisterItemSaverHandler(s, h)
}
