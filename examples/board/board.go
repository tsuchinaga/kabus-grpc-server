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
		res, err := cli.GetBoard(
			context.Background(),
			&kabuspb.GetBoardRequest{
				SymbolCode: "1320",
				Exchange:   kabuspb.Exchange_EXCHANGE_TOUSHOU,
			})
		log.Println("board", res, err)
	}
}
