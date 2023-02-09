package grpc_server

import (
	"fmt"
	"log"
	"net"

	"github.com/Nav1Cr0ss/s-lib/configuration"
	"github.com/Nav1Cr0ss/s-lib/interceptor"
	"github.com/Nav1Cr0ss/s-lib/logger"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	cfg configuration.Configuration
	log *logger.Logger
	Reg grpc.ServiceRegistrar
	lis net.Listener
}

func NewGRPCServer(
	cfg configuration.Configuration,
	log *logger.Logger,
) *GRPCServer {

	srv := GRPCServer{
		cfg: cfg,
		log: log,
	}
	srv.Reg = srv.initGRPC(log)
	srv.lis = srv.initListener()
	return &srv
}

func (s *GRPCServer) initGRPC(log *logger.Logger) *grpc.Server {
	options := grpc.UnaryInterceptor(interceptor.ChainUnaryServer(
		interceptor.LoggerInterceptor(log),
		interceptor.UserInterceptor(),
		interceptor.ValidatorInterceptor(),
	))
	return grpc.NewServer(options)
}

func (s *GRPCServer) initListener() net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.cfg.GetHost(), s.cfg.GetPort()))
	if err != nil {
		_ = lis.Close()
		s.log.Fatalf("error on starting listening : %s", err)
	}
	return lis
}

func Serve(s *GRPCServer) {
	server, ok := s.Reg.(*grpc.Server)
	if !ok {
		log.Fatal("error on parsing navix type")
	}

	s.log.Infof("Listening tcp: %s", s.lis.Addr())
	s.log.Infof("Debug: %t", s.cfg.GetDebug())

	err := server.Serve(s.lis)
	if err != nil {
		_ = s.lis.Close()
		server.GracefulStop()
		log.Fatal("error on starting serving")
	}

}
