package main

import (
	"fmt"
	"log"
	"context"
	"time"

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

	ctx := context.Background()

	/* GetComment
	c := pb.NewCommentServiceClient(conn)
	cr := pb.CommentRequest{Id: 1}
	com, err := c.GetComment(ctx, &cr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(com.Comment)
	*/

	/* GetCommentsByPlaylistId
	csr := pb.CommentsRequest{PlaylistId: 1}
	coms, err := c.GetCommentsByPlaylistId(ctx, &csr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(coms.GetComment()[1].Comment)
	*/
	c := pb.NewPlaylistServiceClient(conn)
	plr := pb.PlaylistRequest{Id: 1}
	pr, err := c.GetPlaylist(ctx, &plr)
	fmt.Println(pr.Name)
	fmt.Println(time.Unix(pr.CreatedAt.Seconds, int64(pr.CreatedAt.Nanos)))
	fmt.Println(pr.Comments.Comment[0].Comment)
}