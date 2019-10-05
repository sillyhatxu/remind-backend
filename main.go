package main

import (
	"flag"
	"fmt"
	"github.com/sillyhatxu/consul-client"
	"github.com/sillyhatxu/environment-config"
	"github.com/sillyhatxu/logrus-client"
	"github.com/sillyhatxu/logrus-client/filehook"
	"github.com/sillyhatxu/logrus-client/logstashhook"
	"github.com/sillyhatxu/remind-backend/api"
	"github.com/sillyhatxu/remind-backend/config"
	"github.com/sillyhatxu/remind-backend/grpc/grpcserver"
	"github.com/sirupsen/logrus"
	"net"
	"os"
)

func init() {
	cfgFile := flag.String("c", "config-local.conf", "config file")
	flag.Parse()
	err := envconfig.ParseConfig(&config.Conf, envconfig.ConfigFile(*cfgFile), envconfig.Environment(os.Getenv("env")))
	if err != nil {
		panic(err)
	}
	fields := logrus.Fields{
		"project":  config.Conf.Project,
		"module":   config.Conf.Module,
		"@version": "1",
		"type":     "project-log",
	}
	logrusconf.NewLogrusConfig(
		logrusconf.Fields(fields),
		logrusconf.FileConfig(filehook.NewFileConfig(config.Conf.Log.FilePath, filehook.Open(config.Conf.Log.OpenLogfile))),
		logrusconf.LogstashConfig(logstashhook.NewLogstashConfig(config.Conf.LogstashURL, logstashhook.Open(config.Conf.Log.OpenLogstash), logstashhook.Fields(fields))),
	).Initial()
}

func main() {
	//for _, pair := range os.Environ() {
	//	fmt.Println(pair)
	//}
	consulServer := consul.NewConsulServer(
		config.Conf.HostConsul,
		config.Conf.Module,
		config.Conf.Host,
		config.Conf.GRPCPort,
		consul.CheckType(consul.HealthCheckGRPC),
	)
	err := consulServer.Register()
	if err != nil {
		panic(err)
	}
	apiListener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Conf.HttpPort))
	if err != nil {
		panic(err)
	}
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Conf.GRPCPort))
	if err != nil {
		panic(err)
	}
	go api.InitialAPI(apiListener)
	go grpcserver.InitialGRPC(grpcListener)
	forever := make(chan bool)
	<-forever
}
