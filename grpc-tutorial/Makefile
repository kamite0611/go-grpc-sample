.PHONY: proto build server client clean

# protoファイルからGoコードを生成
proto:
	protoc --go_out=. --go-grpc_out=. proto/greeter.proto

# 全てビルド
build: build-server build-client

# サーバーをビルド
build-server:
	go build -o bin/server ./server

# クライアントをビルド
build-client:
	go build -o bin/client ./client

# サーバーを実行
server:
	go run ./server

# クライアントを実行
client:
	go run ./client

# クリーンアップ
clean:
	rm -rf bin/

# 依存関係をインストール
deps:
	go mod tidy
