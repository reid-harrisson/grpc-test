package main

import (
	"log"
	"net"

	pb "grpc-practise/proto"
	"grpc-practise/services"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "[::1]:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("tcp started")

	grpcServer := grpc.NewServer()
	service := &services.UserServiceServer{}

	log.Println("service created")

	pb.RegisterUserServiceServer(grpcServer, service)

	log.Println("register user service")

	err = grpcServer.Serve(lis)

	log.Println(err)

	if err != nil {
		log.Fatalf("Error strating server: %v", err)
	}

	log.Println("grpc service created")
}
