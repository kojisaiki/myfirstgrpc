package service

import (
	"context"
	"errors"

	pb "../../pb"
)

type MyCatService struct {
}

func (s *MyCatService) GetMyCat(ctx context.Context, message *pb.GetMyCatMessage) (*pb.GetMyCatResponse, error) {
	switch message.TargetCat {
	case "tama":
		return &pb.GetMyCatResponse{
			Name: "tama",
			Kind: "mainecoon",
		}, nil
	case "mike":
		return &pb.GetMyCatResponse{
			Name: "mike",
			Kind: "Norweian Forest Cat",
		}, nil
	}
	return nil, errors.New("Not Found YourCat")
}
