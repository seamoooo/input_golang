package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	// https://qiita.com/fetaro/items/31b02b940ce9ec579baf
	// 内部パッケージを呼ぶ際はgo modに記載したpfg名を利用する
	hellopb "github.com/seamoooo/input_golang/FistTimeGrpc/pkg/grpc"
)

type myServer struct {
	hellopb.UnimplementedGreetingServiceServer
}

// HelloRequest型のリクエストを受け取って、HelloResponse型のレスポンスを返す」Helloメソッド
func (s *myServer) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {

	stat := status.New(codes.Unknown, "unkown error occurred")
	stat, _ = stat.WithDetails(&errdetails.DebugInfo{
		Detail: "detail reason of err",
	})
	err := stat.Err()

	// 直接HelloResponse型をretrunしている。
	return nil, err
}

func (s *myServer) HelloServerStream(req *hellopb.HelloRequest, stream hellopb.GreetingService_HelloServerStreamServer) error {
	// 5回レスポンスを返す
	resCount := 5

	for i := 0; i < resCount; i++ {
		// レスポンスを返したいときには、Sendメソッドの引数にHelloResponse型を渡すことでそれがクライアントに送信
		if err := stream.Send(&hellopb.HelloResponse{
			Message: fmt.Sprintf("[%d] Hello, %s!", i, req.GetName()),
		}); err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	}

	// return文でメソッドを終了させる
	return nil
}

func (s *myServer) HelloClientStream(stream hellopb.GreetingService_HelloClientStreamServer) error {
	nameList := make([]string, 0)

	for {
		// 複数回のリクエストをserver側で受ける必要があるので、
		// helloRequest型をRecvを経由して受け取る
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			message := fmt.Sprintf("hello. %v", nameList)

			// SendAndCloseメソッドを呼ぶことでレスポンスを返す
			return stream.SendAndClose(&hellopb.HelloResponse{
				Message: message,
			})
		}
		if err != nil {
			return err
		}
		nameList = append(nameList, req.GetName())
	}
}

func (s *myServer) HelloBiStream(stream hellopb.GreetingService_HelloBiStreamServer) error {
	for {
		req, err := stream.Recv()

		// 終端処理
		if errors.Is(err, io.EOF) {
			return nil
		}
		// エラーハンドリング
		if err != nil {
			return nil
		}
		message := fmt.Sprintf("Hello %v", req.GetName())

		if err := stream.Send(&hellopb.HelloResponse{
			Message: message,
		}); err != nil {
			return err
		}
	}
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
