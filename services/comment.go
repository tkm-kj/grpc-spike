package services

import (
	"context"
	pb "github.com/tkm-kj/grpc-spike/proto"
)

type CommentService struct {}

func(s *CommentService) GetComment(ctx context.Context, in *pb.CommentRequest) (*pb.Comment, error) {
	c1 := pb.Comment{Id: 1, PlaylistId: 1, UserName: "Bob", Comment: "面白かった!"}

	return &c1, nil
}

func (s *CommentService) GetCommentsByPlaylistId(ctx context.Context, in *pb.CommentsRequest) (*pb.Comments, error) {
	c1 := pb.Comment{Id: 1, PlaylistId: 1, UserName: "Bob", Comment: "面白かった!"}
	c2 := pb.Comment{Id: 2, PlaylistId: 1, UserName: "Alice", Comment: "最&高"}

	cs := pb.Comments{Comment: []*pb.Comment{&c1, &c2}}

	return &cs, nil
}
