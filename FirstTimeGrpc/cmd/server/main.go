package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	// https://qiita.com/fetaro/items/31b02b940ce9ec579baf
	// 内部パッケージを呼ぶ際はgo modに記載したpfg名を利用する
	hellopb "github.com/seamoooo/input_golang/FistTimeGrpc/pkg/grpc"
)

type myServer struct {
	hellopb.UnimplementedGreetingServiceServer
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
	hellopb.RegisterGreetingServiceServer(s)

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
