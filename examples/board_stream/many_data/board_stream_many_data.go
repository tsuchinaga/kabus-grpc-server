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

	// 登録銘柄クリア
	{
		_, err := cli.UnregisterAllSymbols(context.Background(), &kabuspb.UnregisterAllSymbolsRequest{})
		if err != nil {
			panic(err)
		}
	}

	symbols := []string{
		"4151", "4502", "3105", "6479", "7201", "7202", "4543", "4902", "9412", "9432",
		"7186", "8303", "8253", "8697", "8601", "8604", "8630", "8725", "1332", "1333",
		"2002", "2269", "3086", "3099", "2413", "2432", "1605", "3101", "3103", "3861",
		"3863", "3405", "3407", "5019", "5020", "5101", "5108", "5201", "5202", "5401",
		"5406", "3436", "5703", "2768", "8001", "1721", "1801", "5631", "6103", "7003",
	}

	// 登録
	{
		registerSymbols := make([]*kabuspb.RegisterSymbol, len(symbols))
		for i, s := range symbols {
			registerSymbols[i] = &kabuspb.RegisterSymbol{SymbolCode: s, Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}
		}
		_, err := cli.RegisterSymbols(context.Background(), &kabuspb.RegisterSymbolsRequest{Symbols: registerSymbols})
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
