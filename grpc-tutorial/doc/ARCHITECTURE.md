# gRPC Tutorial アーキテクチャ解説

このドキュメントでは、grpc-tutorialプロジェクト全体の動きを解説します。

## 全体構成

```
┌─────────────────────────────────────────────────────────────────┐
│                        grpc-tutorial                            │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────┐                        ┌─────────────┐        │
│  │   Client    │  ──── gRPC通信 ────▶   │   Server    │        │
│  │  (client/)  │       (HTTP/2)         │  (server/)  │        │
│  └─────────────┘                        └─────────────┘        │
│         │                                      │                │
│         │                                      │                │
│         ▼                                      ▼                │
│  ┌─────────────────────────────────────────────────────┐       │
│  │              pb/ (生成されたGoコード)                │       │
│  │  - greeter.pb.go      (メッセージ定義)              │       │
│  │  - greeter_grpc.pb.go (gRPCクライアント/サーバー)   │       │
│  └─────────────────────────────────────────────────────┘       │
│                            ▲                                    │
│                            │ protoc で生成                      │
│                            │                                    │
│                   ┌─────────────────┐                          │
│                   │ proto/          │                          │
│                   │ greeter.proto   │                          │
│                   │ (サービス定義)   │                          │
│                   └─────────────────┘                          │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

## 処理フロー

```
┌────────┐          ┌────────┐          ┌────────┐
│ Client │          │  gRPC  │          │ Server │
└───┬────┘          └───┬────┘          └───┬────┘
    │                   │                   │
    │ 1. 接続要求        │                   │
    │ ─────────────────▶│                   │
    │                   │ 2. TCP接続確立     │
    │                   │ ─────────────────▶│
    │                   │                   │
    │ 3. SayHello RPC   │                   │
    │   (HelloRequest)  │                   │
    │ ─────────────────▶│ 4. リクエスト転送  │
    │                   │ ─────────────────▶│
    │                   │                   │
    │                   │ 5. 処理実行        │
    │                   │   "こんにちは..."   │
    │                   │                   │
    │                   │ 6. HelloResponse  │
    │                   │ ◀─────────────────│
    │ 7. レスポンス受信  │                   │
    │ ◀─────────────────│                   │
    │                   │                   │
    │ 8. 接続クローズ    │                   │
    │ ─────────────────▶│                   │
    │                   │                   │
```

## 各コンポーネントの役割

### 1. proto/greeter.proto（サービス定義）

Protocol Buffersでサービスのインターフェースを定義します。

```protobuf
service Greeter {
  rpc SayHello (HelloRequest) returns (HelloResponse);
}
```

**役割:**
- メッセージ構造の定義（`HelloRequest`, `HelloResponse`）
- RPCメソッドのシグネチャ定義
- クライアント・サーバー間の契約（Contract）

### 2. pb/（生成コード）

`protoc`コマンドで自動生成されるGoコードです。

| ファイル | 内容 |
|---------|------|
| `greeter.pb.go` | メッセージ型（`HelloRequest`, `HelloResponse`）のGo構造体 |
| `greeter_grpc.pb.go` | `GreeterClient`インターフェース、`GreeterServer`インターフェース |

**生成コマンド:**
```bash
make proto
```

### 3. server/main.go（サーバー実装）

gRPCサーバーの実装です。

```go
// サーバー構造体
type server struct {
    pb.UnimplementedGreeterServer  // 前方互換性のため埋め込み
}

// RPCメソッドの実装
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
    return &pb.HelloResponse{
        Message: fmt.Sprintf("こんにちは、%s さん！", req.GetName()),
    }, nil
}
```

**処理の流れ:**
1. TCPリスナーを作成（ポート50051）
2. gRPCサーバーインスタンスを作成
3. `Greeter`サービスを登録
4. サーバーを起動してリクエストを待機

### 4. client/main.go（クライアント実装）

gRPCクライアントの実装です。

```go
// サーバーに接続
conn, err := grpc.NewClient("localhost:50051",
    grpc.WithTransportCredentials(insecure.NewCredentials()),
)

// クライアント作成
client := pb.NewGreeterClient(conn)

// RPC呼び出し
res, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})
```

**処理の流れ:**
1. サーバーへの接続を確立
2. `GreeterClient`を作成
3. タイムアウト付きコンテキストを設定（1秒）
4. `SayHello` RPCを呼び出し
5. レスポンスを受信・表示
6. 接続をクローズ

## 実行例

**ターミナル1（サーバー起動）:**
```bash
$ make server
gRPCサーバー起動中... ポート:50051
リクエスト受信: Alice
```

**ターミナル2（クライアント実行）:**
```bash
$ make client ARGS="Alice"
サーバーからの応答: こんにちは、Alice さん！
```

## 重要なポイント

### UnimplementedGreeterServer の埋め込み

```go
type server struct {
    pb.UnimplementedGreeterServer
}
```

この埋め込みにより、将来protoファイルに新しいRPCメソッドが追加されても、既存のサーバー実装がコンパイルエラーにならずに動作します（前方互換性）。

### insecure.NewCredentials()

```go
grpc.WithTransportCredentials(insecure.NewCredentials())
```

開発環境ではTLS無しで接続するための設定です。本番環境では適切なTLS設定が必要です。

### context.WithTimeout

```go
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()
```

RPCにタイムアウトを設定することで、サーバーが応答しない場合でもクライアントがハングしません。本番環境では必須のパターンです。
