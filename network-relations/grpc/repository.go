package main

import (
	"context"
	pbGrpcAccount "grpc/proto/account"
	pbGrpcRoom "grpc/proto/room"
)

type Client interface {
	GetAccountInfo(ctx context.Context, in *pbGrpcAccount.GetAccountInfoReq) (*pbGrpcAccount.GetAccountInfoResp, error)
	NewAccountInfo(ctx context.Context, in *pbGrpcAccount.NewAccountInfoReq) (*pbGrpcAccount.NewAccountInfoResp, error)
	DeleteAccount(ctx context.Context, in *pbGrpcAccount.DeleteAccountReq) (*pbGrpcAccount.DeleteAccountResp, error)
	NewRoom(ctx context.Context, in *pbGrpcRoom.NewRoomReq) (*pbGrpcRoom.NewRoomResp, error)
	GetRoom(ctx context.Context, in *pbGrpcRoom.GetRoomReq) (*pbGrpcRoom.GetRoomResp, error)
}
