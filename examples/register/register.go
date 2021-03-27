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
		res, err := cli.RegisterSymbols(
			context.Background(),
			&kabuspb.RegisterSymbolsRequest{Symbols: []*kabuspb.RegisterSymbol{{Symbol: "9433", Exchange: kabuspb.Exchange_TOUSHOU}}})
		log.Println("register 1", res, err)
	}

	{
		res, err := cli.RegisterSymbols(
			context.Background(),
			&kabuspb.RegisterSymbolsRequest{Symbols: []*kabuspb.RegisterSymbol{
				{Symbol: "1320", Exchange: kabuspb.Exchange_TOUSHOU},
				{Symbol: "1329", Exchange: kabuspb.Exchange_TOUSHOU},
			}})
		log.Println("register 2", res, err)
	}

	{
		res, err := cli.RegisterSymbols(
			context.Background(),
			&kabuspb.RegisterSymbolsRequest{Symbols: []*kabuspb.RegisterSymbol{{Symbol: "166060018", Exchange: kabuspb.Exchange_ALL_SESSION}}})
		log.Println("register 3", res, err)
	}

	{
		res, err := cli.GetRegisteredSymbols(context.Background(), &kabuspb.GetRegisteredSymbolsRequest{})
		log.Println("get", res, err)
	}

	{
		res, err := cli.UnregisterSymbols(context.Background(), &kabuspb.UnregisterSymbolsRequest{Symbols: []*kabuspb.RegisterSymbol{
			{Symbol: "1329", Exchange: kabuspb.Exchange_TOUSHOU},
		}})
		log.Println("unregister", res, err)
	}

	{
		res, err := cli.UnregisterAllSymbols(context.Background(), &kabuspb.UnregisterAllSymbolsRequest{})
		log.Println("unregister all", res, err)
	}
}
