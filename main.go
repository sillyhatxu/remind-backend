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
	"github.com/sillyhatxu/remind-backend/dao"
	"github.com/sillyhatxu/remind-backend/grpc/grpcserver"
	"github.com/sillyhatxu/remind-backend/service/remindbirthday"
	"github.com/sirupsen/logrus"
	"net"
	"os"
)

func init() {
	cfgFile := flag.String("c", "config-local.conf", "config file")
	flag.Parse()
	if os.Getenv("env") == "" {
		err := envconfig.LoadEnv(".env")
		if err != nil {
			panic(err)
		}
	}
	err := envconfig.ParseDefaultConfig(envconfig.ConfigFile(*cfgFile))
	if err != nil {
		panic(err)
	}
	fields := logrus.Fields{
		"project":  envconfig.Conf.Project,
		"module":   envconfig.Conf.Module,
		"@version": "1",
		"type":     "project-log",
	}
	logrusconf.NewLogrusConfig(
		logrusconf.Fields(fields),
		logrusconf.FileConfig(filehook.NewFileConfig(envconfig.Conf.Log.FilePath, filehook.Open(envconfig.Conf.Log.OpenLogfile))),
		logrusconf.LogstashConfig(logstashhook.NewLogstashConfig(envconfig.Conf.LogstashURL, logstashhook.Open(envconfig.Conf.Log.OpenLogstash), logstashhook.Fields(fields))),
	).Initial()
	logrus.Infof("config : %#v", envconfig.Conf)
}

func main() {
	//for _, pair := range os.Environ() {
	//	fmt.Println(pair)
	//}
	err := dao.InitialMysqlClient()
	if err != nil {
		panic(err)
	}
	err = remindbirthday.InitialBirthdayClient()
	if err != nil {
		panic(err)
	}
	consulServer := consul.NewConsulServer(
		envconfig.Conf.ConsulAddress,
		envconfig.Conf.Module,
		envconfig.Conf.Host,
		envconfig.Conf.GRPCPort,
		consul.CheckType(consul.HealthCheckGRPC),
	)
	err = consulServer.Register()
	if err != nil {
		panic(err)
	}
	apiListener, err := net.Listen("tcp", fmt.Sprintf(":%d", envconfig.Conf.HttpPort))
	if err != nil {
		panic(err)
	}
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", envconfig.Conf.GRPCPort))
	if err != nil {
		panic(err)
	}
	go api.InitialAPI(apiListener)
	go grpcserver.InitialGRPC(grpcListener)
	forever := make(chan bool)
	<-forever
}
