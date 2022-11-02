package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	hellopb "github.com/seamoooo/input_golang/FistTimeGrpc/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var (
	scanner *bufio.Scanner
	client  hellopb.GreetingServiceClient
)

func main() {
	fmt.Println("grpc client started.")

	// 標準入力から文字列を受け取るスキャナ
	scanner = bufio.NewScanner(os.Stdin)

	address := "localhost:8080"

	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()), //コネクションでSSL/TLSを使用しない
		grpc.WithBlock(), //コネクションが確立されるまで待機する(同期処理をする)
	)

	if err != nil {
		log.Fatal("connnection faild")
		return
	}

	defer conn.Close()

	client = hellopb.NewGreetingServiceClient(conn)

	for {
		fmt.Println("1: send Request")
		fmt.Println("2: HelloServerStream")
		fmt.Println("3: HelloClientStream")
		fmt.Println("4: HelloBiStream")
		fmt.Println("5: exit")
		fmt.Print("please enter")

		scanner.Scan()
		in := scanner.Text()

		switch in {
		case "1":
			Hello()
		case "2":
			HelloServerStream()
		case "3":
			HelloClientStream()
		case "4":
			HelloBiStream()
		case "5":
			fmt.Println("bye")
			goto M
		}
	}
M:
}

func Hello() {
	fmt.Println("please enter your name")
	scanner.Scan()

	name := scanner.Text()

	req := &hellopb.HelloRequest{
		Name: name,
	}

	res, err := client.Hello(context.Background(), req)

	if err != nil {
		if stat, ok := status.FromError(err); ok {
			fmt.Printf("codeL: %s\n", stat.Code())
			fmt.Printf("message: %s\n", stat.Message())
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(res.GetMessage())
	}
}

func HelloServerStream() {
	fmt.Println("please enter name")
	scanner.Scan()

	name := scanner.Text()

	req := &hellopb.HelloRequest{
		Name: name,
	}

	// 1.クライアントが持つHelloServerStreamメソッドを呼んで、サーバーからレスポンスが送られてくるストリーム(GreetingService_HelloServerStreamClientインターフェース型)を取得
	stream, err := client.HelloServerStream(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		// ストリームのRecvメソッドを呼ぶことでレスポンスを得る
		res, err := stream.Recv()

		// Recvメソッドでレスポンスを受け取るとき、これ以上受け取るレスポンスがないという状態なら、
		// 第一戻り値にはnil、第二戻り値のerrにはio.EOFが格納されています。
		if errors.Is(err, io.EOF) {
			fmt.Println("all the response have already received")
			break
		}

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	}
}

func HelloClientStream() {
	// サーバーに複数回リクエストを送るためのストリームを得る
	stream, err := client.HelloClientStream(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	sendCount := 5

	fmt.Printf("Please enter %d names.\n", sendCount)

	for i := 0; i < sendCount; i++ {
		scanner.Scan()
		name := scanner.Text()

		// ストリームを通じてリクエストを送信
		if err := stream.Send(&hellopb.HelloRequest{
			Name: name,
		}); err != nil {
			fmt.Println(err)
			return
		}
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.GetMessage())
	}
}

func HelloBiStream() {
	stream, err := client.HelloBiStream(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	sendNum := 5
	fmt.Printf("Please ender %d names \n", sendNum)

	var senfEnd, recvEnd bool
	sendCount := 0

	for !(senfEnd && recvEnd) {
		//送信処理
		if !senfEnd {
			scanner.Scan()
			name := scanner.Text()

			sendCount++
			if err := stream.Send(&hellopb.HelloRequest{
				Name: name,
			}); err != nil {
				fmt.Println(err)
				senfEnd = true
			}

			if sendCount == sendNum {
				senfEnd = true
				if err := stream.CloseSend(); err != nil {
					fmt.Println(err)
				}
			}
		}

		// 受信処理
		if !recvEnd {
			if res, err := stream.Recv(); err != nil {
				if !errors.Is(err, io.EOF) {
					fmt.Println(err)
				}
				recvEnd = true
			} else {
				fmt.Println(res.GetMessage())
			}
		}
	}
}
