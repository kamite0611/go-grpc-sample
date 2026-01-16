package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc-tutorial/pb"

	"google.golang.org/grpc"
)

// server は GreeterServer インターフェースを実装する構造体
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello は Greeter サービスの RPC メソッドを実装
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("リクエスト受信: %s", req.GetName())
	return &pb.HelloResponse{
		Message: fmt.Sprintf("こんにちは、%s さん！", req.GetName()),
	}, nil
}

func main() {
	// TCPリスナーを作成（ポート50051）
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("リッスン失敗: %v", err)
	}

	// gRPCサーバーを作成
	s := grpc.NewServer()

	// サービスを登録
	pb.RegisterGreeterServer(s, &server{})

	log.Println("gRPCサーバー起動中... ポート:50051")

	// サーバー起動
	if err := s.Serve(lis); err != nil {
		log.Fatalf("サーバー起動失敗: %v", err)
	}
}
