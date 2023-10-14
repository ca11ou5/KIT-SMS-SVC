package main

import (
	"SMS_Service/configs"
	"SMS_Service/internal/pb"
	"SMS_Service/internal/service"
	"github.com/go-resty/resty"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cfg := configs.InitConfig()

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatal("failed to listen")
	}

	grpcServer := grpc.NewServer()
	s := service.Server{
		ApiToken:   cfg.ApiToken,
		RestClient: resty.New(),
	}

	pb.RegisterSmsServiceServer(grpcServer, &s)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("failed to serve: " + err.Error())
	}
}
