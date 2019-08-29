package main

import (
	"cloudali/plugin/logs"
	"flag"
	"fmt"
	"protocol"
	"protocol/config"
	"protocol/version"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-web"
	"github.com/robfig/cron"
)

func main() {
	port := flag.String("port", "8191", "port")
	flag.Parse()
	// Create service
	if !config.ConsulAlive() {
		// logs.Error("Consul is not alive.(consul服务链接失败,请检查consul的配置" + config.ConsulAddress() + ")")
		return
	}
	service := web.NewService(
		web.Name(protocol.WebAPIService),
		web.Address("0.0.0.0:"+*port),
		web.Registry(consul.NewRegistry(registry.Addrs(config.CConsulAddr()))),
		web.RegisterTTL(config.CRegisterTtL()),
		web.RegisterInterval(config.CRegisterInterval()),
		web.Flags(*version.BuildVersionFlag()),
	)
	service.Init()
	//Register Restful Handler
	service.Handle("/", GetRouterContainer())
	// Crontab
	InitCrontab()
}

func InitCrontab() {
	logs.Info("begin init crontab")
	crontable := cron.New()
	// golang 6位 精确秒
	spec := `0 0 0 * * ?`
	crontable.AddFunc(spec, func() {
		logs.Info("crontab sync and check running...")
	})

	crontable.AddFunc(spec, funcName)
	crontable.Start()
}

func funcName() {
	fmt.Println("test...1")
}
