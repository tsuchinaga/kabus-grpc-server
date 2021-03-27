package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"gitlab.com/tsuchinaga/kabus-grpc-server/di"
	"gitlab.com/tsuchinaga/kabus-grpc-server/infra"
	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"
	"google.golang.org/grpc"
)

func main() {
	// パスワード、本番か検証か
	isProd := flag.String("e", "d", "environment d(develop) or p(production)")
	password := flag.String("p", "", "password")
	flag.Parse()

	if *password == "" {
		fmt.Println("-p is required")
		return
	}

	// 設定の初期化
	infra.InitSetting(*isProd == "p", *password)

	// サーバーの起動
	ln, err := net.Listen("tcp", ":18082")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	kabuspb.RegisterKabusServiceServer(s, di.InjectedServer())

	if err := s.Serve(ln); err != nil {
		log.Fatalln(err)
	}
}
