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

	{
		res, err := cli.GetMarginPremium(
			context.Background(),
			&kabuspb.GetMarginPremiumRequest{SymbolCode: "1319"})
		log.Println("margin_premium", res, err)
	}
}
