package main

import (
	"log"
	"net"

	pb "github.com/KolManis/go-grpc/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

//go run $(Get-ChildItem -Filter *.go -Name)
//go run *.go

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	// ↑ "Когда придет запрос к GreetService,
	//   вызывай методы helloServer"
	log.Printf("server started at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Falide to start: %v", err)
	}
}
