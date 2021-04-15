package main

import (
	"context"
	"log"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.DialContext(context.Background(), "localhost:18082", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	cli := kabuspb.NewKabusServiceClient(conn)

	// 注文一覧
	{
		res, err := cli.GetOrders(context.Background(), &kabuspb.GetOrdersRequest{GetDetails: true})
		if err != nil {
			panic(err)
		}
		log.Printf("%+v\n", res)
	}
}
