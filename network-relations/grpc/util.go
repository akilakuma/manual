package main

import (
	pbGrpcAccount "grpc/proto/account"
	pbGrpcRoom "grpc/proto/room"
)

type GrpcServer struct {
	pbGrpcAccount.UnimplementedAccountServer
	pbGrpcRoom.UnimplementedRoomServer
}

var grpcServerIns *GrpcServer

func init() {
	grpcServerIns = GetGrpcServer()
}

func GetGrpcServer() *GrpcServer {
	return &GrpcServer{}
}
