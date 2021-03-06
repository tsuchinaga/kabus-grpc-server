# kabus-grpc-server

![pipeline.svg](https://gitlab.com/tsuchinaga/kabus-grpc-server/badges/master/pipeline.svg)
![coverage.svg](https://gitlab.com/tsuchinaga/kabus-grpc-server/badges/master/coverage.svg)

kabusapiを叩くためのgRPCサーバー

複数ツールからkabusapiを利用するためにリクエストの受け口を一つにするために使う。

## build

`$ protoc --go_out=./kabuspb --go-grpc_out=./kabuspb kabuspb/kabus.proto`

* `--go_out`: protobufのmessageとかenumとかが吐かれる先
* `--go-grpc_out`: protobufのserviceから作られるserverとかclientが吐かれる先

## run

`$ go run cmd/kabus_grpc_server.go -e=p -p=Password1234`

* `e`: 環境。本番がp、検証がd。デフォルトd
* `p`: パスワード。
* `port`: ポート。デフォルト18082

## 定義

[protobufファイル](./kabuspb/kabus.proto)

[protobufドキュメント](https://tsuchinaga.gitlab.io/kabus-grpc-server/#kabuspb.KabusService)

## 注意

まだ開発中で全機能実装できていません。

また、kabusapiを叩く以外の機能も追加すると思います。

[github.com/tsuchinaga/kabus-grpc-server](https://github.com/tsuchinaga/kabus-grpc-server) にミラーリングしていますが、オリジナルは [gitlab.com/tsuchinaga/kabus-grpc-server](https://gitlab.com/tsuchinaga/kabus-grpc-server) にあります。
