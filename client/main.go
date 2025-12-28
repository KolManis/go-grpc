package main

import (
	"log"

	pb "github.com/KolManis/go-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	// С TLS/SSL шифрованием
	//creds, _ := credentials.NewClientTLSFromFile("cert.pem", "")
	//conn, err := grpc.Dial("example.com:443", grpc.WithTransportCredentials(creds))
	// dial устарел
	// без шифрования
	//conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(
		"localhost"+port,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)
	names := &pb.NameList{
		Names: []string{"Kolmanis", "Danya", "Kolya"},
	}

	//callSayHello(client)
	//callSayHelloServerStream(client, names)
	//callSayHelloClientStream(client, names)
	callHelloBidirectionalStream(client, names)
}
