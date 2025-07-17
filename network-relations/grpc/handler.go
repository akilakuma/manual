package main

import (
	"context"
	"fmt"
	pbGrpcAccount "grpc/proto/account"
	pbGrpcRoom "grpc/proto/room"
)

func (g *GrpcServer) GetAccountInfo(ctx context.Context, in *pbGrpcAccount.GetAccountInfoReq) (*pbGrpcAccount.GetAccountInfoResp, error) {

	fmt.Println("GetAccountInfo called with request:", in)
	return nil, nil
}

func (g *GrpcServer) NewAccountInfo(ctx context.Context, in *pbGrpcAccount.NewAccountInfoReq) (*pbGrpcAccount.NewAccountInfoResp, error) {

	fmt.Println("NewAccountInfo called with request:", in)
	return nil, nil
}

func (g *GrpcServer) DeleteAccount(ctx context.Context, in *pbGrpcAccount.DeleteAccountReq) (*pbGrpcAccount.DeleteAccountResp, error) {

	fmt.Println("DeleteAccount called with request:", in)
	return nil, nil
}

func (g *GrpcServer) NewRoom(ctx context.Context, in *pbGrpcRoom.NewRoomReq) (*pbGrpcRoom.NewRoomResp, error) {

	fmt.Println("NewRoom called with request:", in)
	return nil, nil
}

func (g *GrpcServer) GetRoom(ctx context.Context, in *pbGrpcRoom.GetRoomReq) (*pbGrpcRoom.GetRoomResp, error) {

	fmt.Println("GetRoom called with request:", in)
	return nil, nil
}
