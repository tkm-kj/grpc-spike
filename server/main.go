package main

import (
	"log"
	"net"

	pb "github.com/tkm-kj/grpc-spike/proto"
	"github.com/tkm-kj/grpc-spike/services"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterCommentServiceServer(s, new(services.CommentService))
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}