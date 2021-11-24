package main

import (
	"context"
	"log"

	"google.golang.org/grpc/status"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"

	"google.golang.org/grpc"
)

var YourKabucomPassword = "" // カブコムのパスワードをセット

func main() {
	conn, err := grpc.DialContext(context.Background(), "localhost:18082", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	cli := kabuspb.NewKabusServiceClient(conn)

	// 制限値幅の下限を取得
	var lowerLimit float64
	{
		res, err := cli.GetSymbol(context.Background(), &kabuspb.GetSymbolRequest{SymbolCode: "1475", Exchange: kabuspb.Exchange_EXCHANGE_TOUSHOU})
		if err != nil {
			panic(err)
		}

		lowerLimit = res.LowerLimit
	}

	// 約定しないように制限値幅の下限で指値注文
	var orderID string
	{
		res, err := cli.SendStockOrder(context.Background(), &kabuspb.SendStockOrderRequest{
			Password:     YourKabucomPassword,
			SymbolCode:   "1475",
			Exchange:     kabuspb.StockExchange_STOCK_EXCHANGE_TOUSHOU,
			Side:         kabuspb.Side_SIDE_BUY,
			DeliveryType: kabuspb.DeliveryType_DELIVERY_TYPE_CASH,
			FundType:     kabuspb.FundType_FUND_TYPE_SUBSTITUTE_MARGIN,
			AccountType:  kabuspb.AccountType_ACCOUNT_TYPE_SPECIFIC,
			Quantity:     1,
			OrderType:    kabuspb.StockOrderType_STOCK_ORDER_TYPE_LO,
			Price:        lowerLimit,
			ExpireDay:    nil,
		})
		if err != nil {
			panic(err)
		}
		orderID = res.OrderId
	}

	// 注文取消し
	{
		_, err := cli.CancelOrder(context.Background(), &kabuspb.CancelOrderRequest{Password: YourKabucomPassword, OrderId: orderID})
		if err != nil {
			panic(err)
		}
	}

	// 失敗する注文取消し
	{
		_, err := cli.CancelOrder(context.Background(), &kabuspb.CancelOrderRequest{Password: YourKabucomPassword, OrderId: orderID})
		st, ok := status.FromError(err)
		if !ok {
			panic(err)
		}
		for _, detail := range st.Details() {
			switch e := detail.(type) {
			case *kabuspb.RequestError:
				log.Printf("status_code: %+v, body: %+v, code: %+v, message: %+v\n", e.StatusCode, e.Body, e.Code, e.Message)
			default:
				panic(err)
			}
		}
	}
}
