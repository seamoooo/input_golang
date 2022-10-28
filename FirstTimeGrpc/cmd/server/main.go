package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// https://qiita.com/fetaro/items/31b02b940ce9ec579baf
	// 内部パッケージを呼ぶ際はgo modに記載したpfg名を利用する
	hellopb "github.com/seamoooo/input_golang/FistTimeGrpc/pkg/grpc"
)

type myServer struct {
	hellopb.UnimplementedGreetingServiceServer
}

// HelloRequest型のリクエストを受け取って、HelloResponse型のレスポンスを返す」Helloメソッド
func (s *myServer) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	return &hellopb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s", req.GetName()),
	}, nil
}

func NewMyServer() *myServer {
	return &myServer{}
}

func main() {
	// 8080portのlisnerを作成
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// grpcサーバーを作成
	// gRPCサーバーに対応する型がgoogle.golang.org/grpcパッケージに用意
	s := grpc.NewServer()

	// GreetingServiceをgRPCサーバーに登録する
	hellopb.RegisterGreetingServiceServer(s, NewMyServer())

	// サーバーリフレクションの設定
	reflection.Register(s)

	// 作成したgRPCサーバーを8080ポートで稼働させる
	go func() {
		log.Printf("Starting GRPC Sserver: %v", port)
		s.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	log.Printf("stop gRPC server")
	s.GracefulStop()
}
