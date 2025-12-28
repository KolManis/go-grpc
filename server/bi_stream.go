package main

import (
	"io"
	"log"

	pb "github.com/KolManis/go-grpc/proto"
)

func (s *helloServer) SayHelloBidirectionStreaming(stream pb.GreetService_SayHelloBidirectionStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("Error: %v", err)
			return err
		}
		log.Printf("Got request with name: %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
