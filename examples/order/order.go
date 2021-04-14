package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"gitlab.com/tsuchinaga/kabus-grpc-server/kabuspb"

	"google.golang.org/grpc"
)

var YourKabucomPassword = "" // カブコムのパスワードをセット

// 売買手数料のかからない銘柄でエントリー -> 約定確認 -> エグジット -> 約定確認するだけのプログラム
// 1ティックが1円なので、2ティック分くらい損する。
func main() {
	conn, err := grpc.DialContext(context.Background(), "localhost:18082", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	cli := kabuspb.NewKabusServiceClient(conn)

	// 成行エントリー
	var entryOrderID string
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
			OrderType:    kabuspb.StockOrderType_STOCK_ORDER_TYPE_MO,
			Price:        0,
			ExpireDay:    nil,
		})
		if err != nil {
			panic(err)
		}
		entryOrderID = res.OrderId
	}

	checkContracted := func(orderID string, orders *kabuspb.Orders) bool {
		for _, order := range orders.Orders {
			if order.Id == orderID && order.OrderState == kabuspb.OrderState_ORDER_STATE_DONE {
				return true
			}
		}
		return false
	}

	// 約定確認
	{
		for {
			res, err := cli.GetOrders(context.Background(), &kabuspb.GetOrdersRequest{})
			if err != nil {
				panic(err)
			}

			if checkContracted(entryOrderID, res) {
				fmt.Println("エントリー注文の約定を確認しました")
				break
			}
			<-time.After(time.Second)
		}
	}

	// 成行エグジット
	var exitOrderID string
	{
		res, err := cli.SendStockOrder(context.Background(), &kabuspb.SendStockOrderRequest{
			Password:     YourKabucomPassword,
			SymbolCode:   "1475",
			Exchange:     kabuspb.StockExchange_STOCK_EXCHANGE_TOUSHOU,
			Side:         kabuspb.Side_SIDE_SELL,
			DeliveryType: kabuspb.DeliveryType_DELIVERY_TYPE_CASH,
			FundType:     kabuspb.FundType_FUND_TYPE_SUBSTITUTE_MARGIN,
			AccountType:  kabuspb.AccountType_ACCOUNT_TYPE_SPECIFIC,
			Quantity:     1,
			OrderType:    kabuspb.StockOrderType_STOCK_ORDER_TYPE_MO,
			Price:        0,
			ExpireDay:    nil,
		})
		if err != nil {
			panic(err)
		}
		exitOrderID = res.OrderId
	}

	// 約定確認
	{
		for {
			res, err := cli.GetOrders(context.Background(), &kabuspb.GetOrdersRequest{})
			if err != nil {
				panic(err)
			}

			if checkContracted(exitOrderID, res) {
				fmt.Println("エグジット注文の約定を確認しました")
				break
			}
			<-time.After(time.Second)
		}
	}
}
