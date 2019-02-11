package services

import (
	"context"
	"time"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "github.com/tkm-kj/grpc-spike/proto"
)

type PlaylistService struct {}

func(s *PlaylistService) GetPlaylist(context.Context, *pb.PlaylistRequest) (*pb.Playlist, error) {
	c1 := pb.Comment{Id: 1, PlaylistId: 1, UserName: "Bob", Comment: "面白かった!"}
	c2 := pb.Comment{Id: 2, PlaylistId: 1, UserName: "Alice", Comment: "最&高"}

	cs := pb.Comments{Comment: []*pb.Comment{&c1, &c2}}
	p := pb.Playlist{Id: 1, Order: 1, Name: "とても楽しい回", CreatedAt: &timestamp.Timestamp{Seconds: time.Now().Unix(), Nanos: int32(time.Now().UnixNano())}, Comments: &cs}
	return &p, nil
}