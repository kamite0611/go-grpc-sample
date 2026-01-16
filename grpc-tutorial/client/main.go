package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "grpc-tutorial/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// サーバーに接続
	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("接続失敗: %v", err)
	}
	defer conn.Close()

	// クライアント作成
	client := pb.NewGreeterClient(conn)

	// 名前を引数から取得（デフォルトは "World"）
	name := "World"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// タイムアウト付きコンテキスト
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// RPCを呼び出し
	res, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("SayHello失敗: %v", err)
	}

	log.Printf("サーバーからの応答: %s", res.GetMessage())
}
