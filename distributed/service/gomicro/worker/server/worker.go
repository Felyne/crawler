package main

import (
	"github.com/Felyne/crawler/distributed/service/gomicro/worker"
	"github.com/Felyne/crawler/distributed/service/gomicro/worker/pb"
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

func setup(s server.Server, _ string) error {
	return pb.RegisterCrawlerHandler(s, worker.CrawlerService{})
}
