# 学習リソース

Go + gRPC + Clean Architecture を学ぶためのおすすめサンプルプロジェクト集です。

## 初心者向け（まずはここから）

| リポジトリ | 特徴 |
|-----------|------|
| [bxcodec/go-clean-arch-grpc](https://github.com/bxcodec/go-clean-arch-grpc) | シンプルで分かりやすい。解説記事付きで入門に最適 |
| [brandon-a-pinto/go-clean-architecture](https://github.com/brandon-a-pinto/go-clean-architecture) | gRPC + PostgreSQL + Docker。構成がシンプル |

## 中級〜実践向け（AleksK1NGシリーズ）⭐ 特におすすめ

AleksK1NGさんのリポジトリは、gRPC、PostgreSQL、Redis、Prometheus、Grafana、Jaegerトレーシングなど、本番環境に近い構成で非常に学習価値が高いです。

| リポジトリ | 技術スタック | 解説記事 |
|-----------|-------------|----------|
| [Go-GRPC-Auth-Microservice](https://github.com/AleksK1NG/Go-GRPC-Auth-Microservice) | gRPC + PostgreSQL + Redis + Jaeger | [DEV.to記事](https://dev.to/aleksk1ng/go-grpc-clean-architecture-microservice-with-prometheus-grafana-monitoring-and-jaeger-opentracing-51om) |
| [Go-gRPC-RabbitMQ-microservice](https://github.com/AleksK1NG/Go-gRPC-RabbitMQ-microservice) | gRPC + RabbitMQ + メール送信 | [DEV.to記事](https://dev.to/aleksk1ng/go-rabbitmq-and-grpc-clean-architecture-microservice-2kdn) |
| [Go-CQRS-Kafka-gRPC-Microservices](https://github.com/AleksK1NG/Go-CQRS-Kafka-gRPC-Microservices) | Kafka + CQRS + gRPC | [DEV.to記事](https://dev.to/aleksk1ng/go-kafka-and-grpc-clean-architecture-cqrs-microservices-with-jaeger-tracing-45bj) |

### AleksK1NGシリーズの良い点

- Docker Composeですぐ動かせる（`make local` → `make run`）
- 解説記事が充実している
- 監視・トレーシングまで含んでいて実務に近い
- テストコードも参考になる

## その他の良質なリポジトリ

| リポジトリ | 特徴 |
|-----------|------|
| [herryg91/go-clean-architecture](https://github.com/herryg91/go-clean-architecture) | 詳細なREADMEでアーキテクチャの説明が充実 |
| [khannedy/golang-clean-architecture](https://github.com/khannedy/golang-clean-architecture) | HTTP/gRPC/Messagingの実装例あり |

## 学習の進め方

```
1. bxcodec/go-clean-arch-grpc で基本構造を理解
   ↓
2. AleksK1NG/Go-GRPC-Auth-Microservice で実践的な実装を学ぶ
   ↓
3. 自分で小さなプロジェクトを作ってみる
```
