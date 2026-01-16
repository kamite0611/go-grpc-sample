# go-grpc-sample

Go言語とgRPCの学習用リポジトリです。

## プロジェクト一覧

| ディレクトリ | 内容 |
|-------------|------|
| [grpc-tutorial](./grpc-tutorial) | gRPC入門 - 基本的なUnary RPC |

## 技術スタック

- **言語**: Go
- **通信プロトコル**: gRPC / Protocol Buffers
- **対象環境**: Google Cloud (Cloud Run, Cloud Spanner等)

## 始め方

```bash
# リポジトリをクローン
git clone <repository-url>
cd go-grpc-sample

# 各チュートリアルへ移動
cd grpc-tutorial
```

詳細は各ディレクトリのREADMEを参照してください。

## 学習ロードマップ

1. **grpc-tutorial** - gRPCの基礎（Unary RPC）
2. *(予定)* ストリーミングRPC
3. *(予定)* 認証・メタデータ
4. *(予定)* Cloud Spanner連携

## 参考リポジトリ

- [go-clean-arch-grpc](https://github.com/bxcodec/go-clean-arch-grpc) - クリーンアーキテクチャ + gRPCの実装例
- [go-clean-architecture](https://github.com/brandon-a-pinto/go-clean-architecture) - Goのクリーンアーキテクチャ実装例

詳細な学習リソースは [LEARNING_RESOURCES.md](./LEARNING_RESOURCES.md) を参照してください。

## 自作プロジェクトのアイデア

学習後に作ってみたいプロジェクト案は [PROJECT_IDEAS.md](./PROJECT_IDEAS.md) を参照してください。
