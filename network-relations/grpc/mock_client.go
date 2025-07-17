package main

import (
	"context"
	"fmt"
	pbGrpcAccount "grpc/proto/account"
	pbGrpcRoom "grpc/proto/room"
)

var DftGrpcClientSet Client

func InitClient() {
	DftGrpcClientSet = &MockClient{}
}

type MockClient struct {
}

func (m *MockClient) GetAccountInfo(ctx context.Context, in *pbGrpcAccount.GetAccountInfoReq) (*pbGrpcAccount.GetAccountInfoResp, error) {

	fmt.Println("GetAccountInfo called with request:", in)
	return nil, nil
}

func (m *MockClient) NewAccountInfo(ctx context.Context, in *pbGrpcAccount.NewAccountInfoReq) (*pbGrpcAccount.NewAccountInfoResp, error) {

	fmt.Println("NewAccountInfo called with request:", in)
	return nil, nil
}

func (m *MockClient) DeleteAccount(ctx context.Context, in *pbGrpcAccount.DeleteAccountReq) (*pbGrpcAccount.DeleteAccountResp, error) {

	fmt.Println("DeleteAccount called with request:", in)
	return nil, nil
}

func (m *MockClient) NewRoom(ctx context.Context, in *pbGrpcRoom.NewRoomReq) (*pbGrpcRoom.NewRoomResp, error) {

	fmt.Println("NewRoom called with request:", in)
	return nil, nil
}

func (m *MockClient) GetRoom(ctx context.Context, in *pbGrpcRoom.GetRoomReq) (*pbGrpcRoom.GetRoomResp, error) {

	fmt.Println("GetRoom called with request:", in)
	return nil, nil
}
