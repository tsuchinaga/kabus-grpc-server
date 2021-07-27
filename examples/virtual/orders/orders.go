package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.DialContext(context.Background(), "localhost:18082", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	cli := kabuspb.NewKabusServiceClient(conn)

	// 銘柄登録
	{
		res, err := cli.RegisterSymbols(context.Background(),
			&kabuspb.RegisterSymbolsRequest{Symbols: []*kabuspb.RegisterSymbol{{SymbolCode: "1475", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU}}, RequesterName: "virtual-test"},
		)
		if err != nil {
			panic(err)
		}
		log.Println(res)
	}

	// 成行エントリー
	{
		res, err := cli.SendStockOrder(context.Background(), &kabuspb.SendStockOrderRequest{
			Password:     "",
			SymbolCode:   "1475",
			Exchange:     kabuspb.StockExchange_STOCK_EXCHANGE_TOUSHOU,
			Side:         kabuspb.Side_SIDE_BUY,
			DeliveryType: kabuspb.DeliveryType_DELIVERY_TYPE_CASH,
			FundType:     kabuspb.FundType_FUND_TYPE_SUBSTITUTE_MARGIN,
			AccountType:  kabuspb.AccountType_ACCOUNT_TYPE_SPECIFIC,
			Quantity:     1,
			OrderType:    kabuspb.StockOrderType_STOCK_ORDER_TYPE_MO,
			Price:        0,
			ExpireDay:    nil,
			IsVirtual:    true,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	}

	// 注文一覧
	for {
		isContinue := false
		res, err := cli.GetOrders(context.Background(), &kabuspb.GetOrdersRequest{GetDetails: true, IsVirtual: true})
		if err != nil {
			panic(err)
		}
		for _, o := range res.Orders {
			log.Println(o)
			if o.State != kabuspb.State_STATE_DONE {
				isContinue = true
			}
		}
		if !isContinue {
			break
		}
		<-time.After(30 * time.Second)
	}
}
