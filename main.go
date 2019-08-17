package main

import (
	"context"
	proto "github.com/ramfan/internal_api/proto"
)

type server struct {}

func main(){}

func (s *server) CheckAuth(ctx context.Context, req *proto.CheckAuthRequest) (proto.CheckAuthResponse, error) {
	session := req.GetSession()

	if session == "trueSession" {
		return proto.CheckAuthResponse{IsAuth: true}, nil
	}

	return proto.CheckAuthResponse{IsAuth: false}, nil
}
