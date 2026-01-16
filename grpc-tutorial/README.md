# gRPC Tutorial

Go言語でgRPCを学ぶためのサンプルプロジェクトです。

## プロジェクト構成

```
grpc-tutorial/
├── proto/              # Protocol Buffers定義
│   └── greeter.proto
├── pb/                 # 生成されたGoコード
│   ├── greeter.pb.go
│   └── greeter_grpc.pb.go
├── server/             # gRPCサーバー
│   └── main.go
├── client/             # gRPCクライアント
│   └── main.go
└── bin/                # ビルド済みバイナリ
```

## 必要なツール

- Go 1.20以上
- Protocol Buffers コンパイラ (`protoc`)
- Go用プラグイン (`protoc-gen-go`, `protoc-gen-go-grpc`)

### インストール

```bash
# macOS (Homebrew)
brew install go protobuf

# Go用プラグイン
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## ビルド方法

```bash
# 依存関係の取得
go mod tidy

# protocでコード生成
protoc --go_out=pb --go_opt=paths=source_relative \
       --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
       -I proto proto/greeter.proto

# バイナリのビルド
go build -o bin/server ./server
go build -o bin/client ./client
```

## 実行方法

```bash
# ターミナル1: サーバー起動
./bin/server
# 出力: gRPCサーバー起動中... ポート:50051

# ターミナル2: クライアント実行
./bin/client [名前]
# 例: ./bin/client Tanaka
# 出力: サーバーからの応答: こんにちは、Tanaka さん！
```

## API仕様

### Greeter サービス

| メソッド | リクエスト | レスポンス | 説明 |
|---------|-----------|-----------|------|
| SayHello | HelloRequest | HelloResponse | 名前を受け取り挨拶を返す |

### メッセージ型

```protobuf
message HelloRequest {
  string name = 1;
}

message HelloResponse {
  string message = 1;
}
```

## 参考資料

- [gRPC公式ドキュメント](https://grpc.io/docs/languages/go/)
- [Protocol Buffers](https://protobuf.dev/)
- [grpc-go GitHub](https://github.com/grpc/grpc-go)
