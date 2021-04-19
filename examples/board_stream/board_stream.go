package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

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
		_, err := cli.UnregisterAllSymbols(context.Background(), &kabuspb.UnregisterAllSymbolsRequest{})
		if err != nil {
			panic(err)
		}
	}

	var symbolCode string
	{
		res, err := cli.GetFutureSymbolCodeInfo(context.Background(), &kabuspb.GetFutureSymbolCodeInfoRequest{FutureCode: kabuspb.FutureCode_FUTURE_CODE_NK225_MINI, DerivativeMonth: timestamppb.New(time.Now().AddDate(0, 1, 0))})
		if err != nil {
			panic(err)
		}
		symbolCode = res.Code
	}

	// 登録
	{
		_, err := cli.RegisterSymbols(context.Background(), &kabuspb.RegisterSymbolsRequest{Symbols: []*kabuspb.RegisterSymbol{{SymbolCode: symbolCode, Exchange: kabuspb.Exchange_EXCHANGE_ALL_SESSION}}})
		if err != nil {
			panic(err)
		}
	}

	// stream開始
	{
		stream, err := cli.GetBoardsStreaming(context.Background(), &kabuspb.GetBoardsStreamingRequest{})
		if err != nil {
			panic(err)
		}

		for {
			board, err := stream.Recv()
			if err != nil {
				panic(err)
			}
			log.Println(board)
		}
	}
}
