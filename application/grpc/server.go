package grpc

import (
	"codePix/application/grpc/pb"
	"codePix/application/useCase"
	"codePix/repository"
	"fmt"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func StartGrpcServer(db *gorm.DB, serverPort int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pixKeyRepository := repository.PixKeyRepository{Db: db}
	pixUseCase := useCase.PixUseCase{PixKeyRepository: &pixKeyRepository}
	pixKeyGrpcServiceController := NewPixGrpcServiceController(pixUseCase)
	pb.RegisterPixServiceControllerServer(grpcServer, pixKeyGrpcServiceController)

	address := fmt.Sprintf("0.0.0.0:%d", serverPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start gRPC Server", err)
	}

	log.Printf("gRPC Server has been started on port %d", serverPort)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC Server", err)
	}
}
