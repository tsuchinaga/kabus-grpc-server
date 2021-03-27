# kabus-grpc-server

kabusapiを叩くためのgRPCサーバー

複数ツールからkabusapiを利用するためにリクエストの受け口を一つにするために使う。

## build

`$ protoc --go_out=./kabuspb --go-grpc_out=./kabuspb kabuspb/kabus.proto`

* `--go_out`: protobufのmessageとかenumとかが吐かれる先
* `--go-grpc_out`: protobufのserviceから作られるserverとかclientが吐かれる先

## run

`$ go run main.go -e=p -p=Password1234`

* `e`: 環境。本番がp、検証がd
* `p`: パスワード。
