package main

import (
	"log"
	"net"

	"github.com/lucasgfsouza/full-cycle-project1/pb"
	"github.com/lucasgfsouza/full-cycle-project1/services"
	"google.golang.org/grpc"
)


func main() {
	fmt.Println('teste')
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &services.NewUserService)
	// reflection.Register(grpcServer)
	reflection.Register(grpcServer)
	


	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}

}
