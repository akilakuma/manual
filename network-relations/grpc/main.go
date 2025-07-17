package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pbAccount "grpc/proto/account"
	pbRoom "grpc/proto/room"
	"net"
	"runtime"
	"runtime/debug"
)

func main() {

	InitClient()

	ctx, _ := context.WithCancel(context.Background())
	lis, err := net.Listen("tcp", ":"+"50052")
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	grpcServer := regGrpcServer()

	fmt.Printf("gRPC server listening at %v\n", lis.Addr())

	go func() {
		<-ctx.Done()
		// 關閉client連線 TODO
		//client.DftGrpcClientSet.Close()
		fmt.Println("Shutting down gRPC server...")
		grpcServer.GracefulStop()

		fmt.Println("gRPC server stopped")

	}()

	if err := grpcServer.Serve(lis); err != nil && err != grpc.ErrServerStopped {
		fmt.Errorf("failed to serve: %v", err)
	}
}

// regGrpcServer 處理註冊不同的proto server
func regGrpcServer() *grpc.Server {

	var (
		serverSet  = grpcServerIns // 自己聚合client對象
		grpcServer = grpc.NewServer(
			grpc.UnaryInterceptor(recoveryInterceptor),
		)
	)

	// 在同一個 server 上註冊所有服務
	pbAccount.RegisterAccountServer(grpcServer, serverSet)
	pbRoom.RegisterRoomServer(grpcServer, serverSet)

	return grpcServer
}

// 定義一個 UnaryInterceptor 來捕捉 panic
func recoveryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			// 可加入詳細 log 記錄
			err = status.Errorf(codes.Internal, "internal server error, may missing neccenary paramter")
			eMsg := fmt.Sprintf("%+v", err)
			fmt.Println("recovering from panic", eMsg)

			_, file, line, ok := runtime.Caller(4)
			if ok {
				fmt.Printf("panic recovered at %s:%d, error: %v\n", file, line, r)
			} else {
				stack := debug.Stack()
				fmt.Printf("panic recovered: %v\n%s\n", r, stack)
			}
		}

	}()
	return handler(ctx, req)
}
