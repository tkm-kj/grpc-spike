package main

import (
	"fmt"
	"log"
	"context"

	pb "github.com/tkm-kj/grpc-spike/proto"
	"google.golang.org/grpc"
)

const (
	address = ":8080"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewCommentServiceClient(conn)
	ctx := context.Background()

	cr := pb.CommentRequest{Id: 1}
	com, err := c.GetComment(ctx, &cr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(com.Comment)

	csr := pb.CommentsRequest{PlaylistId: 1}
	coms, err := c.GetCommentsByPlaylistId(ctx, &csr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(coms.GetComment()[1].Comment)
}