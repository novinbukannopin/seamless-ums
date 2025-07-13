package cmd

import (
	"log"
	"net"
	tokenvalidation "seamless-ums/cmd/proto/token_validation"
	"seamless-ums/helpers"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ServeGRPC() {

	dependency := DI()

	lis, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", "7000"))
	if err != nil {
		log.Fatal("failed to listen grpc port: ", err)
	}

	s := grpc.NewServer()

	tokenvalidation.RegisterTokenValidationServer(s, dependency.TokenValidationAPI)

	logrus.Info("start listening grpc on port:" + helpers.GetEnv("GRPC_PORT", "7000"))
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve grpc port: ", err)
	}
}
