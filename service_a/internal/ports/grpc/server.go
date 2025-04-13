package grpc

import (
	service "addition_service/internal/domain/addition"
	"addition_service/internal/outbox"

	pb "addition_service/internal/api/addition"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGRPCServer(outboxRepo *outbox.SQLRepository) *grpc.Server {
	grpcServer := grpc.NewServer()

	pb.RegisterAdditionServiceServer(grpcServer, service.NewAdditionService(outboxRepo))

	reflection.Register(grpcServer)

	return grpcServer
}
