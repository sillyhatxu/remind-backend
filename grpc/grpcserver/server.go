package grpcserver

import (
	"github.com/sillyhatxu/consul-client"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	hv1 "google.golang.org/grpc/health/grpc_health_v1"
	"net"
)

func InitialGRPC(listener net.Listener) {
	logrus.Info("initial grpc server")
	server := grpc.NewServer()

	healthServer := health.NewServer()
	healthServer.SetServingStatus(consul.DefaultHealthCheckGRPCServerName, hv1.HealthCheckResponse_SERVING)
	hv1.RegisterHealthServer(server, healthServer)

	err := server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
